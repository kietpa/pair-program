package controller

import (
	"pair/model"
	"pair/repository"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionController(transactionRepository repository.TransactionRepository) *TransactionController {
	return &TransactionController{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransactionController) CreateTranscation(c echo.Context) error {
	var newTranscation model.Transaction
	if err := c.Bind(&newTranscation); err != nil {
		return c.JSON(400, echo.Map{
			"message": "invalid request",
		})
	}

	if err := t.TransactionRepository.Create(&newTranscation); err != nil {
		return c.JSON(500, echo.Map{
			"message": "failed to create transaction",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "transaction created successfully",
	})
}

func (t *TransactionController) GetAllTransaction(c echo.Context) error {
	transaction, err := t.TransactionRepository.ReadAll()
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "transaction not found",
		})
	}

	return c.JSON(200, echo.Map{
		"transactions": transaction,
	})
}

func (t *TransactionController) GetTransactionByID(c echo.Context) error {
	transactionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "invalid transaction ID",
		})
	}

	transaction, err := t.TransactionRepository.ReadID(transactionID)
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": "transaction not found",
		})
	}

	return c.JSON(200, transaction)
}

func (t *TransactionController) UpdateTransaction(c echo.Context) error {
	transactionID := c.Param("id")
	var updatedTransaction model.Transaction
	if err := c.Bind(&updatedTransaction); err != nil {
		return c.JSON(400, echo.Map{
			"message": "invalid request",
		})
	}

	err := t.TransactionRepository.Update(transactionID, updatedTransaction)
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "failed to update transaction",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "transaction updated successfully",
	})
}

func (t *TransactionController) Deletetransaction(c echo.Context) error {
	transactionID := c.Param("id")

	if err := t.TransactionRepository.Delete(transactionID); err != nil {
		return c.JSON(500, echo.Map{
			"message": "failed to delete transaction",
		})
	}

	return c.JSON(200, echo.Map{
		"message": "transaction deleted successfully",
	})
}
