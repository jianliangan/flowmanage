package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jianliangan/flowmanage/control/concommon"
	"github.com/jianliangan/flowmanage/control/conflowform"
	"github.com/jianliangan/flowmanage/control/conlogin"
	"github.com/jianliangan/flowmanage/control/consyssetting"
	"github.com/jianliangan/flowmanage/model/modcontext"

	"github.com/jianliangan/flowmanage/libs/libdb"
	"github.com/jianliangan/flowmanage/libs/tools/goldflake"
	"github.com/jianliangan/flowmanage/libs/tools/jrlog"
	"github.com/jianliangan/flowmanage/libs/tools/unit"
)

func init() {
	//init log
	err := jrlog.Loginit(true)
	if err != nil {
		log.Fatal("[ERR]", err)
	}
	//init config
	err = modcontext.Coninit()
	if err != nil {
		jrlog.Logfatal("[ERR]", "conf.json:", err)
	}

	//int mysql
	dbenv := libdb.Mysql{}
	dbenv.Dburi = modcontext.Default().Dburi
	dbenv.Maxidle = modcontext.Default().Maxidle
	dbenv.Maxopen = modcontext.Default().Maxopen

	err = libdb.NewMySQL(&dbenv)
	if err != nil {
		jrlog.Logfatal("[ERR]", "mysql:", err)
	}
	//
	var st goldflake.Settings
	st.MachineID = func() (uint16, error) {
		return 1, nil
	}
	err = goldflake.NewGoldflake(st)
	if err != nil {
		jrlog.Logfatal("sonyflake not created")
	}
}
func main() {

	r := gin.Default()
	r.Use(unit.Core())

	r.POST("/api/demo/ver", conlogin.H_Ver)
	r.GET("/api/demo/ver", conlogin.H_Ver)
	r.POST("/api/token", conlogin.H_Token)
	r.GET("/api/token", conlogin.H_Token)

	r.POST("/system/menuinfo", conlogin.H_Menuinfo)
	r.GET("/system/menuinfo", conlogin.H_Menuinfo)

	r.GET("/this/global/getid", concommon.H_GetGlobalId)
	r.GET("/flow/f", conflowform.H_GetListDev)
	r.POST("/flow/a", conflowform.H_AddRowDev)
	r.POST("/flow/e", conflowform.H_EditRowDev)
	r.POST("/flow/d", conflowform.H_DelRowDev)
	r.GET("/flow/form/f", conflowform.H_FecthFormField)
	r.POST("/flow/form/p", conflowform.H_PushFormField)
	r.POST("/flow/formfield/p", conflowform.H_PushFormFieldRow)
	r.POST("/flow/node/p", conflowform.H_PushNodeRowNew)
	r.POST("/flow/button/p", conflowform.H_PushButtonRowNew)
	r.POST("/flow/path/p", conflowform.H_PushPathRowNew)
	r.POST("/flow/btncheck/p", conflowform.H_PushBtnCheckRowNew)
	r.POST("/flow/btnok/p", conflowform.H_PushBtnOkRowNew)
	r.POST("/flow/btnact/p", conflowform.H_PushBtnActRowNew)
	r.GET("/flow/node/f", conflowform.H_FecthNodeInfos)
	r.GET("/flow/noderow/f", conflowform.H_FecthNodeRowInfos)

	//业务架构和组织架构管理
	r.GET("/sys/business/f", consyssetting.H_FecthBusiness)
	r.POST("/sys/business/p", consyssetting.H_PushBusiness)
	r.GET("/sys/organize/f", consyssetting.H_FecthOrganize)
	r.POST("/sys/organize/p", consyssetting.H_PushOrganize)
	//用户
	r.GET("/sys/user/f", consyssetting.H_FecthUser)
	//角色
	r.GET("/sys/role/f", consyssetting.H_FecthRole)
	r.POST("/sys/role/p", consyssetting.H_PushRole)
	r.Run("0.0.0.0:82") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
