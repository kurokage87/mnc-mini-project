package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mnc/constant"
	"mnc/model"
	"mnc/service"
	"net/http"
)

func Topup(c *gin.Context) {
	userId := c.MustGet("user_id")
	var req model.RequestTransaction
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
		return
	}

	req.TransactionType = constant.TopUpTransaction
	data, err := service.Transaction(fmt.Sprintf("%s", userId), req)
	if err != nil {
		if err != nil {
			fmt.Errorf("insert topup failed, err : %v", err)
			c.JSON(http.StatusBadRequest, model.FailedResponse{
				Status:  "failed",
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})
	return
}

func Payment(c *gin.Context) {
	userId := c.MustGet("user_id")
	var req model.RequestTransaction
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
		return
	}

	req.TransactionType = constant.PaymentTransaction
	data, err := service.Transaction(fmt.Sprintf("%s", userId), req)
	if err != nil {
		if err != nil {
			fmt.Errorf("insert topup failed, err : %v", err)
			c.JSON(http.StatusBadRequest, model.FailedResponse{
				Status:  "failed",
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})
	return
}

func Transfer(c *gin.Context) {
	userId := c.MustGet("user_id")
	var req model.RequestTransaction
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
		return
	}

	if req.TargetUser == uuid.Nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "target user can't blank",
		})
		return
	}

	data, err := service.TransferTransaction(fmt.Sprintf("%s", userId), req)
	if err != nil {
		if err != nil {
			fmt.Errorf("insert topup failed, err : %v", err)
			c.JSON(http.StatusBadRequest, model.FailedResponse{
				Status:  "failed",
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})
	return
}

func Report(c *gin.Context) {
	userId := c.MustGet("user_id")

	data, err := service.ListTransactions(fmt.Sprintf("%s", userId))
	if err != nil {
		if err != nil {
			fmt.Errorf("insert topup failed, err : %v", err)
			c.JSON(http.StatusBadRequest, model.FailedResponse{
				Status:  "failed",
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})
	return
}
