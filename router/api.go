package router

import (
	"github.com/labstack/echo/v4"
	"go-backend/handler"
	"go-backend/middleware"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	CateHandler    handler.CateHandler
}

func (api *API) SetupRouter() {

	// user
	user := api.Echo.Group("/user")
	user.POST("/sign-up", api.UserHandler.HandleSignUp)
	user.POST("/sign-in", api.UserHandler.HandleSignIn)
	user.GET("/profile/", api.UserHandler.HandleProfile, middleware.JWTMiddleware())
	user.GET("/list/", api.UserHandler.HandleListUsers, middleware.JWTMiddleware())

	// categories
	categories := api.Echo.Group("/cate")

	categories.POST("/add", api.CateHandler.HandleAddCate)
	categories.PUT("/edit", api.CateHandler.HandleEditCate)
	categories.GET("/detail/:id", api.CateHandler.HandleCateDetail)
	categories.GET("/list", api.CateHandler.HandleCateList)

}

func (api *API) SetupAdminRouter()  {
	admin := api.Echo.Group("/admin")
	admin.GET("/token", api.UserHandler.GenToken)
	admin.POST("/sign-in", api.UserHandler.HandleAdminSignIn)
	admin.POST("/sign-up", api.UserHandler.HandleAdminSignUp , middleware.JWTMiddleware())
}
