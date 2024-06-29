package constant

const (
	BalanceStatusLatest         = "L"
	BalanceStatusOldest         = "O"
	TransactionTypeDebit        = "D"
	TransactionTypeCredit       = "C"
	TransactionStatusSuccess    = "S"
	TransactionStatusFailed     = "F"
	TransactionStatusPending    = "P"
	TransactionCategoryTopUp    = "TP"
	TransactionCategoryTransfer = "TF"
	TransactionCategoryPayment  = "PY"

	DateLayout         = "2006-1-2 15:04:05"
	PaymentTransaction = "PAYMENT"
	TopUpTransaction   = "TOP-UP"

	Payment  = "PAYMENT"
	Transfer = "TRANSFER"
	TOPUP    = "TOP UP"

	TransactionPrefix = "transaction:%s"
)
