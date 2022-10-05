package controller

import (
	"net/http"
	"transactionPoint-service/handler"
	"transactionPoint-service/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service    helper.TransactionPointService
	jwtService handler.JwtService
	// db          *gorm.DB
}

func NewAuthController(Service helper.TransactionPointService, jwtService handler.JwtService) *AuthController {
	return &AuthController{
		Service:    Service,
		jwtService: jwtService,
	}
}

func (controller *AuthController) StoreTransactionPoint(c echo.Context) (err error) {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		resp := helper.BuildErrorResponse("No Authorization, No Token", helper.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, resp)
	}

	if validate != true {
		resp := helper.BuildErrorResponse("Token expired or invalid", helper.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, resp)
	}

	req := new(helper.TransactionPoint)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Trx_Type_Id == 0 || req.Reward_Id == 0 || req.Point == 0 || req.Own_Pct == 0 || req.Referral_Pct == 0 || req.Branch_Pct == 0 {
		response := helper.BuildErrorResponse("All fields are required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.TransactionPoint{
		Trx_Type_Id:  req.Trx_Type_Id,
		Reward_Id:    req.Reward_Id,
		Point:        req.Point,
		Own_Pct:      req.Own_Pct,
		Referral_Pct: req.Referral_Pct,
		Branch_Pct:   req.Branch_Pct,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.StoreTransactionPointService(id, store)

	response := helper.BuildSuccessResponse(true, "Success store transaction point", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)

}

func (controller *AuthController) UpdateTransactionPoint(c echo.Context) (err error) {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		resp := helper.BuildErrorResponse("No Authorization, No Token", helper.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, resp)
	}

	if validate != true {
		resp := helper.BuildErrorResponse("Token expired or invalid", helper.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, resp)
	}

	req := new(helper.TransactionPoint)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	reqId := new(helper.IdTransactionPoint)
	if err = c.Bind(reqId); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if reqId.Id_Transaction_Point == 0 {
		response := helper.BuildErrorResponse("Id Transaction is required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.TransactionPoint{
		Trx_Type_Id:  req.Trx_Type_Id,
		Reward_Id:    req.Reward_Id,
		Point:        req.Point,
		Own_Pct:      req.Own_Pct,
		Referral_Pct: req.Referral_Pct,
		Branch_Pct:   req.Branch_Pct,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.UpdateTransactionPointService(id, reqId.Id_Transaction_Point, store)
	if err != nil {
		response := helper.BuildErrorResponse("Id Transaction Point not found", helper.EmptyObj{})
		return c.JSON(http.StatusOK, response)
	}

	response := helper.BuildSuccessResponse(true, "Success edit transaction point", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
