package conapicurd

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RowPostParams struct {
	Row map[string]string `form:"row"`
	Dt  string            `form:"dt"`
	Cmd string            `form:"cmd"`
}

func H_CurdPush(c *gin.Context) {
	pageinfo := RowPostParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
}
func H_CurdFetch(c *gin.Context) {

}
