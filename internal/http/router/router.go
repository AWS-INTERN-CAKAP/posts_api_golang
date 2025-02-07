package router

import (
	"net/http"

	"github.com/Davidafdal/e-commerce/internal/http/handler"
	"github.com/Davidafdal/e-commerce/pkg/route"
)

func AppPublicRoutes(appHandler handler.AppHandler) []*route.Route {
	userHandler := appHandler.UserHandler
	categoryHandler := appHandler.CategoryHandler
	postHandler := appHandler.PostHandler

	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: userHandler.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/register",
			Handler: userHandler.Register,
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories",
			Handler: categoryHandler.GetCategories,
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories/:id",
			Handler: categoryHandler.GetCategoryByID,
		},
		{
			Method:  http.MethodGet,
			Path:    "/posts",
			Handler: postHandler.GetPosts,
		},
		{
			Method:  http.MethodGet,
			Path:    "/posts/:id",
			Handler: postHandler.GetPostByID,
		},	
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: appHandler.UserHandler.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
		},
	}
}

func AppPrivateRoutes(appHandler handler.AppHandler) []*route.Route {

	userHandler := appHandler.UserHandler
	categoryHandler := appHandler.CategoryHandler
	postHandler := appHandler.PostHandler

	return []*route.Route{
	
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: categoryHandler.CreateCategory,
		},
		{
			Method:  http.MethodPut,
			Path:    "/categories/:id",
			Handler: categoryHandler.UpdateCategory,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/categories/:id",
			Handler: categoryHandler.DeleteCategory,
		},
		{
			Method:  http.MethodPost,
			Path:    "/posts",
			Handler: postHandler.CreatePost,
		},
		{
			Method:  http.MethodPut,
			Path:    "/posts/:id",
			Handler: postHandler.UpdatePost,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/posts/:id",
			Handler: postHandler.DeletePost,
		},
		{
			Method:  http.MethodGet,
			Path:    "/profile",
			Handler: userHandler.GetProfile,
		},
		{
			Method:  http.MethodPost,
			Path:    "/logout",
			Handler: userHandler.Logout,
		},
	}
}