package builder

import (
	"github.com/Davidafdal/e-commerce/internal/http/handler"
	"github.com/Davidafdal/e-commerce/internal/http/router"
	"github.com/Davidafdal/e-commerce/internal/repositories"
	"github.com/Davidafdal/e-commerce/internal/services"
	"github.com/Davidafdal/e-commerce/pkg/route"
	"github.com/Davidafdal/e-commerce/pkg/token"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	appHandler := handler.NewAppHandler(userHandler, categoryHandler, postHandler)


	return router.AppPublicRoutes(appHandler)
}

func BuildAppPrivateRoutes(db *gorm.DB, tokenUseCase token.TokenUseCase) []*route.Route {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, tokenUseCase)
	userHandler := handler.NewUserHandler(userService)


	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	appHandler := handler.NewAppHandler(userHandler, categoryHandler, postHandler)

	return router.AppPrivateRoutes(appHandler)
}