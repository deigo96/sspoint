package controller

import (
	"net/http"
	"pointHistory-service/handler"
	"pointHistory-service/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service    helper.PointHistoryService
	jwtService handler.JwtService
}

func NewAuthController(Service helper.PointHistoryService, jwtService handler.JwtService) *AuthController {
	return &AuthController{
		Service:    Service,
		jwtService: jwtService,
	}
}

func (controller *AuthController) GetPointHistoryService(c echo.Context) (err error) {
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

	req := new(helper.TrxPointRequest)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.FromTrxDate == "" || req.ToTrxDate == "" || req.FromTrxDate > req.ToTrxDate {
		response := helper.BuildErrorResponse("Invalid fromTrxDate or to toTrxDate", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	store := helper.TrxPointRequest{
		FromTrxDate: req.FromTrxDate,
		ToTrxDate:   req.ToTrxDate,
	}

	id, err := helper.GetIdToken(header)
	res, err := controller.Service.GetPointHistoryService(id, store)

	data := helper.AllPointHisList(res)
	response := helper.BuildSuccessResponse(true, "Success get point history", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) StorePointHistoryService(c echo.Context) (err error) {
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

	req := new(helper.StoreTrxPointReq)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Trx_date == "" || req.Trx_point_id == 0 || req.Reference_trx_id == 0 {
		response := helper.BuildErrorResponse("All fields are required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	id, err := helper.GetIdToken(header)

	store := helper.StoreTrxPointReq{
		Trx_date:         req.Trx_date,
		User_id:          id,
		Is_branch:        req.Is_branch,
		Trx_point_id:     req.Trx_point_id,
		Reference_trx_id: req.Reference_trx_id,
		Ss_point_before:  req.Ss_point_before,
		Ss_point_trx:     req.Ss_point_trx,
		Ss_point_after:   req.Ss_point_after,
	}

	if err = controller.Service.StorePointHistoryService(store); err != nil {
		response := helper.BuildErrorResponse("Fail to store data", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.BuildSuccessResponse(true, "Success post point history", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
