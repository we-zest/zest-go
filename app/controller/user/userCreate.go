package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zest-go/app/model"
)

func Create(context *gin.Context) {
	var user model.User
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.JSON(http.StatusBadRequest, "输入的数据不合法")
	}
	user.Create()
	context.JSON(http.StatusOK, "用户创建成功")
}
