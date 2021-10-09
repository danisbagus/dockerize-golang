package repo

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type TransactionModel struct {
	TransactionID string `db:"transaction_id"`
	MerchantID    int64  `db:"merchant_id"`
	SKUID         string `db:"sku_id"`
	SuppierID     int64  `db:"supplier_id"`
	Quantity      int64  `db:"quantity"`
	TotalPrice    int64  `db:"total_price"`
	CreatedAt     string `db:"created_at"`
}

type ITransactionRepo interface {
	FetchAll() ([]TransactionModel, error)
}

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) ITransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (r TransactionRepo) FetchAll() ([]TransactionModel, error) {
	transactions := make([]TransactionModel, 0)

	fetchAllTrasactionQuery := `select * from transactions`

	err := r.db.Select(&transactions, fetchAllTrasactionQuery)

	if err != nil {
		log.Printf("Error while quering find all purchase transaction by merchant id " + err.Error())
		return nil, err
	}

	return transactions, nil
}
