package daopath

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

func SelectPathRowsByBtnId(btn_id string) ([]*dtoflow.BtnPathinfo, error) {

	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
	)
	rows := []*dtoflow.BtnPathinfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.BtnPathinfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Node_id)
		ret = append(ret, &m.Flow_id)
		ret = append(ret, &m.Btn_id)

		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.SelectExe(cols, "wf_btn_path", where, 0, 1500, &rows, fetchFlowAttr)
	return rows, err

}
func DeletePathRowByBtnId(btn_id string) error {

	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.DeleteExe("wf_btn_path", where)
	return err
}
func DeletePathRowByNodeid(node_id string) error {

	where := map[string]any{}
	where["node_id"] = node_id
	err := basecurd.DeleteExe("wf_btn_path", where)
	return err
}
func pathConvert(col *dtoflow.BtnPathinfo) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id", Value: col.Btn_id},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id", Value: col.Node_id},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id", Value: col.Flow_id},
		)
	}
	return cols
}
func AddPathRow(row *dtoflow.BtnPathinfo, btn_id string) error {
	cols := pathConvert(row)
	err := basecurd.AddExe(cols, "wf_btn_path")
	return err
}
