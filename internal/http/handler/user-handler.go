package handler

import (
	"net/http"
	"strings"

	"github.com/Davidafdal/e-commerce/internal/http/binder"
	"github.com/Davidafdal/e-commerce/internal/services"
	"github.com/Davidafdal/e-commerce/pkg/response"
	"github.com/Davidafdal/e-commerce/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService}
}


func (h *UserHandler) Login(ctx echo.Context) error {
	var input binder.LoginRequst

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}


	responseData, err := h.userService.Login(&input)

	if err != nil {
		if err == services.ErrorInvalidUsers {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}



	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success login", responseData))
}

func (h *UserHandler) Register(ctx echo.Context) error {
	var input binder.RegisterRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}


	responseData, err := h.userService.Register(&input)

	if err != nil {
		if err == services.ErrorExitedUser {
			return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success register", responseData))
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	users, err := h.userService.GetUsers()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get users", users))
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	var input binder.GetByID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	
	user, err := h.userService.GetUserByID(input.ID)

	if err != nil {
		if err == services.ErrorUserNotFound {
			return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get user", user))
}


func (h *UserHandler) GetProfile(ctx echo.Context) error {

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	user, err := h.userService.GetUserByID(claims.ID)

	if err != nil {
		if err == services.ErrorUserNotFound {
			return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))	
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get profile", user))
}

func (h *UserHandler) Logout(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = tokenString[7:]
	} else {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid authorization header format"))
	}

	err := h.userService.Logout(tokenString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "sukses logout", nil))
}





