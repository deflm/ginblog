package v1

import (
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Upload(c *gin.Context) {
	var code int
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "img/"+file.Filename)
	if err != nil {
		code = errmsg.ERROR
		log.Fatalln(err)
	}
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
