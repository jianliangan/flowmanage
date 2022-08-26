package conlogin

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func H_Menuinfo(c *gin.Context) {
	//relist := make(map[string]any)
	tmp := `{"menu":[{"name":"home","path":"/home","meta":{"title":"项目看板","icon":"el-icon-eleme-filled","type":"menu"},"children":[{"name":"dashboard","path":"/dashboard","meta":{"title":"控制台","icon":"el-icon-menu","affix":true},"component":"home"}]},{"name":"userCenter","path":"/usercenter","meta":{"title":"帐号信息","icon":"el-icon-user","tag":"NEW"},"component":"userCenter"},{"name":"projectmanage","path":"/projectmanage","meta":{"title":"项目管理","icon":"el-icon-takeaway-box","type":"menu"},"children":[{"path":"/flowlist","name":"flowlist","meta":{"title":"项目列表","icon":"el-icon-magic-stick","type":"menu"},"component":"workflow/WorkFlowList"},{"path":"/flowedit","name":"flowedit","meta":{"title":"流程编辑","icon":"el-icon-orange","type":"menu","hidden":true},"component":"workflow/WorkFlowEdit"}]},{"name":"setting","path":"/setting","meta":{"title":"设置","icon":"el-icon-setting","type":"menu"},"children":[{"path":"/flowlist","name":"flowlist","meta":{"title":"流程设置","icon":"el-icon-magic-stick","type":"menu"},"component":"workflow/WorkFlowList"},{"path":"/flowedit","name":"flowedit","meta":{"title":"流程编辑","icon":"el-icon-orange","type":"menu","hidden":true},"component":"workflow/WorkFlowEdit"},{"path":"/setting/user","name":"user","meta":{"title":"用户管理","icon":"el-icon-user-filled","type":"menu"},"component":"setting/user"},{"path":"/setting/role","name":"role","meta":{"title":"角色管理","icon":"el-icon-notebook","type":"menu"},"component":"setting/role"},{"path":"/setting/organize","name":"dept","meta":{"title":"部门管理","icon":"sc-icon-organization","type":"menu"},"component":"setting/organize"},{"path":"/setting/business","name":"business","meta":{"title":"业务架构","icon":"el-icon-scale-to-original","type":"menu"},"component":"setting/business"},{"path":"/setting/dic","name":"dic","meta":{"title":"字典管理","icon":"el-icon-document","type":"menu"},"component":"setting/dic"},{"path":"/setting/table","name":"tableSetting","meta":{"title":"表格列管理","icon":"el-icon-scale-to-original","type":"menu"},"component":"setting/table"},{"path":"/setting/menu","name":"settingMenu","meta":{"title":"菜单管理","icon":"el-icon-fold","type":"menu"},"component":"setting/menu"},{"path":"/setting/task","name":"task","meta":{"title":"计划任务","icon":"el-icon-alarm-clock","type":"menu"},"component":"setting/task"},{"path":"/setting/client","name":"client","meta":{"title":"应用管理","icon":"el-icon-help-filled","type":"menu"},"component":"setting/client"},{"path":"/setting/log","name":"log","meta":{"title":"系统日志","icon":"el-icon-warning","type":"menu"},"component":"setting/log"}]},{"path":"/other/about","name":"about","meta":{"title":"关于","icon":"el-icon-info-filled","type":"menu"},"component":"other/about"}],"permissions":["list.add","list.edit","list.delete","user.add","user.edit","user.delete"]}`

	tmpmap := map[string]interface{}{}
	json.Unmarshal([]byte(tmp), &tmpmap)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "",
		"data":    tmpmap,
	})
}

func H_Token(c *gin.Context) {
	//relist := make(map[string]any)
	//scui 用了scui的layout和登录注册部分
	tmp := `{"token":"SCUI.Administrator.Auth","userInfo":{"userId":"1","userName":"Administrator","dashboard":"0","role":["SA","admin","Auditor"]}}`
	tmpmap := map[string]interface{}{}
	json.Unmarshal([]byte(tmp), &tmpmap)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "",
		"data":    tmpmap,
	})
}
func H_Ver(c *gin.Context) {
	//relist := make(map[string]any)

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "",
		"data":    "1.6.6",
	})
}
