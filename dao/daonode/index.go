package daonode

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dao/daobutton"
	"github.com/jianliangan/flowmanage/dao/daocheckokact"
	"github.com/jianliangan/flowmanage/dao/daopath"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

func SelectNodeRowsByFlowId(flow_id string) ([]*dtoflow.Nodeinfo, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "node_name"},
		basecurd.FieldProperty{NeedDate: false, Name: "timeout"},
		basecurd.FieldProperty{NeedDate: false, Name: "vx"},
		basecurd.FieldProperty{NeedDate: false, Name: "vy"},
	)
	rows := []*dtoflow.Nodeinfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.Nodeinfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Node_id)
		ret = append(ret, &m.Flow_id)
		ret = append(ret, &m.Node_name)
		ret = append(ret, &m.Timeout)
		ret = append(ret, &m.Vx)
		ret = append(ret, &m.Vy)

		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["flow_id"] = flow_id
	err := basecurd.SelectExe(cols, "wf_nodes", where, 0, 1500, &rows, fetchFlowAttr)
	return rows, err

}

func DeleteNodeRowByNodeId(node_id string) error {
	btnrows, _ := daobutton.SelectBtnRowsByNodeId(node_id)
	//删除按钮
	daobutton.DeleteBtnRowBynodeid(node_id)
	//删除path
	daopath.DeletePathRowByNodeid(node_id)
	//删除判断条件
	//删除生效条件
	//删除回调
	for _, value := range btnrows {
		daocheckokact.DeleteCheckRowByBtnid(value.Btn_id)
		daocheckokact.DeleteBtnActRowByBtnid(value.Btn_id)
		daocheckokact.DeleteBtnActRowByBtnid(value.Btn_id)
		daopath.DeletePathRowByBtnId(value.Btn_id)
	}

	/////////
	where := map[string]any{}
	where["node_id"] = node_id
	err := basecurd.DeleteExe("wf_nodes", where)

	return err
}
func nodeConvert(col *dtoflow.Nodeinfo) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "node_name"},
			basecurd.FieldProperty{NeedDate: false, Name: "vx"},
			basecurd.FieldProperty{NeedDate: false, Name: "vy"},
			basecurd.FieldProperty{NeedDate: false, Name: "timeout"},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "node_name", Value: col.Node_name},
			basecurd.FieldProperty{NeedDate: false, Name: "vx", Value: col.Vx},
			basecurd.FieldProperty{NeedDate: false, Name: "vy", Value: col.Vy},
			basecurd.FieldProperty{NeedDate: false, Name: "timeout", Value: col.Timeout},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id", Value: col.Flow_id},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id", Value: col.Node_id},
		)
	}
	return cols
}
func EditNodeRow(row *dtoflow.Nodeinfo, flow_id string, node_id string) error {

	cols := nodeConvert(row)
	where := map[string]any{}
	where["flow_id"] = flow_id
	where["node_id"] = node_id
	err := basecurd.EditExe(cols, "wf_nodes", where)

	return err
}
func AddNodeRow(row *dtoflow.Nodeinfo, flow_id string, node_id string) error {
	cols := nodeConvert(row)
	err := basecurd.AddExe(cols, "wf_nodes")
	return err
}
