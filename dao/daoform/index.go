package daoform

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dao/daocheckokact"
	"github.com/jianliangan/flowmanage/dto/dtoflow"
)

var FieldMap = make(map[string]string)

func Init() {
	FieldMap["sort"] = ""
	FieldMap["field_id"] = ""
	FieldMap["flow_id"] = ""
	FieldMap["fieldname"] = ""
	FieldMap["fieldtype"] = ""
	FieldMap["tablestruct"] = ""
	FieldMap["apistr"] = ""
	FieldMap["onlyapi"] = ""
}

func FetchFormRows(flow_id string) ([]*dtoflow.FormField, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "sort"},
		basecurd.FieldProperty{NeedDate: false, Name: "field_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "fieldname"},
		basecurd.FieldProperty{NeedDate: false, Name: "fieldtype"},
		basecurd.FieldProperty{NeedDate: false, Name: "tablestruct"},
		basecurd.FieldProperty{NeedDate: false, Name: "apistr"},
		basecurd.FieldProperty{NeedDate: false, Name: "onlyapi"},
	)
	rows := []*dtoflow.FormField{}
	fetchFlowAttr := func() []any {
		m := &dtoflow.FormField{}
		ret := make([]any, 0)
		ret = append(ret, &m.Sort)
		ret = append(ret, &m.Field_id)
		ret = append(ret, &m.Flow_id)
		ret = append(ret, &m.Fieldname)
		ret = append(ret, &m.Fieldtype)
		ret = append(ret, &m.Tablestruct)
		ret = append(ret, &m.Apistr)
		ret = append(ret, &m.Onlyapi)

		rows = append(rows, m)
		return ret
	}
	where := map[string]any{}
	where["flow_id"] = flow_id
	err := basecurd.SelectExe(cols, "wf_form", where, 0, 150, &rows, fetchFlowAttr)

	return rows, err
}
func fieldConvert(col *dtoflow.FormField) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "sort"},
			basecurd.FieldProperty{NeedDate: false, Name: "field_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "fieldname"},
			basecurd.FieldProperty{NeedDate: false, Name: "fieldtype"},
			basecurd.FieldProperty{NeedDate: false, Name: "tablestruct"},
			basecurd.FieldProperty{NeedDate: false, Name: "apistr"},
			basecurd.FieldProperty{NeedDate: false, Name: "onlyapi"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "sort", Value: col.Sort},
			basecurd.FieldProperty{NeedDate: false, Name: "field_id", Value: col.Field_id},
			basecurd.FieldProperty{NeedDate: false, Name: "flow_id", Value: col.Flow_id},
			basecurd.FieldProperty{NeedDate: false, Name: "fieldname", Value: col.Fieldname},
			basecurd.FieldProperty{NeedDate: false, Name: "fieldtype", Value: col.Fieldtype},
			basecurd.FieldProperty{NeedDate: false, Name: "tablestruct", Value: col.Tablestruct},
			basecurd.FieldProperty{NeedDate: false, Name: "apistr", Value: col.Apistr},
			basecurd.FieldProperty{NeedDate: false, Name: "onlyapi", Value: col.Onlyapi},
		)
	}
	return cols
}
func AddFormRow(row *dtoflow.FormField, formid string) error {

	cols := fieldConvert(row)
	err := basecurd.AddExe(cols, "wf_form")
	return err

}
func EditFormRow(row *dtoflow.FormField, formid string) error {

	cols := fieldConvert(row)
	where := map[string]any{}
	where["flow_id"] = row.Flow_id
	where["field_id"] = row.Field_id
	err := basecurd.EditExe(cols, "wf_form", where)

	return err
}
func DelFormByField(id string, formid string) error {
	daocheckokact.DeleteCheckRowByFieldid(id)
	daocheckokact.DeleteBtnOkRowByFieldid(id)

	where := map[string]any{}
	where["flow_id"] = formid
	where["field_id"] = id
	err := basecurd.DeleteExe("wf_form", where)

	return err
}
func DelFormByFormId(id string) error {

	fetchrows, _ := FetchFormRows(id)
	for _, value := range fetchrows {
		DelFormByField(value.Field_id, id)
	}

	where := map[string]any{}
	where["flow_id"] = id
	err := basecurd.DeleteExe("wf_form", where)

	return err
}
