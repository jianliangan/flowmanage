package consyssetting

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jianliangan/flowmanage/control/concommon"
	"github.com/jianliangan/flowmanage/dto/dtosetting"
	"github.com/jianliangan/flowmanage/model/modcontext"
	"github.com/jianliangan/flowmanage/model/modsyssetting"
)

func treeOrganizeParse(d0 *dtosetting.Organize_str, wflist []*dtosetting.Organize_str) *dtosetting.Organize_str {

	list := []*dtosetting.Organize_str{}

	for i := 0; i < len(wflist); i++ {
		value := (wflist)[i]
		if value.Parentid == d0.Organize_str_id {

			wflist = append((wflist)[:i], (wflist)[(i+1):]...)
			list = append(list, treeOrganizeParse(value, wflist))
			i = -1
		}
	}
	d0.Children = list
	return d0
}
func treeBusinessParse(d0 *dtosetting.Business_str, wflist []*dtosetting.Business_str) *dtosetting.Business_str {

	list := []*dtosetting.Business_str{}

	for i := 0; i < len(wflist); i++ {
		value := (wflist)[i]
		if value.Parentid == d0.Business_str_id {

			wflist = append((wflist)[:i], (wflist)[(i+1):]...)
			list = append(list, treeBusinessParse(value, wflist))
			i = -1
		}
	}
	d0.Children = list
	return d0
}

//node信息拉取
func H_FecthBusiness(c *gin.Context) {

	pageinfo := concommon.SimpleParams{}
	c.Bind(&pageinfo)

	wflist, err := modsyssetting.GetBusinessList()
	relist := []*dtosetting.Business_str{}
	for i := 0; i < len(wflist); i++ {
		value := (wflist)[i]
		if value.Parentid == "" {
			wflist = append((wflist)[:i], (wflist)[(i+1):]...)
			relist = append(relist, treeBusinessParse(value, wflist))
			i = -1
		}

	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": relist,
			"num":  len(relist),
			"size": 500,
		})
	}

}
func H_PushBusiness(c *gin.Context) {
	pageinfo := dtosetting.Business_str{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modsyssetting.PushBusinessNewCmd(&pageinfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}

func H_FecthOrganize(c *gin.Context) {

	pageinfo := concommon.SimpleParams{}
	c.Bind(&pageinfo)

	wflist, err := modsyssetting.GetOrganizeList()
	relist := []*dtosetting.Organize_str{}
	for i := 0; i < len(wflist); i++ {
		value := (wflist)[i]
		if value.Parentid == "" {
			wflist = append((wflist)[:i], (wflist)[(i+1):]...)
			relist = append(relist, treeOrganizeParse(value, wflist))
			i = -1
		}

	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": relist,
			"num":  len(relist),
			"size": 500,
		})
	}

}
func H_PushOrganize(c *gin.Context) {
	pageinfo := dtosetting.Organize_str{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modsyssetting.PushOrganizeNewCmd(&pageinfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}

//user
func H_FecthUser(c *gin.Context) {

	pageinfo := concommon.SimpleParams{}

	wflist, err := modsyssetting.GetUserList(pageinfo.Page, pageinfo.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": wflist,
			"num":  wflist.Num,
			"size": modcontext.Default().Pagesize,
		})
	}

}

//role
func H_FecthRole(c *gin.Context) {
	wflist, err := modsyssetting.GetRoleList()
	num, _ := modsyssetting.GetRoleNum()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": wflist,
			"num":  num,
			"size": modcontext.Default().Pagesize,
		})
	}

}
func H_PushRole(c *gin.Context) {
	pageinfo := dtosetting.Role_str{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modsyssetting.PushRoleNewCmd(&pageinfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}
