package conflowform

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jianliangan/flowmanage/dao/daoflow"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
	"github.com/jianliangan/flowmanage/model/modworkflow"

	"github.com/jianliangan/flowmanage/libs/tools/goldflake"
)

type simpleParams struct {
	Page int    `form:"page"`
	Id   string `form:"id"`
}
type simpleFormParams struct {
	Page    int    `form:"page"`
	Flow_id string `form:"flow_id"`
}
type simpleflowid_nodeid struct {
	Node_id string `form:"node_id"`
	Flow_id string `form:"flow_id"`
}
type formFieldPushParams struct {
	Flow_id string               `json:"flow_id"`
	Rows    []*dtoflow.FormField `json:"rows"`
}
type formFieldPushParamsRow struct {
	Flow_id string               `json:"flow_id"`
	Rows    []*dtoflow.FormField `json:"rows"`
	Cmd     string               `json:"cmd"`
}
type nodeInfoPushParams struct {
	Flow_id       string                                        `json:"flow_id"`
	Ui_btninfos   map[string][]*dtoflow.Btninfo                 `json:"ui_btninfos"`
	Ui_paths      map[string]map[string][]*dtoflow.Nodeinfo     `json:"ui_paths"`
	Ui_checkinfos map[string]map[string][]*dtoflow.Btncheckinfo `json:"ui_checkinfos"`
	Ui_btnokinfos map[string]map[string][]*dtoflow.Btncheckinfo `json:"ui_btnokinfos"`
	Ui_btnactinfs map[string]map[string][]*dtoflow.Btnactinfo   `json:"ui_btnactinfs"`
	Ui_nodes      map[string]*dtoflow.Nodeinfo                  `json:"ui_nodes"`
}

func H_GetListDev(c *gin.Context) {

	pageinfo := simpleParams{}
	c.Bind(&pageinfo)

	var wflist *modworkflow.FlowListInfos //[]*dtoflow.FlowRow
	var err error
	wflist, err = modworkflow.GetList(pageinfo.Page, pageinfo.Id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": wflist.Rows,
			"num":  wflist.Num,
			"size": wflist.Size,
		})
	}

}
func H_AddRowDev(c *gin.Context) {

	var row dtoflow.FlowRow
	c.ShouldBindBodyWith(&row, binding.JSON)
	id, _ := goldflake.Default().NextID()
	row.Flow_id = id
	err := daoflow.AddRow(&row)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "1",
		})
	}

}
func H_EditRowDev(c *gin.Context) {

	var row dtoflow.FlowRow
	c.ShouldBindBodyWith(&row, binding.JSON)

	err := daoflow.EditRow(&row)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "1",
		})
	}

}
func H_DelRowDev(c *gin.Context) {

	var row dtoflow.FlowRow
	c.ShouldBindBodyWith(&row, binding.JSON)

	err := daoflow.DelRow(&row)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "1",
		})
	}

}

/**
form
*/
func H_FecthFormField(c *gin.Context) {

	pageinfo := simpleFormParams{}
	c.Bind(&pageinfo)

	var wflist *modworkflow.FormFieldInfos //[]*dtoflow.FlowRow
	var err error
	wflist, err = modworkflow.GetFormFields(pageinfo.Flow_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"list": wflist.Rows,
		})
	}

}

func H_PushFormField(c *gin.Context) {

	pageinfo := formFieldPushParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)

	err := modworkflow.PushFormFields(pageinfo.Rows, pageinfo.Flow_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"obj": "1",
		})
	}

}
func H_PushFormFieldRow(c *gin.Context) {

	pageinfo := formFieldPushParamsRow{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)

	err := modworkflow.PushFormFieldsCmd(pageinfo.Rows, pageinfo.Flow_id, pageinfo.Cmd)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"obj": "1",
		})
	}

}

//node 信息推送
func H_PushNodeRowNew(c *gin.Context) {
	pageinfo := []*dtoflow.Nodeinfo{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushNodeNewCmd(pageinfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}

//node 信息推送
func H_PushButtonRowNew(c *gin.Context) {
	pageinfo := []*dtoflow.Btninfo{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushButtonNewCmd(pageinfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}

type PushPathParams struct {
	List   []*dtoflow.BtnPathinfo `json:"list"`
	Btn_id string                 `json:"btn_id"`
}

func H_PushPathRowNew(c *gin.Context) {
	pageinfo := PushPathParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushPathNewCmd(pageinfo.List, pageinfo.Btn_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}

}

type PushBtnCheckParams struct {
	List   []*dtoflow.Btncheckinfo `json:"list"`
	Btn_id string                  `json:"btn_id"`
}

func H_PushBtnCheckRowNew(c *gin.Context) {
	pageinfo := PushBtnCheckParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushBtnCheckNewCmd(pageinfo.List, pageinfo.Btn_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}
}

type PushBtnOkParams struct {
	List   []*dtoflow.Btncheckinfo `json:"list"`
	Btn_id string                  `json:"btn_id"`
}

func H_PushBtnOkRowNew(c *gin.Context) {
	pageinfo := PushBtnOkParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushBtnOkNewCmd(pageinfo.List, pageinfo.Btn_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}
}

type PushBtnActParams struct {
	List   []*dtoflow.Btnactinfo `json:"list"`
	Btn_id string                `json:"btn_id"`
}

func H_PushBtnActRowNew(c *gin.Context) {
	pageinfo := PushBtnActParams{}
	c.ShouldBindBodyWith(&pageinfo, binding.JSON)
	reid, err := modworkflow.PushBtnActNewCmd(pageinfo.List, pageinfo.Btn_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"reid": reid,
			"obj":  "1",
		})
	}
}

//node信息拉取
func H_FecthNodeInfos(c *gin.Context) {

	pageinfo := simpleFormParams{}
	c.Bind(&pageinfo)

	var wfobj *modworkflow.FlowNodeInfos //[]*dtoflow.FlowRow
	var err error
	wfobj, err = modworkflow.GetFlowNodeInfos(pageinfo.Flow_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"obj": wfobj,
		})
	}

}
func H_FecthNodeRowInfos(c *gin.Context) {

	pageinfo := simpleflowid_nodeid{}
	c.Bind(&pageinfo)

	var wfobj *modworkflow.FlowNodeResInfos //[]*dtoflow.FlowRow
	var err error
	wfobj, err = modworkflow.GetFlowNodeRowInfos(pageinfo.Flow_id, pageinfo.Node_id)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "发生错误" + err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"obj": wfobj,
		})
	}

}
