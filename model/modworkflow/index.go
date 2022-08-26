package modworkflow

import (
	"github.com/jianliangan/flowmanage/dao/daobutton"
	"github.com/jianliangan/flowmanage/dao/daocheckokact"
	"github.com/jianliangan/flowmanage/dao/daoflow"
	"github.com/jianliangan/flowmanage/dao/daoform"
	"github.com/jianliangan/flowmanage/dao/daonode"
	"github.com/jianliangan/flowmanage/dao/daopath"

	"github.com/jianliangan/flowmanage/dto/dtoflow"
	"github.com/jianliangan/flowmanage/dto/dtoothers"
	"github.com/jianliangan/flowmanage/libs/tools/goldflake"
	"github.com/jianliangan/flowmanage/model/modcontext"
)

type FlowListInfos struct {
	Num  int                `json:"Num"`
	Rows []*dtoflow.FlowRow `json:"Rows"`
	Size int                `json:"Size"`
}

type FormFieldInfos struct {
	Rows []*dtoflow.FormField `json:"Rows"`
}
type FlowNodeInfos struct {
	NodeRows []*dtoflow.Nodeinfo               `json:"NodeRows"`
	PathRows map[string][]*dtoflow.BtnPathinfo `json:"PathRows"`
}
type FlowNodeResInfos struct {
	Btnconfoptions []*dtoothers.Defaultbutton         `json:"Btnconfoptions"`
	Fieldinfo      map[string]*dtoflow.FormField      `json:"Fieldinfo"`
	Btninfo        map[string]*dtoflow.Btninfo        `json:"Btninfo"`
	Btnpathinfo    map[string][]*dtoflow.BtnPathinfo  `json:"Btnpathinfo"`
	Btncheckinfo   map[string][]*dtoflow.Btncheckinfo `json:"Btncheckinfo"`
	Btnokinfo      map[string][]*dtoflow.Btncheckinfo `json:"Btnokinfo"`
	Btnactinfo     map[string][]*dtoflow.Btnactinfo   `json:"Btnactinfo"`
}
type FlowALLUpdateInfos struct {
	Ui_btninfos    map[string][]*dtoflow.Btninfo                 `json:"Ui_btninfos"`
	Ui_paths       map[string]map[string][]*dtoflow.Nodeinfo     `json:"Ui_paths"`
	Ui_checkinfos  map[string]map[string][]*dtoflow.Btncheckinfo `json:"Ui_checkinfos"`
	Ui_btnokinfos  map[string]map[string][]*dtoflow.Btncheckinfo `json:"Ui_btnokinfos"`
	Ui_btnactinfos map[string]map[string][]*dtoflow.Btnactinfo   `json:"Ui_btnactinfos"`
	Ui_nodes       map[string]*dtoflow.Nodeinfo                  `json:"Ui_nodes"`
}

func GetList(page int, id string) (*FlowListInfos, error) {

	flowListInfos := FlowListInfos{}
	if page == 0 {
		page = 1
	}
	if id != "" {
		fetchRow, _ := daoflow.FetchRow(id)
		if fetchRow != nil {
			flowListInfos.Rows = append(flowListInfos.Rows, fetchRow)
		}

	} else {
		limitlen := modcontext.Default().Pagesize
		limitstart := (page - 1) * limitlen

		fetchRows, _ := daoflow.FetchRows(limitstart, limitlen)
		num, _ := daoflow.CountRows()
		flowListInfos.Rows = fetchRows
		flowListInfos.Num = num
		flowListInfos.Size = limitlen
	}

	return &flowListInfos, nil

}

//form
func GetFormFields(id string) (*FormFieldInfos, error) {
	formFieldInfos := FormFieldInfos{}
	fetchFormRows, err := daoform.FetchFormRows(id)
	if err == nil {
		formFieldInfos.Rows = fetchFormRows
	}
	return &formFieldInfos, nil

}
func PushFormFieldsCmd(rows []*dtoflow.FormField, formid string, cmd string) error {
	var err error
	for _, row := range rows {
		if cmd == "edit" {
			err = daoform.EditFormRow(row, formid)
		} else if cmd == "delete" {
			err = daoform.DelFormByField(row.Field_id, formid)
		} else if cmd == "add" {
			row.Field_id, _ = goldflake.Default().NextID()
			err = daoform.AddFormRow(row, formid)
		}
		if err != nil {
			break
		}
	}

	return err
}
func PushFormFields(rows []*dtoflow.FormField, formid string) error {

	formfieldinfos, err := GetFormFields(formid)
	for _, v1 := range rows {
		isup := false
		for _, v2 := range formfieldinfos.Rows {
			if v2.Field_id == v1.Field_id {
				isup = true
				break
			}
		}
		if isup {
			daoform.EditFormRow(v1, formid)
		} else {
			id, _ := goldflake.Default().NextID()
			v1.Field_id = id
			daoform.AddFormRow(v1, formid)
		}
	}

	for _, v1 := range formfieldinfos.Rows {
		isdel := true
		for _, v2 := range rows {
			if v2.Field_id == v1.Field_id {
				isdel = false
				break
			}
		}
		if isdel {
			daoform.DelFormByField(v1.Field_id, formid)
		}
	}
	return err
}

/*
	Ui_btninfos   map[string][]*dtoflow.Btninfo                 `json:"ui_btninfos"`
	Ui_paths      map[string]map[string][]*dtoflow.Nodeinfo     `json:"ui_paths"`
	Ui_checkinfs  map[string]map[string][]*dtoflow.Btncheckinfo `json:"ui_checkinfs"`
	Ui_btnokinfs  map[string]map[string][]*dtoflow.Btncheckinfo `json:"ui_btnokinfs"`
	Ui_btnactinfs map[string]map[string][]*dtoflow.Btnactinfo   `json:"ui_btnactinfs"`
	Ui_nodes      map[string]*nodeinfoParams                 `json:"ui_nodes"`
*/
//node all info

func PushNodeNewCmd(rows []*dtoflow.Nodeinfo) (string, error) {
	var err error
	reid := "0"
	for _, row := range rows {
		if row.Cmd_ == "edit" {
			err = daonode.EditNodeRow(row, row.Flow_id, row.Node_id)
		} else if row.Cmd_ == "delete" {
			err = daonode.DeleteNodeRowByNodeId(row.Node_id)
		} else if row.Cmd_ == "add" {
			row.Node_id, _ = goldflake.Default().NextID()
			reid = row.Node_id
			err = daonode.AddNodeRow(row, row.Flow_id, row.Node_id)
		}
		if err != nil {
			break
		}
	}

	return reid, err
}
func PushButtonNewCmd(rows []*dtoflow.Btninfo) (string, error) {
	var err error
	reid := "0"
	for _, row := range rows {
		if row.Cmd_ == "edit" {
			err = daobutton.EditBtnRow(row, row.Node_id, row.Btn_id)
		} else if row.Cmd_ == "delete" {
			err = daobutton.DeleteBtnRowBybtnid(row.Btn_id)
		} else if row.Cmd_ == "add" {
			row.Btn_id, _ = goldflake.Default().NextID()
			reid = row.Btn_id
			err = daobutton.AddBtnRow(row, row.Node_id, row.Btn_id)
		}
		if err != nil {
			break
		}
	}

	return reid, err
}
func PushPathNewCmd(rows []*dtoflow.BtnPathinfo, btn_id string) (string, error) {
	var err error
	reid := "0"
	daopath.DeletePathRowByBtnId(btn_id)
	for _, row := range rows {
		err = daopath.AddPathRow(row, btn_id)
		if err != nil {
			break
		}
	}

	return reid, err
}
func PushBtnCheckNewCmd(rows []*dtoflow.Btncheckinfo, btn_id string) (string, error) {
	var err error
	reid := "0"
	daocheckokact.DeleteCheckRowByBtnid(btn_id)
	for _, row := range rows {
		err = daocheckokact.AddCheckRow(row, btn_id)
		if err != nil {
			break
		}
	}

	return reid, err
}
func PushBtnOkNewCmd(rows []*dtoflow.Btncheckinfo, btn_id string) (string, error) {
	var err error
	reid := "0"
	daocheckokact.DeleteBtnOkRowBybtnid(btn_id)
	for _, row := range rows {
		err = daocheckokact.AddBtnOkRow(row, btn_id)
		if err != nil {
			break
		}
	}

	return reid, err
}
func PushBtnActNewCmd(rows []*dtoflow.Btnactinfo, btn_id string) (string, error) {
	var err error
	reid := "0"
	daocheckokact.DeleteBtnActRowByBtnid(btn_id)
	for _, row := range rows {
		err = daocheckokact.AddBtnActRow(row, btn_id)
		if err != nil {
			break
		}
	}

	return reid, err
}

//拿到所有节点信息
func GetFlowNodeInfos(flow_id string) (*FlowNodeInfos, error) {
	//拿节点信息
	flowNodeInfos := FlowNodeInfos{}
	nodeinfos, err := daonode.SelectNodeRowsByFlowId(flow_id)
	if err != nil {
		return nil, err
	}
	flowNodeInfos.NodeRows = nodeinfos

	//拿路径信息
	if flowNodeInfos.PathRows == nil {
		flowNodeInfos.PathRows = make(map[string][]*dtoflow.BtnPathinfo)
	}
	for _, nodeinfosvalue := range nodeinfos {
		nodeid := nodeinfosvalue.Node_id
		btnrows, err := daobutton.SelectBtnRowsByNodeId(nodeid)
		if err != nil {
			continue
		}
		btnPathinfo := []*dtoflow.BtnPathinfo{}
		for _, btnrowsvalue := range btnrows {
			btn_id := btnrowsvalue.Btn_id
			pathrows, err := daopath.SelectPathRowsByBtnId(btn_id)
			if err != nil {
				continue
			}
			btnPathinfo = append(btnPathinfo, pathrows...)
		}
		if len(btnPathinfo) > 0 {
			flowNodeInfos.PathRows[nodeid] = btnPathinfo
		}

	}
	return &flowNodeInfos, nil
}

//拿到节点信息
func GetFlowNodeRowInfos(flow_id string, node_id string) (*FlowNodeResInfos, error) {
	//拿节点信息
	flowNodeResInfos := FlowNodeResInfos{}
	//Fieldinfo
	flowNodeResInfos.Fieldinfo = make(map[string]*dtoflow.FormField)
	formFields, err := daoform.FetchFormRows(flow_id)
	if err != nil {
		return nil, err
	}
	for _, formfield := range formFields {
		flowNodeResInfos.Fieldinfo[formfield.Field_id] = formfield
	}
	//Btninfo
	flowNodeResInfos.Btninfo = make(map[string]*dtoflow.Btninfo)
	formbtninfo, err := daobutton.SelectBtnRowsByNodeId(node_id)
	if err != nil {
		return nil, err
	}
	for _, btninfo := range formbtninfo {
		flowNodeResInfos.Btninfo[btninfo.Btn_id] = btninfo
	}
	//Btnpathinfo
	flowNodeResInfos.Btnpathinfo = make(map[string][]*dtoflow.BtnPathinfo)
	for _, btninfovalue := range flowNodeResInfos.Btninfo {
		btnid := btninfovalue.Btn_id
		btnpathinfo, err := daopath.SelectPathRowsByBtnId(btnid)
		if err != nil {
			continue
		}
		if len(btnpathinfo) > 0 {
			flowNodeResInfos.Btnpathinfo[btnid] = btnpathinfo
		}
	}
	//Btncheckinfo
	flowNodeResInfos.Btncheckinfo = make(map[string][]*dtoflow.Btncheckinfo)
	for _, btninfovalue := range flowNodeResInfos.Btninfo {
		btnid := btninfovalue.Btn_id
		btncheckinfo, err := daocheckokact.SelectBtnChecksByBtnId(btnid)
		if err != nil {
			continue
		}
		if len(btncheckinfo) > 0 {
			flowNodeResInfos.Btncheckinfo[btnid] = btncheckinfo
		}
	}
	//BtnOkinfo
	flowNodeResInfos.Btnokinfo = make(map[string][]*dtoflow.Btncheckinfo)
	for _, btninfovalue := range flowNodeResInfos.Btninfo {
		btnid := btninfovalue.Btn_id
		btnokinfo, err := daocheckokact.SelectBtnOkByBtnId(btnid)
		if err != nil {
			continue
		}
		if len(btnokinfo) > 0 {
			flowNodeResInfos.Btnokinfo[btnid] = btnokinfo
		}
	}
	//BtnOkinfo
	flowNodeResInfos.Btnactinfo = make(map[string][]*dtoflow.Btnactinfo)
	for _, btninfovalue := range flowNodeResInfos.Btninfo {
		btnid := btninfovalue.Btn_id

		btnactinfo, err := daocheckokact.SelectBtnActByBtnId(btnid)

		if err != nil {
			continue
		}

		if len(btnactinfo) > 0 {
			flowNodeResInfos.Btnactinfo[btnid] = btnactinfo
		}
	}

	return &flowNodeResInfos, nil
}
