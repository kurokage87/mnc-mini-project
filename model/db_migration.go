package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserId      uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	FirstName   string
	LastName    string
	PhoneNumber string `gorm:"unique"`
	Pin         string
	Address     string
	CreatedDate time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	UpdateDate  time.Time `gorm:"type:timestamptz;"`
}

type Balance struct {
	BalanceId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	UserId        uuid.UUID `gorm:"type:uuid"`
	BalanceBefore float64
	BalanceAfter  float64
	CreatedDate   time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
	Status        string    `gorm:"type:varchar(1);comment:L = Latest, O = Oldest"`
}

type Transaction struct {
	TransactionId       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	UserId              uuid.UUID `gorm:"type:uuid"`
	BalanceId           uuid.UUID `gorm:"type:uuid"`
	TransactionType     string    `gorm:"type:varchar(1);comment:D =  DebitC = Credit"`
	Amount              float64
	Status              string    `gorm:"type:varchar(1);comment:S = Success, F = Failed, P = Pending"`
	TransactionCategory string    `gorm:"type:varchar(2);comment:TP = Top Up, PY = Payment, TF = Transfer"`
	Remarks             string    `gorm:"type:varchar"`
	TargetUser          uuid.UUID `gorm:"type:varchar"`
	CreatedDate         time.Time `gorm:"type:timestamptz;default:current_timestamp;"`
}
