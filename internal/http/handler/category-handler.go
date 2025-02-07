package handler

import (
	"net/http"

	"github.com/Davidafdal/e-commerce/internal/http/binder"
	"github.com/Davidafdal/e-commerce/internal/services"
	"github.com/Davidafdal/e-commerce/pkg/response"
	"github.com/Davidafdal/e-commerce/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService services.CategoryService
}

func  NewCategoryHandler(categoryService services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService}
}


func (h *CategoryHandler) GetCategories(ctx echo.Context) error{
	categories, err := h.categoryService.GetCategories()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get categories", categories))
}

func (h *CategoryHandler) GetCategoryByID(ctx echo.Context) error{
	var input binder.GetByID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	categoryID := uuid.MustParse(input.ID)
	category, err := h.categoryService.GetCategoryByID(categoryID)

	if err != nil {
		if err == services.ErrorCategoryNotFound {
			return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get category", category))
}

func (h *CategoryHandler) CreateCategory(ctx echo.Context) error{
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	userID := uuid.MustParse(claims.ID)

	var input binder.CreateCategory

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	input.CreatedBy = userID

	data, err := h.categoryService.CreateCategory(&input)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success create category", data))
}

func (h *CategoryHandler) UpdateCategory(ctx echo.Context) error{

	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	userID := uuid.MustParse(claims.ID)

	var input binder.UpdateCategory


	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	input.CreatedBy = userID

	data, err := h.categoryService.UpdateCategory(&input)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update category", data))
}

func (h *CategoryHandler) DeleteCategory(ctx echo.Context) error{
	var input binder.GetByID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	categoryID := uuid.MustParse(input.ID)
	err := h.categoryService.DeleteCategory(categoryID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success delete category", nil))
}


