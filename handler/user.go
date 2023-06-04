package handler

import (
	"go-sample/model"
	"go-sample/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Create(c *gin.Context) {
	input := model.UserInput{}

	// ヘッダーのJSONをinputにバインド
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	// サービス層をNewする
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	// サービス層のCreateメソッドを呼び出す
	payload, err := user.Create(&input)
	if err != nil {
		log.Fatal(err)
	}

	// ステータス200と、payloadを返す
	c.JSON(http.StatusOK, payload)
}

func GetAll(c *gin.Context) {
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func GetOne(c *gin.Context) {
	// user-idのパラメータを数字に変換
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	payload, err := user.GetOne(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Update(c *gin.Context) {
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	input := model.UserInput{}
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()

	payload, err := user.Update(&input, userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, payload)
}

func Delete(c *gin.Context) {
	userID, err := strconv.Atoi((c.Param("user-id")))
	if err != nil {
		log.Fatal(err)
	}
	service := c.MustGet(service.ServiceKey).(service.Servicer)
	user := service.NewUser()
	err = user.Delete(userID)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, "削除されました")
}
