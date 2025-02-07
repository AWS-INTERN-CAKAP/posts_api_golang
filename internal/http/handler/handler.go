package handler

import "github.com/Davidafdal/e-commerce/pkg/validator"

type AppHandler struct {
	UserHandler *UserHandler
	CategoryHandler *CategoryHandler
	PostHandler *PostHandler
}

func NewAppHandler(userHandler *UserHandler, categoryHandler *CategoryHandler, postHandler *PostHandler) AppHandler {
	return AppHandler{
		UserHandler: userHandler,
		CategoryHandler: categoryHandler,
		PostHandler: postHandler,
	}
}

func checkValidation(input interface{}) (errorMessage string, data interface{}) {
	validationErrors := validator.Validate(input)
	if validationErrors != nil {
		if _, exists := validationErrors["error"]; exists {
			return "validasi input gagal", nil
		}
		return "validasi input gagal", validationErrors
	}
	return "", nil
}