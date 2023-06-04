package main

import (
	"log"

	"go-sample/handler"

	"go-sample/infra"

	"go-sample/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// engineを作成
	engine := infra.DBInit()

	// サービスを作成
	factory := service.NewService(engine)

	// 最後にengineを閉じる
	defer func() {
		log.Println("engine closed")
		engine.Close()
	}()

	// Ginを作成
	g := gin.Default()

	// サービス層のMiddlewareを作成
	g.Use(service.ServiceFactoryMiddleware(factory))

	// v1というグループを作成
	routes := g.Group("/v1")

	// ルーティング
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}
	// デフォルトの8080から3000に変更
	g.Run(":3000")
}
