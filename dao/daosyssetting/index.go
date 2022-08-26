package daosyssetting

import (
	"github.com/jianliangan/flowmanage/dao/basecurd"
	"github.com/jianliangan/flowmanage/dto/dtosetting"
)

func businessConvert(col *dtosetting.Business_str) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "business_str_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "parentid"},
			basecurd.FieldProperty{NeedDate: false, Name: "name"},
			//basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
			basecurd.FieldProperty{NeedDate: false, Name: "desc"},
			basecurd.FieldProperty{NeedDate: false, Name: "status"},
			basecurd.FieldProperty{NeedDate: false, Name: "sort"})
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "business_str_id", Value: col.Business_str_id},
			basecurd.FieldProperty{NeedDate: false, Name: "parentid", Value: col.Parentid},
			basecurd.FieldProperty{NeedDate: false, Name: "name", Value: col.Name},
			//basecurd.FieldProperty{NeedDate: false, Name: "dtime", Value: col.Dtime},
			basecurd.FieldProperty{NeedDate: false, Name: "desc", Value: col.Desc},
			basecurd.FieldProperty{NeedDate: false, Name: "status", Value: col.Status},
			basecurd.FieldProperty{NeedDate: false, Name: "sort", Value: col.Sort})
	}

	return cols
}
func organizeConvert(col *dtosetting.Organize_str) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "organize_str_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "parentid"},
			basecurd.FieldProperty{NeedDate: false, Name: "name"},
			//basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
			basecurd.FieldProperty{NeedDate: false, Name: "desc"},
			basecurd.FieldProperty{NeedDate: false, Name: "status"},
			basecurd.FieldProperty{NeedDate: false, Name: "sort"})
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "organize_str_id", Value: col.Organize_str_id},
			basecurd.FieldProperty{NeedDate: false, Name: "parentid", Value: col.Parentid},
			basecurd.FieldProperty{NeedDate: false, Name: "name", Value: col.Name},
			//basecurd.FieldProperty{NeedDate: false, Name: "dtime", Value: col.Dtime},
			basecurd.FieldProperty{NeedDate: false, Name: "desc", Value: col.Desc},
			basecurd.FieldProperty{NeedDate: false, Name: "status", Value: col.Status},
			basecurd.FieldProperty{NeedDate: false, Name: "sort", Value: col.Sort})
	}
	return cols
}

//`role_id`, `rolename`, `sort`, `status`, `dtime`
func roleConvert(col *dtosetting.Role_str) []basecurd.FieldProperty {
	cols := []basecurd.FieldProperty{}
	if col == nil {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "role_id"},
			basecurd.FieldProperty{NeedDate: false, Name: "rolename"},
			basecurd.FieldProperty{NeedDate: false, Name: "sort"},
			basecurd.FieldProperty{NeedDate: false, Name: "status"},
			basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
		)
	} else {
		cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "role_id", Value: col.Role_id},
			basecurd.FieldProperty{NeedDate: false, Name: "rolename", Value: col.Rolename},
			basecurd.FieldProperty{NeedDate: false, Name: "sort", Value: col.Sort},
			basecurd.FieldProperty{NeedDate: false, Name: "status", Value: col.Status},
			basecurd.FieldProperty{NeedDate: false, Name: "dtime", Value: col.Dtime},
		)
	}
	return cols
}

func FetchBusinessRows(limitstart int, limitlen int, where map[string]any) ([]*dtosetting.Business_str, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "business_str_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "parentid"},
		basecurd.FieldProperty{NeedDate: false, Name: "name"},
		basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
		basecurd.FieldProperty{NeedDate: false, Name: "desc"},
		basecurd.FieldProperty{NeedDate: false, Name: "status"},
		basecurd.FieldProperty{NeedDate: false, Name: "sort"})

	//SelectExe(cols map[string]FieldProperty, dt string, where map[string]interface{}, limitstart int, limitend int, rows any, fetchAttr func(any) []any)
	rows := []*dtosetting.Business_str{}
	fetchBusinessAttr := func() []any {

		m := &dtosetting.Business_str{}
		ret := make([]any, 0)
		ret = append(ret, &m.Business_str_id)
		ret = append(ret, &m.Parentid)
		ret = append(ret, &m.Name)
		ret = append(ret, &m.Dtime)
		ret = append(ret, &m.Desc)
		ret = append(ret, &m.Status)
		ret = append(ret, &m.Sort)
		rows = append(rows, m)

		return ret
	}
	err := basecurd.SelectExe(cols, "business_struct", where, limitstart, limitlen, &rows, fetchBusinessAttr)

	return rows, err
}
func FetchBusinessCount() (int, error) {
	count, err := basecurd.CountRows("business_struct", nil)
	return count, err
}
func EditBusinessExe(row *dtosetting.Business_str) error {
	//for key,value :=range rows{
	cols := businessConvert(row)
	where := map[string]any{}
	where["business_str_id"] = row.Business_str_id
	err := basecurd.EditExe(cols, "business_struct", where)
	//}

	return err
}
func AddBusinessExe(row *dtosetting.Business_str) error {
	cols := businessConvert(row)
	err := basecurd.AddExe(cols, "business_struct")
	return err
}
func DelBusinessExe(row *dtosetting.Business_str) error {
	where := map[string]any{}
	where["business_str_id"] = row.Business_str_id
	err := basecurd.DeleteExe("business_struct", where)
	return err
}

//Organize
func FetchOrganizeRows(limitstart int, limitlen int, where map[string]any) ([]*dtosetting.Organize_str, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "organize_str_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "parentid"},
		basecurd.FieldProperty{NeedDate: false, Name: "name"},
		basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
		basecurd.FieldProperty{NeedDate: false, Name: "desc"},
		basecurd.FieldProperty{NeedDate: false, Name: "status"},
		basecurd.FieldProperty{NeedDate: false, Name: "sort"})

	//SelectExe(cols map[string]FieldProperty, dt string, where map[string]interface{}, limitstart int, limitend int, rows any, fetchAttr func(any) []any)
	rows := []*dtosetting.Organize_str{}

	fetchOrganizeAttr := func() []any {

		m := &dtosetting.Organize_str{}
		ret := make([]any, 0)
		ret = append(ret, &m.Organize_str_id)
		ret = append(ret, &m.Parentid)
		ret = append(ret, &m.Name)
		ret = append(ret, &m.Dtime)
		ret = append(ret, &m.Desc)
		ret = append(ret, &m.Status)
		ret = append(ret, &m.Sort)
		rows = append(rows, m)

		return ret
	}
	err := basecurd.SelectExe(cols, "organize_struct", where, limitstart, limitlen, &rows, fetchOrganizeAttr)
	// logv, _ := json.Marshal(rows)
	// jrlog.Logprintln("[INFO]", string(logv))

	return rows, err
}
func FetchOrganizeCount() (int, error) {
	count, err := basecurd.CountRows("organize_struct", nil)
	return count, err
}
func EditOrganizeExe(row *dtosetting.Organize_str) error {
	//for key,value :=range rows{
	cols := organizeConvert(row)
	where := map[string]any{}
	where["organize_str_id"] = row.Organize_str_id
	err := basecurd.EditExe(cols, "organize_struct", where)
	//}

	return err
}
func AddOrganizeExe(row *dtosetting.Organize_str) error {
	cols := organizeConvert(row)
	err := basecurd.AddExe(cols, "organize_struct")
	return err
}
func DelOrganizeExe(row *dtosetting.Organize_str) error {
	where := map[string]any{}
	where["organize_str_id"] = row.Organize_str_id
	err := basecurd.DeleteExe("organize_struct", where)
	return err
}

//user

func FetchUserRows(limitstart int, limitlen int, where map[string]any) ([]*dtosetting.User_str, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "user_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "username"},
		basecurd.FieldProperty{NeedDate: false, Name: "avatar"},
		basecurd.FieldProperty{NeedDate: false, Name: "email"},
		basecurd.FieldProperty{NeedDate: false, Name: "phone"},
		basecurd.FieldProperty{NeedDate: false, Name: "organize_str_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "business_str_id"})

	//SelectExe(cols map[string]FieldProperty, dt string, where map[string]interface{}, limitstart int, limitend int, rows any, fetchAttr func(any) []any)
	rows := []*dtosetting.User_str{}
	fetchUserAttr := func() []any {
		m := &dtosetting.User_str{}
		ret := make([]any, 0)
		ret = append(ret, &m.User_id)
		ret = append(ret, &m.Username)
		ret = append(ret, &m.Avatar)
		ret = append(ret, &m.Email)
		ret = append(ret, &m.Phone)
		ret = append(ret, &m.Organize_str_id)
		ret = append(ret, &m.Business_str_id)
		rows = append(rows, m)

		return ret
	}
	err := basecurd.SelectExe(cols, "user", where, limitstart, limitlen, &rows, fetchUserAttr)

	return rows, err
}
func FetchUserCount(where map[string]any) (int, error) {

	count, err := basecurd.CountRows("user", where)
	return count, err
}

//role
func FetchRoleRows(limitstart int, limitlen int, where map[string]any) ([]*dtosetting.Role_str, error) {
	cols := []basecurd.FieldProperty{}
	cols = append(cols, basecurd.FieldProperty{NeedDate: false, Name: "role_id"},
		basecurd.FieldProperty{NeedDate: false, Name: "rolename"},
		basecurd.FieldProperty{NeedDate: false, Name: "sort"},
		basecurd.FieldProperty{NeedDate: false, Name: "status"},
		basecurd.FieldProperty{NeedDate: false, Name: "dtime"},
	)
	rows := []*dtosetting.Role_str{}

	fetchRoleAttr := func() []any {
		m := &dtosetting.Role_str{}
		ret := make([]any, 0)
		ret = append(ret, &m.Role_id)
		ret = append(ret, &m.Rolename)
		ret = append(ret, &m.Dtime)
		ret = append(ret, &m.Sort)
		ret = append(ret, &m.Status)
		rows = append(rows, m)
		return ret
	}

	err := basecurd.SelectExe(cols, "role", where, limitstart, limitlen, &rows, fetchRoleAttr)

	return rows, err
}
func FetchRoleCount(where map[string]any) (int, error) {

	count, err := basecurd.CountRows("role", where)
	return count, err
}
func EditRoleExe(row *dtosetting.Role_str) error {
	//for key,value :=range rows{
	cols := roleConvert(row)
	where := map[string]any{}
	where["role_id"] = row.Role_id
	err := basecurd.EditExe(cols, "role", where)
	//}

	return err
}
func AddRoleExe(row *dtosetting.Role_str) error {
	cols := roleConvert(row)
	err := basecurd.AddExe(cols, "role")
	return err
}
func DelRoleExe(row *dtosetting.Role_str) error {
	where := map[string]any{}
	where["role_id"] = row.Role_id
	err := basecurd.DeleteExe("role", where)
	return err
}
