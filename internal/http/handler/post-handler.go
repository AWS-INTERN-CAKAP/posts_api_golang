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

type PostHandler struct {
	postService services.PostService
}

func NewPostHandler(postService services.PostService) *PostHandler {
	return &PostHandler{postService: postService}
}


func (h *PostHandler) CreatePost(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)


	var input binder.CreatePost

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}
	input.AuthorID = claims.ID

	data, err := h.postService.CreatePost(&input)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success create post", data))
}
func (h *PostHandler) GetPosts(ctx echo.Context) error {
	posts, err := h.postService.GetPosts()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get posts", posts))
}

func (h *PostHandler) GetPostByID(ctx echo.Context) error {
	var input binder.GetByID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}


	post, err := h.postService.GetPostByID(uuid.MustParse(input.ID))

	if err != nil {
		if err == services.ErrorPostNotFound {
			return ctx.JSON(http.StatusNotFound, response.ErrorResponse(http.StatusNotFound, err.Error()))
		}
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get post", post))
}

func (h *PostHandler) UpdatePost(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)


	var input binder.UpdatePost

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	input.AuthorID = claims.ID

	data, err := h.postService.UpdatePost(&input)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update post", data))
}

func (h *PostHandler) DeletePost(ctx echo.Context) error {
	var input binder.GetByID

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	postID := uuid.MustParse(input.ID)

	err := h.postService.DeletePost(postID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success delete post", nil))
}