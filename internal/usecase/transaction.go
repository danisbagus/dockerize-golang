package usecase

import (
	"github.com/danisbagus/dockerize-golang/internal/dto"
	"github.com/danisbagus/dockerize-golang/internal/repo"
)

type ITransactionUsecase interface {
	GetAll() (*dto.TransactionListResponse, error)
}

type TransactionUsecase struct {
	repo repo.ITransactionRepo
}

func NewTransactionUsecase(repo repo.ITransactionRepo) ITransactionUsecase {
	return &TransactionUsecase{
		repo: repo,
	}
}

func (r TransactionUsecase) GetAll() (*dto.TransactionListResponse, error) {

	transactionsList, err := r.repo.FetchAll()
	if err != nil {
		return nil, err
	}
	response := dto.NewTransactionListResponse(transactionsList)

	return response, nil

}
