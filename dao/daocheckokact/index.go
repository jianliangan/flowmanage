package daocheckokact

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

func DeleteCheckRowByBtnid(btn_id string) error {

	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.DeleteExe("wfbutton_check", where)
	return err
}

func DeleteBtnOkRowBybtnid(btn_id string) error {
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.DeleteExe("wf_button_ok", where)
	return err

}
func DeleteBtnActRowByBtnid(btn_id string) error {
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.DeleteExe("wf_buttons_acts", where)
	return err
}
func DeleteCheckRowByFieldid(field_id string) error {

	where := map[string]any{}
	where["field_id"] = field_id
	err := basecurd.DeleteExe("wfbutton_check", where)
	return err

}

func DeleteBtnOkRowByFieldid(field_id string) error {
	where := map[string]any{}
	where["field_id"] = field_id
	err := basecurd.DeleteExe("wf_button_ok", where)
	return err
}
func checkConvert(col *dtoflow.Btncheckinfo) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "field_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "type"},
			basecurd.FieldProperty{NeedDate: false, Name: "value"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id", Value: col.Btn_id},
			basecurd.FieldProperty{NeedDate: false, Name: "field_id", Value: col.Field_id},
			basecurd.FieldProperty{NeedDate: false, Name: "type", Value: col.Type},
			basecurd.FieldProperty{NeedDate: false, Name: "value", Value: col.Value},
		)
	}
	return cols
}
func AddCheckRow(row *dtoflow.Btncheckinfo, btn_id string) error {

	cols := checkConvert(row)
	err := basecurd.AddExe(cols, "wfbutton_check")
	return err
}

func AddBtnOkRow(row *dtoflow.Btncheckinfo, btn_id string) error {

	cols := checkConvert(row)
	err := basecurd.AddExe(cols, "wf_button_ok")
	return err

}
func actConvert(col *dtoflow.Btnactinfo) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "body"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id", Value: col.Btn_id},
			basecurd.FieldProperty{NeedDate: false, Name: "body", Value: col.Body},
		)
	}
	return cols
}
func AddBtnActRow(row *dtoflow.Btnactinfo, btn_id string) error {

	cols := actConvert(row)
	err := basecurd.AddExe(cols, "wf_buttons_acts")
	return err
}

//check
func SelectBtnChecksByBtnId(btn_id string) ([]*dtoflow.Btncheckinfo, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "field_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "type"},
		basecurd.FieldProperty{NeedDate: false, Name: "value"},
	)
	rows := []*dtoflow.Btncheckinfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.Btncheckinfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Btn_id)
		ret = append(ret, &m.Field_id)
		ret = append(ret, &m.Type)
		ret = append(ret, &m.Value)
		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.SelectExe(cols, "wfbutton_check", where, 0, 100, &rows, fetchFlowAttr)
	return rows, err
}

//ok
func SelectBtnOkByBtnId(btn_id string) ([]*dtoflow.Btncheckinfo, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "field_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "type"},
		basecurd.FieldProperty{NeedDate: false, Name: "value"},
	)
	rows := []*dtoflow.Btncheckinfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.Btncheckinfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Btn_id)
		ret = append(ret, &m.Field_id)
		ret = append(ret, &m.Type)
		ret = append(ret, &m.Value)
		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.SelectExe(cols, "wf_button_ok", where, 0, 100, &rows, fetchFlowAttr)
	return rows, err

}

//act
func SelectBtnActByBtnId(btn_id string) ([]*dtoflow.Btnactinfo, error) {

	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "btn_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "body"},
	)
	rows := []*dtoflow.Btnactinfo{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.Btnactinfo{}
		ret := make([]any, 0)
		ret = append(ret, &m.Btn_id)
		ret = append(ret, &m.Body)
		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["btn_id"] = btn_id
	err := basecurd.SelectExe(cols, "wf_buttons_acts", where, 0, 100, &rows, fetchFlowAttr)
	return rows, err

}
