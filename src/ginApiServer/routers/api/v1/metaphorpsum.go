package v1

import (
	"api_server/src/ginApiServer/pkg/e"
	"api_server/src/ginApiServer/routers/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 取得 Metaphorpsum Sentence 內容
func GetSentence(c *gin.Context) {

	var (
		requestURL string = "http://metaphorpsum.com/sentences/3"
		err        error
	)

	sentence := service.NewDailySentence(requestURL)
	if sentence, err = sentence.GetSentence(); err != nil {
		// 400 回傳錯誤代碼
		c.JSON(http.StatusBadRequest, gin.H{
			"code": e.ERROR,
			"msg":  e.GetMsg(e.ERROR),
			"data": "",
		})
		c.Abort()

	} else {
		// 200 回傳成功代碼
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": sentence.RespData,
		})
	}

	c.Next()
}
