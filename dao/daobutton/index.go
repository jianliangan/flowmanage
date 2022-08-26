package daobutton

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

func SelectBtnRowsByNodeId(node_id string) ([]*dtoflow.Btninfo, error) {

	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_name"},
		basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "sort"},
	)
	rows := []*dtoflow.Btninfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.Btninfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Btn_name)
		ret = append(ret, &m.Node_id)
		ret = append(ret, &m.Btn_id)
		ret = append(ret, &m.Sort)
		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["node_id"] = node_id
	err := basecurd.SelectExe(cols, "wfnode_buttons", nil, 0, 1500, &rows, fetchFlowAttr)
	return rows, err

}
func DeleteBtnRowBynodeid(node_id string) error {

	where := map[string]any{}
	where["node_id"] = node_id
	err := basecurd.DeleteExe("wfnode_buttons", where)
	return err
}
func DeleteBtnRowBybtnid(btn_id string) error {
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.DeleteExe("wfnode_buttons", where)
	return err
}
func btnConvert(col *dtoflow.Btninfo) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_name"},
			basecurd.FieldProperty{NeedDate: false, Name: "sort"},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_name", Value: col.Btn_name},
			basecurd.FieldProperty{NeedDate: false, Name: "sort", Value: col.Sort},
			basecurd.FieldProperty{NeedDate: false, Name: "node_id", Value: col.Node_id},
			basecurd.FieldProperty{NeedDate: false, Name: "btn_id", Value: col.Btn_id},
		)
	}
	return cols
}
func EditBtnRow(row *dtoflow.Btninfo, node_id string, btn_id string) error {
	cols := btnConvert(row)
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.EditExe(cols, "wfnode_buttons", where)
	return err
}
func AddBtnRow(row *dtoflow.Btninfo, node_id string, btn_id string) error {
	cols := btnConvert(row)
	err := basecurd.AddExe(cols, "wfnode_buttons")
	return err
}
