package concommon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianliangan/flowmanage/libs/tools/goldflake"
)

type SimpleParams struct {
	Page int    `form:"page"`
	Id   string `form:"id"`
}

func H_GetGlobalId(c *gin.Context) {
	//res := []string{}
	//for i := 0; i < 10; i++ {
	//	id, _ := goldflake.Default().NextID()
	//	res = append(res, id)
	//}
	id, _ := goldflake.Default().NextID()
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
