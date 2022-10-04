package controller

import (
	"net/http"
	"reward-list-service/handler"
	"reward-list-service/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service    helper.RewardService
	jwtService handler.JwtService
}

func NewAuthController(Service helper.RewardService, jwtService handler.JwtService) *AuthController {
	return &AuthController{
		Service:    Service,
		jwtService: jwtService,
	}
}

func (controller *AuthController) GetReward(c echo.Context) (err error) {
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

	req := new(helper.GetData)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	res, err := controller.Service.GetReward()

	data := helper.AllRewardList(res)
	response := helper.BuildSuccessResponse(true, "Success get reward list", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) StoreReward(c echo.Context) (err error) {
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

	req := new(helper.StoreRewardRequest)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Reward_Description == "" || req.Reward_Image == "" || req.Reward_Name == "" {
		response := helper.BuildErrorResponse("All fields are required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.StoreRewardRequest{
		Reward_Name:        req.Reward_Name,
		Reward_Image:       req.Reward_Image,
		Reward_Description: req.Reward_Description,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.StoreReward(id, store)

	response := helper.BuildSuccessResponse(true, "Success store reward", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) UpdateReward(c echo.Context) (err error) {
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

	req := new(helper.StoreRewardRequest)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	reqId := new(helper.Idreward)
	if err = c.Bind(reqId); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if reqId.Id_Reward == 0 {
		response := helper.BuildErrorResponse("Id reward is required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	var store = helper.StoreRewardRequest{
		Reward_Name:        req.Reward_Name,
		Reward_Image:       req.Reward_Image,
		Reward_Description: req.Reward_Description,
	}

	id, err := helper.GetIdToken(header)
	_, err = controller.Service.UpdateReward(id, reqId.Id_Reward, store)
	if err != nil {
		response := helper.BuildErrorResponse("Id Reward not found", helper.EmptyObj{})
		return c.JSON(http.StatusOK, response)
	}

	response := helper.BuildSuccessResponse(true, "Success store reward", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}
