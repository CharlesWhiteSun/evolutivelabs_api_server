package v1

import (
	"api_server/src/ginApiServer/pkg/e"
	"api_server/src/ginApiServer/routers/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 取得 ItsThisForThat Sentence 內容
func GetItsThisForThatSentence(c *gin.Context) {

	var (
		requestURL string = "https://itsthisforthat.com/api.php?text=" + c.Param("param")
		err        error
	)

	sentence := service.NewItsThisForThatDailySentence(requestURL)
	if sentence, err = sentence.GetItsThisForThatSentence(); err != nil {
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
