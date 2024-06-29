package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mnc/model"
	"mnc/service"
	"net/http"
)

func Register(c *gin.Context) {
	var req model.RequestRegister
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
		return
	}

	data, err := service.Register(req)
	if err != nil {
		fmt.Errorf("can't insert data to database, err : %v", err)
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})
}

func Login(c *gin.Context) {
	var req model.RequestLogin
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
	}

	data, err := service.Login(req)
	if err != nil {
		fmt.Errorf("login failed, err : %v", err)
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})

	return
}

func UpdateProfile(c *gin.Context) {

	userId := c.MustGet("user_id")
	var req model.RequestUpdateProfile
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: "Invalid input request",
		})
	}

	data, err := service.Update(fmt.Sprintf("%s", userId), req)
	if err != nil {
		fmt.Errorf("update failed, err : %v", err)
		c.JSON(http.StatusBadRequest, model.FailedResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Status: "success",
		Data:   data,
	})

	return
}
