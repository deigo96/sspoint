package controller

import (
	"net/http"
	"referralUser-service/handler"
	"referralUser-service/helper"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Service    helper.RefferralService
	jwtService handler.JwtService
	// db          *gorm.DB
}

type Connection struct {
}

func NewAuthController(Service helper.RefferralService, jwtService handler.JwtService) *AuthController {
	return &AuthController{
		Service:    Service,
		jwtService: jwtService,
	}
}

func (controller *AuthController) GetReferralUser(c echo.Context) (err error) {
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

	// d := model.

	// id, err := extractClaims(header, c)
	userData, err := helper.GetIdToken(header)
	var id int
	for _, i := range userData.Results {
		id = i.Parent_id
	}

	if err != nil {
		response := helper.BuildErrorResponse("Failed to extract token", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := controller.Service.GetRef(id, header)
	if err != nil {
		return err
	}

	data := helper.AllUserReferral(res)

	response := helper.BuildSuccessResponse(true, "Success get user referral", data)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) RegisterReferral(c echo.Context) (err error) {
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

	req := new(helper.RegisterReferral)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	userData, err := helper.GetIdToken(header)
	var id int
	for _, i := range userData.Results {
		id = i.Parent_id
	}
	if err != nil {
		response := helper.BuildErrorResponse("Failed to extract token", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.Child_id == 0 {
		response := helper.BuildErrorResponse("Child id is required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = controller.Service.RegisterRef(req.Parent_id, req.Child_id, id)
	if err != nil {
		return err
	}

	response := helper.BuildSuccessResponse(true, "Success create upline user referral", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)

}

func (controller *AuthController) UpdateReferral(c echo.Context) (err error) {
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

	req := new(helper.RegisterReferral)
	if err = c.Bind(req); err != nil {
		response := helper.BuildErrorResponse("Invalid request body", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	userData, err := helper.GetIdToken(header)
	var id int
	for _, i := range userData.Results {
		id = i.Parent_id
	}
	if err != nil {
		response := helper.BuildErrorResponse("Failed to extract token", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if req.To_Parent_id == 0 || req.Child_id == 0 {
		response := helper.BuildErrorResponse("To Parent id and Child id is required", helper.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = controller.Service.UpdateRef(req.Parent_id, req.Child_id, req.To_Parent_id, id)
	if err != nil {
		return err
	}

	response := helper.BuildSuccessResponse(true, "Success Update upline user referral", helper.EmptyObj{})
	return c.JSON(http.StatusOK, response)
}

// func extractClaims(tokenStr string, c echo.Context) (int, error) {
// 	token, _ := jwt.Parse(tokenStr, nil)
// 	if token == nil {
// 		return 0, nil
// 	}
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	id := fmt.Sprintf("%v", claims["sub"])
// 	intID, _ := strconv.Atoi(id)

// 	return intID, nil
// }
