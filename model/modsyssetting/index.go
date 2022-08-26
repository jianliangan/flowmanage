package modsyssetting

import (
	"errors"

	"github.com/jianliangan/flowmanage/dao/daosyssetting"
	"github.com/jianliangan/flowmanage/dto/dtosetting"
	"github.com/jianliangan/flowmanage/libs/tools/goldflake"
	"github.com/jianliangan/flowmanage/model/modcontext"
)

type BusinessListInfos struct {
	Num  int                        `json:"Num"`
	Rows []*dtosetting.Business_str `json:"Rows"`
	Size int                        `json:"Size"`
}

func GetBusinessList() ([]*dtosetting.Business_str, error) {

	fetchRows, err := daosyssetting.FetchBusinessRows(0, 500, nil)
	//num, _ := daosyssetting.FetchBusinessCount()

	return fetchRows, err
}
func PushBusinessNewCmd(row *dtosetting.Business_str) (string, error) {
	var err error
	reid := "0"
	//for _, row := range *rows {
	if row.Cmd_ == "edit" {
		err = daosyssetting.EditBusinessExe(row)
	} else if row.Cmd_ == "delete" {
		where := map[string]any{}
		where["parentid"] = row.Business_str_id
		rowtmp, _ := daosyssetting.FetchBusinessRows(0, 500, where)
		if len(rowtmp) > 0 {
			err = errors.New("有子项")
		} else {
			err = daosyssetting.DelBusinessExe(row)
		}

	} else if row.Cmd_ == "add" {
		row.Business_str_id, _ = goldflake.Default().NextID()
		reid = row.Business_str_id
		err = daosyssetting.AddBusinessExe(row)
	}

	//}

	return reid, err
}

func GetOrganizeList() ([]*dtosetting.Organize_str, error) {

	fetchRows, err := daosyssetting.FetchOrganizeRows(0, 500, nil)

	return fetchRows, err
}
func PushOrganizeNewCmd(row *dtosetting.Organize_str) (string, error) {
	var err error
	reid := "0"
	//for _, row := range *rows {
	if row.Cmd_ == "edit" {
		err = daosyssetting.EditOrganizeExe(row)
	} else if row.Cmd_ == "delete" {
		where := map[string]any{}
		where["parentid"] = row.Organize_str_id
		rowtmp, _ := daosyssetting.FetchOrganizeRows(0, 500, where)
		if len(rowtmp) > 0 {
			err = errors.New("有子项")
		} else {
			err = daosyssetting.DelOrganizeExe(row)
		}

	} else if row.Cmd_ == "add" {
		row.Organize_str_id, _ = goldflake.Default().NextID()
		reid = row.Organize_str_id
		err = daosyssetting.AddOrganizeExe(row)
	}

	//}

	return reid, err
}

type ResUserListInfo struct {
	User     []*dtosetting.User_str     `json:"user"`
	Business []*dtosetting.Business_str `json:"business"`
	Organize []*dtosetting.Organize_str `json:"organize"`
	Num      int                        `json:"num"`
}

//user
func GetUserList(page int, organizeid string) (*ResUserListInfo, error) {
	if page == 0 {
		page = 1
	}

	limitlen := modcontext.Default().Pagesize
	limitstart := (page - 1) * limitlen
	where0 := map[string]any{}
	if organizeid != "" {
		where0["organize_str_id"] = organizeid
	}
	fetchRows, err := daosyssetting.FetchUserRows(limitstart, limitlen, where0)
	num, _ := daosyssetting.FetchUserCount(where0)
	organizeids := map[string]bool{}
	businessids := map[string]bool{}
	for _, value := range fetchRows {
		if value.Organize_str_id != "" {
			organizeids[value.Organize_str_id] = true
		}

		if value.Business_str_id != "" {
			businessids[value.Business_str_id] = true
		}
	}
	organizeidslist := make([]string, 0, len(organizeids))
	for key := range organizeids {
		organizeidslist = append(organizeidslist, key)
	}
	businessidslist := make([]string, 0, len(businessids))
	for key := range organizeids {
		businessidslist = append(businessidslist, key)
	}

	where1 := map[string]any{}
	where1["business_str_id"] = businessidslist
	resUserListInfo := &ResUserListInfo{}
	fetchBusinessRows, _ := daosyssetting.FetchBusinessRows(limitstart, limitlen, where1)
	where2 := map[string]any{}
	where2["organize_str_id"] = organizeidslist
	fetchOrganizeRows, _ := daosyssetting.FetchOrganizeRows(limitstart, limitlen, where2)

	resUserListInfo.Business = fetchBusinessRows
	resUserListInfo.Organize = fetchOrganizeRows
	resUserListInfo.Num = num
	resUserListInfo.User = fetchRows
	return resUserListInfo, err
}

//role
func GetRoleList() ([]*dtosetting.Role_str, error) {

	fetchRows, err := daosyssetting.FetchRoleRows(0, 500, nil)

	return fetchRows, err
}
func GetRoleNum() (int, error) {
	return daosyssetting.FetchRoleCount(nil)
}
func PushRoleNewCmd(row *dtosetting.Role_str) (string, error) {
	var err error
	reid := "0"
	//for _, row := range *rows {
	if row.Cmd_ == "edit" {
		err = daosyssetting.EditRoleExe(row)
	} else if row.Cmd_ == "delete" {
		where := map[string]any{}
		where["parentid"] = row.Role_id
		rowtmp, _ := daosyssetting.FetchRoleRows(0, 500, where)
		if len(rowtmp) > 0 {
			err = errors.New("有子项")
		} else {
			err = daosyssetting.DelRoleExe(row)
		}

	} else if row.Cmd_ == "add" {
		row.Role_id, _ = goldflake.Default().NextID()
		reid = row.Role_id
		err = daosyssetting.AddRoleExe(row)
	}

	//}

	return reid, err
}
