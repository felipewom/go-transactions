package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-transactions/api/presenter"
	"github.com/rafaelbmateus/go-transactions/internal/domain/entity/transaction"
	"github.com/rafaelbmateus/go-transactions/internal/domain/usecase"
)

// TransactionRouters is the endpoints of the server
func TransactionRouters(e *gin.Engine, u usecase.UseCase) {
	e.GET("/transactions", GetTransactions(u))
	e.POST("/transactions", RegisterTransaction(u))
}

// GetTransactions get all transactions
func GetTransactions(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactions, err := u.GetTransactions()
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"Erro to get the transactions",
				err.Error(), "", http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, transactions)
	}
}

// RegisterTransaction register a new transaction
func RegisterTransaction(u usecase.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var t presenter.Transaction
		err := c.BindJSON(&t)
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusBadRequest),
				"Transactions can't be created",
				"Error when converting the parameters sent to json", "", http.StatusBadRequest)
			return
		}

		err = u.RegisterTransaction(&transaction.Transaction{
			ID:              t.ID,
			AccountID:       t.AccountID,
			OperationTypeID: t.OperationTypeID,
			Amount:          t.Amount,
		})
		if err != nil {
			responseFailure(c, http.StatusText(http.StatusInternalServerError),
				"account can't be created",
				"internal server error when creating a new account", "", http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusCreated)
	}
}
