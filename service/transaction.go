package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mnc/config"
	"mnc/constant"
	"mnc/model"
	"time"
)

func Transaction(user_id string, req model.RequestTransaction) (data interface{}, err error) {

	db := config.DB
	//Check balance latest
	var balance model.Balance
	if err = db.Where("user_id = ? and status = ?", user_id, constant.BalanceStatusLatest).Order("created_date desc").First(&balance).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println(err)
		return
	}

	var lastBalance float64 = 0
	var newBalance float64 = 0
	now := time.Now()

	uuidUser, err := uuid.Parse(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if balance.BalanceId == uuid.Nil {
		lastBalance = 0

		if req.TransactionType == constant.TopUpTransaction {
			newBalance = lastBalance + req.Amount
		} else {
			newBalance = lastBalance - req.Amount
			if newBalance < 0 {
				err = errors.New("“Balance is not enough")
				return
			}
		}
	} else {
		lastBalance = balance.BalanceAfter
		if req.TransactionType == constant.TopUpTransaction {
			newBalance = lastBalance + req.Amount
		} else {
			newBalance = lastBalance - req.Amount
			if newBalance < 0 {
				err = errors.New("Balance is not enough")
				return
			}
		}

		// disable latest balance data
		if err = db.Model(&model.Balance{}).Where("user_id = ? and balance_id = ?", user_id, balance.BalanceId).Update("status", constant.BalanceStatusOldest).Error; err != nil {
			err = errors.New(fmt.Sprintf("failed update data balance for user_id %s and balance_id %s", user_id, balance.BalanceId))
			return
		}
	}

	//insert new balance data
	newInstBalance := model.Balance{
		BalanceId:     uuid.New(),
		UserId:        uuidUser,
		Status:        constant.BalanceStatusLatest,
		BalanceBefore: lastBalance,
		BalanceAfter:  newBalance,
		CreatedDate:   now,
	}
	if err = db.Create(&newInstBalance).Error; err != nil {
		err = errors.New(fmt.Sprintf("failed insert new data balance user_id %s", user_id))
		return
	}

	// insert data to table transaction
	var transaction model.Transaction
	switch req.TransactionType {
	case constant.TopUpTransaction:
		transaction = model.Transaction{
			TransactionId:       uuid.New(),
			BalanceId:           newInstBalance.BalanceId,
			UserId:              uuidUser,
			Amount:              req.Amount,
			TransactionType:     constant.TransactionTypeCredit,
			Status:              constant.TransactionStatusPending,
			TransactionCategory: constant.TransactionCategoryTopUp,
			CreatedDate:         now,
		}
	case constant.PaymentTransaction:
		transaction = model.Transaction{
			TransactionId:       uuid.New(),
			BalanceId:           newInstBalance.BalanceId,
			UserId:              uuidUser,
			Amount:              req.Amount,
			Remarks:             req.Remarks,
			TransactionType:     constant.TransactionTypeDebit,
			Status:              constant.TransactionStatusPending,
			TransactionCategory: constant.TransactionCategoryPayment,
			CreatedDate:         now,
		}
	default:
		err = errors.New("error transaction type")
		return
	}

	if err = db.Create(&transaction).Error; err != nil {
		err = errors.New("failed insert new data to table transactions")
	}

	switch req.TransactionType {
	case constant.TopUpTransaction:
		data = model.DataResponseTopUp{
			TopUpId:       transaction.TransactionId,
			AmountTopUp:   transaction.Amount,
			BalanceBefore: newInstBalance.BalanceBefore,
			BalanceAfter:  newInstBalance.BalanceAfter,
			CreatedDate:   now.Format(constant.DateLayout),
		}
	case constant.PaymentTransaction:
		data = model.DataResponsePayment{
			PaymentId:     transaction.TransactionId,
			Amount:        transaction.Amount,
			Remarks:       transaction.Remarks,
			BalanceBefore: newInstBalance.BalanceBefore,
			BalanceAfter:  newInstBalance.BalanceAfter,
			CreatedDate:   now.Format(constant.DateLayout),
		}
	default:
		err = errors.New("error get response transaction")
		return
	}

	//save queue to redis
	redisClient := config.RedisClient
	err = redisClient.Set(context.Background(), fmt.Sprintf(constant.TransactionPrefix, transaction.TransactionId), fmt.Sprintf("%s", transaction.TransactionId), 0).Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("failed set transaction to redis id %s", transaction.TransactionId))
		return
	}
	return
}

func TransferTransaction(user_id string, req model.RequestTransaction) (data model.DataResponseTransfer, err error) {
	db := config.DB
	now := time.Now()
	uuidUser, err := uuid.Parse(user_id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// get latest balance current user
	var latestBalanceCurrentUser model.Balance
	if err = db.Where("user_id = ? and status = ?", user_id, constant.BalanceStatusLatest).Order("created_date desc").First(&latestBalanceCurrentUser).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New(fmt.Sprintf("failed get data balance for current user_id %s", req.TargetUser))
		return
	}

	// get latest balance target user
	var latestBalanceTargetUser model.Balance
	if err = db.Model(&model.Balance{}).Where("user_id = ? and status = ?", req.TargetUser, constant.BalanceStatusLatest).Order("created_date desc").First(&latestBalanceTargetUser).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New(fmt.Sprintf("failed get data balance for target user_id %s", req.TargetUser))
		return
	}

	var lastCurrentBalance float64 = 0
	var newCurrentBalance float64 = 0

	var lastTargetBalance float64 = 0
	var newTargetBalance float64 = 0

	if latestBalanceCurrentUser.BalanceId == uuid.Nil && latestBalanceTargetUser.BalanceId == uuid.Nil {
		lastCurrentBalance = 0
		lastTargetBalance = 0
		newCurrentBalance = lastCurrentBalance - req.Amount
		newTargetBalance = lastTargetBalance + req.Amount

		if newCurrentBalance < 0 {
			err = errors.New("balance is not enough")
			return
		}
	} else if latestBalanceCurrentUser.BalanceId != uuid.Nil && latestBalanceTargetUser.BalanceId == uuid.Nil {
		lastCurrentBalance = latestBalanceCurrentUser.BalanceAfter
		lastTargetBalance = 0
		newCurrentBalance = lastCurrentBalance - req.Amount
		newTargetBalance = lastTargetBalance + req.Amount

		if newCurrentBalance < 0 {
			err = errors.New("balance is not enough")
			return
		}

		// disable latest balance data
		if err = db.Model(&model.Balance{}).Where("user_id = ? and balance_id = ?", user_id, latestBalanceCurrentUser.BalanceId).Update("status", constant.BalanceStatusOldest).Error; err != nil {
			err = errors.New(fmt.Sprintf("failed update data balance for user_id %s and balance_id %s", user_id, latestBalanceCurrentUser.BalanceId))
			return
		}
	} else {
		lastCurrentBalance = latestBalanceCurrentUser.BalanceAfter
		lastTargetBalance = latestBalanceTargetUser.BalanceAfter
		newCurrentBalance = lastCurrentBalance - req.Amount
		newTargetBalance = lastTargetBalance + req.Amount

		if newCurrentBalance < 0 {
			err = errors.New("balance is not enough")
			return
		}

		arrUserId := make([]uuid.UUID, 0)
		arrBalanceId := make([]uuid.UUID, 0)

		arrUserId = append(arrUserId, latestBalanceCurrentUser.UserId, latestBalanceTargetUser.UserId)
		arrBalanceId = append(arrBalanceId, latestBalanceCurrentUser.BalanceId, latestBalanceTargetUser.BalanceId)

		// disable latest balance data
		if err = db.Model(&model.Balance{}).Where("user_id IN (?) and balance_id IN (?)", arrUserId, arrBalanceId).Update("status", constant.BalanceStatusOldest).Error; err != nil {
			err = errors.New(fmt.Sprintf("failed update data balance for user_id %s and balance_id %s", arrUserId, arrBalanceId))
			return
		}
	}

	//insert new balance current data
	newInsertCurrentBalance := model.Balance{
		BalanceId:     uuid.New(),
		UserId:        uuidUser,
		Status:        constant.BalanceStatusLatest,
		BalanceBefore: lastCurrentBalance,
		BalanceAfter:  newCurrentBalance,
		CreatedDate:   now,
	}

	newInsertTargetBalance := model.Balance{
		BalanceId:     uuid.New(),
		UserId:        req.TargetUser,
		Status:        constant.BalanceStatusLatest,
		BalanceBefore: lastTargetBalance,
		BalanceAfter:  newTargetBalance,
		CreatedDate:   now,
	}
	if err = db.Create(&newInsertCurrentBalance).Error; err != nil {
		err = errors.New(fmt.Sprintf("failed insert new data balance user_id %s", user_id))
		return
	}

	if err = db.Create(&newInsertTargetBalance).Error; err != nil {
		err = errors.New(fmt.Sprintf("failed insert new data balance user_id %s", user_id))
		return
	}

	//insert data transaction
	var transactionCurrent model.Transaction
	var transactionTarget model.Transaction
	transactionCurrent = model.Transaction{
		TransactionId:       uuid.New(),
		BalanceId:           newInsertCurrentBalance.BalanceId,
		UserId:              newInsertCurrentBalance.UserId,
		Amount:              req.Amount,
		TransactionType:     constant.TransactionTypeCredit,
		Status:              constant.TransactionStatusPending,
		TransactionCategory: constant.TransactionCategoryTransfer,
		Remarks:             req.Remarks,
		CreatedDate:         now,
	}

	transactionTarget = model.Transaction{
		TransactionId:       uuid.New(),
		BalanceId:           newInsertCurrentBalance.BalanceId,
		UserId:              newInsertCurrentBalance.UserId,
		Amount:              req.Amount,
		TransactionType:     constant.TransactionTypeDebit,
		Status:              constant.TransactionStatusPending,
		TransactionCategory: constant.TransactionCategoryTransfer,
		CreatedDate:         now,
	}

	if err = db.Create(&transactionCurrent).Error; err != nil {
		err = errors.New("failed insert new data current user to table transactions")
		return
	}

	if err = db.Create(&transactionTarget).Error; err != nil {
		err = errors.New("failed insert new data target to table transactions")
		return
	}

	data = model.DataResponseTransfer{
		TransferId:    transactionCurrent.TransactionId,
		Amount:        transactionCurrent.Amount,
		Remarks:       transactionCurrent.Remarks,
		BalanceBefore: newInsertCurrentBalance.BalanceBefore,
		BalanceAfter:  newInsertCurrentBalance.BalanceAfter,
		CreatedDate:   now.Format(constant.DateLayout),
	}

	//save queue to redis
	redisClient := config.RedisClient
	err = redisClient.Set(context.Background(), fmt.Sprintf(constant.TransactionPrefix, transactionCurrent.TransactionId), fmt.Sprintf("%s", transactionCurrent.TransactionId), 0).Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("failed set transaction to redis id %s", transactionCurrent.TransactionId))
		return
	}

	err = redisClient.Set(context.Background(), fmt.Sprintf(constant.TransactionPrefix, transactionTarget.TransactionId), fmt.Sprintf("%s", transactionTarget.TransactionId), 0).Err()
	if err != nil {
		err = errors.New(fmt.Sprintf("failed set transaction to redis id %s", transactionTarget.TransactionId))
		return
	}

	return
}

func ListTransactions(user_id string) (data []interface{}, err error) {
	db := config.DB

	var listTransaction []model.ListTransactions
	query := `select
				t.transaction_id ,
				t.user_id ,
				case
					when t.transaction_type = 'C' then 'CREDIT'
					else 'DEBIT'
				end transaction_type,
				case
					when t.transaction_category = 'TP' then 'TOP UP'
					when t.transaction_category = 'PY' then 'PAYMENT'
					else 'TRANSFER'
				end transaction_category,
				case
					when t.status = 'P' then 'PENDING'
					when t.status = 'S' then 'SUCCESS'
					else 'FAILED'
				end status,
    			t.amount,
				t.remarks,
				b.balance_before ,
				b.balance_after ,
				t.created_date
			from
				transactions t
			inner join users u on
				u.user_id = t.user_id
			inner join balances b on
				b.balance_id = t.balance_id
			where
				t.user_id = ?`

	if err = db.Raw(query, user_id).Scan(&listTransaction).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New(fmt.Sprintf("failed get list transactions for user_id %s", user_id))
		return
	}

	for _, v := range listTransaction {
		if v.TransactionCategory == constant.Payment {
			data = append(data, model.PaymentListTransactions{
				PaymentId:       v.TransactionID,
				Status:          v.Status,
				UserId:          v.UserId,
				TransactionType: v.TransactionType,
				Amount:          v.Amount,
				Remarks:         v.Remarks,
				BalanceBefore:   v.BalanceBefore,
				BalanceAfter:    v.BalanceAfter,
				CreatedDate:     v.CreatedDate,
			})
		} else if v.TransactionCategory == constant.Transfer {
			data = append(data, model.TransferListTransactions{
				TransferId:      v.TransactionID,
				Status:          v.Status,
				UserId:          v.UserId,
				TransactionType: v.TransactionType,
				Amount:          v.Amount,
				Remarks:         v.Remarks,
				BalanceBefore:   v.BalanceBefore,
				BalanceAfter:    v.BalanceAfter,
				CreatedDate:     v.CreatedDate,
			})
		} else {
			data = append(data, model.TopUpListTransactions{
				TopUpId:         v.TransactionID,
				Status:          v.Status,
				UserId:          v.UserId,
				TransactionType: v.TransactionType,
				Amount:          v.Amount,
				Remarks:         v.Remarks,
				BalanceBefore:   v.BalanceBefore,
				BalanceAfter:    v.BalanceAfter,
				CreatedDate:     v.CreatedDate,
			})
		}
	}
	return
}

func UpdateStatus(ctx context.Context) (err error) {

	// Scan to get the keys with the specified prefix
	db := config.DB
	var cursor uint64
	var keys []string
	keys, cursor, err = config.RedisClient.Scan(ctx, cursor, fmt.Sprintf(constant.TransactionPrefix, "*"), 10).Result()
	if err != nil {
		fmt.Printf("Could not scan keys with prefix %s: %v", constant.TransactionPrefix, err)
		return
	}

	if len(keys) == 0 {
		fmt.Printf("No keys found with prefix %s\n", constant.TransactionPrefix)
		return
	}

	// Fetch the values for the keys
	values, err := config.RedisClient.MGet(ctx, keys...).Result()
	if err != nil {
		fmt.Printf("Could not get values for keys: %v", err)
		return
	}

	// loop data for update database and delete data in redis
	go func() (err error) {
		for i, _ := range keys {
			if errUpdate := db.Model(&model.Transaction{}).Where("transaction_id = ?", values[i]).Update("status", constant.TransactionStatusSuccess).Error; err != nil {
				errUpdate = errors.New(fmt.Sprintf("failed update transaction to redis id %s", values[i]))
				return errUpdate
			}

			_, errDeleteRedis := config.RedisClient.Del(ctx, fmt.Sprintf(constant.TransactionPrefix, values[i])).Result()
			if errDeleteRedis != nil {
				db.Model(&model.Transaction{}).Where("transaction_id = ?", values[i]).Update("status", constant.TransactionStatusFailed)
				errDeleteRedis = errors.New(fmt.Sprintf("failed delete transaction to redis id %s", values[i]))
				return errDeleteRedis
			}
		}
		return
	}()
	return
}
