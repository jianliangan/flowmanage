package daoflow

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

func CountRows() (int, error) {
	count, err := basecurd.CountRows("workflow_release", nil)
	return count, err
}
func FetchRow(flow_id string) (*dtoflow.FlowRow, error) {

	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "flow_name"},
		basecurd.FieldProperty{NeedDate: true, Name: "dtime"},
		basecurd.FieldProperty{NeedDate: false, Name: "description"},
	)
	rows := []*dtoflow.FlowRow{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.FlowRow{}
		ret := make([]any, 0)
		ret = append(ret, &m.Flow_id)
		ret = append(ret, &m.Flowname)
		ret = append(ret, &m.Description)
		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["flow_id"] = flow_id

	err := basecurd.SelectExe(cols, "workflow_release", where, 0, 1, &rows, fetchFlowAttr)

	var m *dtoflow.FlowRow
	if len(rows) > 0 {
		m = rows[0]
	}
	return m, err

}
func FetchRows(limitstart int, limitlen int) ([]*dtoflow.FlowRow, error) {

	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "flow_name"},
		basecurd.FieldProperty{NeedDate: true, Name: "dtime"},
		basecurd.FieldProperty{NeedDate: false, Name: "description"},
	)
	rows := []*dtoflow.FlowRow{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.FlowRow{}
		ret := make([]any, 0)
		ret = append(ret, &m.Flow_id)
		ret = append(ret, &m.Flowname)
		ret = append(ret, &m.Dtime)
		ret = append(ret, &m.Description)
		rows = append(rows, m)
		return ret
	}

	err := basecurd.SelectExe(cols, "workflow_release", nil, limitstart, limitlen, &rows, fetchFlowAttr)
	return rows, err

}
func flowConvert(col *dtoflow.FlowRow) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_name"},
			basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
			basecurd.FieldProperty{NeedDate: false, Name: "description"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "flow_id", Value: col.Flow_id},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_name", Value: col.Flowname},
			basecurd.FieldProperty{NeedDate: false, Name: "description", Value: col.Description},
		)
	}
	return cols
}
func AddRow(row *dtoflow.FlowRow) error {
	cols := flowConvert(row)
	err := basecurd.AddExe(cols, "workflow_release")
	return err
}

func EditRow(row *dtoflow.FlowRow) error {
	cols := flowConvert(row)
	where := map[string]any{}
	where["flow_id"] = row.Flow_id
	err := basecurd.EditExe(cols, "workflow_release", where)

	return err
}
func DelRow(row *dtoflow.FlowRow) error {
	where := map[string]any{}
	where["flow_id"] = row.Flow_id
	err := basecurd.DeleteExe("workflow_release", where)

	return err
}
