package router

import (
	"brostools-api-person/infrastructure"
	"brostools-api-person/interfaces/auth"
	"brostools-api-person/interfaces/handler"
	"brostools-api-person/usecase"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) error {
	r := e.Group("/api/person")
	r.Use(auth.Secret())
	profileBasicRepository := infrastructure.NewApiPersonfrastructure()
	profilRepository := infrastructure.NewApiPersonBrantectfrastructure("pgsql")
	expFrUseCase := usecase.NewApiPersonUsecase(profileBasicRepository, profilRepository)
	expFreHandler := handler.NewApiPersonHandler(expFrUseCase)
	r.GET("/", expFreHandler.HandleGetAllPerson())
	r.POST("/", expFreHandler.HandleAddPerson())
	r.GET("/:client_cd/:person_cd", expFreHandler.HandleGePersontById())
	r.PUT("/:client_cd/:person_cd", expFreHandler.HandleUpdatePerson())
	r.DELETE("/:client_cd/:person_cd", expFreHandler.HandleDeletePerson())
	return nil
}
