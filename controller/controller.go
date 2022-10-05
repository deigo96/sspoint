package controller

import (
	"net/http"
	"referralUser-service/handler"
	"referralUser-service/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service    helper.TransactionTypeService
	jwtService handler.JwtService
	// db          *gorm.DB
}

func NewAuthController(Service helper.TransactionTypeService, jwtService handler.JwtService) *AuthController {
	return &AuthController{
		Service:    Service,
		jwtService: jwtService,
	}
}

func (controller *AuthController) StoreTransactionType(c echo.Context) (err error) {
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

	req := new(helper.TransactionType)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Type_Name == "" || req.Type_Description == "" {
		response := helper.BuildErrorResponse("All fields are required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.TransactionType{
		Type_Name:        req.Type_Name,
		Type_Description: req.Type_Description,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.StoreTransactionTypeService(id, store)

	response := helper.BuildSuccessResponse(true, "Success add new transaction type", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)

}

func (controller *AuthController) UpdateTransactionType(c echo.Context) (err error) {
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

	req := new(helper.TransactionType)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	reqId := new(helper.IdTransactionType)
	if err = c.Bind(reqId); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if reqId.Id_Transaction_type == 0 {
		response := helper.BuildErrorResponse("Id Transaction is required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.TransactionType{
		Type_Name:        req.Type_Name,
		Type_Description: req.Type_Description,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.UpdateTransactionTypeService(id, reqId.Id_Transaction_type, store)
	if err != nil {
		response := helper.BuildErrorResponse("Id Transaction not found", helper.EmptyObj{})
		return c.JSON(http.StatusOK, response)
	}

	response := helper.BuildSuccessResponse(true, "Success edit transaction type", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
