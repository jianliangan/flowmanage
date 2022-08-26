package basecurd

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"github.com/jianliangan/flowmanage/libs/libdb"
)

type FieldProperty struct {
	Name     string //字段名称
	NeedDate bool   //select 时用到
	Value    any    //字段值
}

func getType(value any) int {
	switch value.(type) {
	case []uint64:
		return 2

	case []int64:
		return 2

	case []int:
		return 2

	case []string:
		return 2

	case []any:
		return 2

	default:
		return 1

	}

}
func SelectExe(cols []FieldProperty, dt string, where map[string]any, limitstart int, limitend int, rows any, fetchAttr func() []any) error {
	colslen := len(cols)
	if colslen == 0 {
		return errors.New("data empty")
	}
	if limitend > 2000 {
		return errors.New("limit too big")
	}
	sqlstr := "SELECT "
	for _, value := range cols {
		if value.NeedDate {
			sqlstr += "DATE_FORMAT(`" + value.Name + "`,'%Y-%m-%d'),"
		} else {
			sqlstr += "`" + value.Name + "`,"
		}
	}
	sqlstr = strings.Trim(sqlstr, ",")
	sqlstr += " FROM `" + dt + "` "
	params := []any{}
	if len(where) > 0 {
		sqlstr += " where "
		isfirst := true
		for key, value := range where {
			valuet := getType(value)
			if valuet == 1 {
				if isfirst {
					sqlstr += " `" + key + "`=? "
					isfirst = false
				} else {
					sqlstr += " and `" + key + "`=? "
				}

				params = append(params, value)
			} else if valuet == 2 {
				num := len(value.([]string))
				if num > 0 {

					if isfirst {
						sqlstr += " `" + key + "` in( "
						isfirst = false
					} else {
						sqlstr += " and `" + key + "` in( "
					}
					for _, value := range value.([]string) {
						params = append(params, value)
					}
					tmp := strings.Repeat("?,", num)
					sqlstr += strings.Trim(tmp, ",") + ") "
				}
			}

		}
	}
	sqlstr += " limit " + strconv.Itoa(limitstart) + "," + strconv.Itoa(limitend)
	res, err := libdb.Query(sqlstr, params...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	defer res.Close()

	for res.Next() {
		set := fetchAttr()
		if err = libdb.Scan(res, set...); err != nil {
			return err
		}

	}
	return nil
}

func EditExe(cols []FieldProperty, dt string, where map[string]any) error {
	colslen := len(cols)
	if colslen == 0 || len(where) == 0 {
		return errors.New("data empty")
	}
	sqlstr := "UPDATE `" + dt + "` set "
	params := []any{}
	isfirst := true
	for _, value := range cols {
		if isfirst {
			sqlstr += " `" + value.Name + "`=? "
			isfirst = false
		} else {
			sqlstr += " , `" + value.Name + "`=? "
		}
		params = append(params, value.Value)
	}
	sqlstr = strings.Trim(sqlstr, ",")
	if len(where) > 0 {
		sqlstr += " where "
		isfirst := true
		for key, value := range where {
			valuet := getType(value)
			if valuet == 1 {
				if isfirst {
					sqlstr += " `" + key + "`=? "
					isfirst = false
				} else {
					sqlstr += " and `" + key + "`=? "
				}

				params = append(params, value)
			} else if valuet == 2 {
				num := len(value.([]string))
				if num > 0 {

					if isfirst {
						sqlstr += " `" + key + "` in( "
						isfirst = false
					} else {
						sqlstr += " and `" + key + "` in( "
					}
					for _, value := range value.([]string) {
						params = append(params, value)
					}
					tmp := strings.Repeat("?,", num)
					sqlstr += strings.Trim(tmp, ",") + ") "
				}
			}

		}
	}

	_, err := libdb.Exec(sqlstr, params...)

	return err

}

func AddExe(cols []FieldProperty, dt string) error {
	colslen := len(cols)
	if colslen == 0 {
		return errors.New("data empty")
	}
	sqlstr := "INSERT INTO  `" + dt + "`("
	params := []any{}
	valuestr := ""
	for _, value := range cols {
		sqlstr += "`" + value.Name + "`,"
		params = append(params, value.Value)
		valuestr += "?,"
	}
	sqlstr = strings.Trim(sqlstr, ",")
	valuestr = strings.Trim(valuestr, ",")
	sqlstr += ") values(" + valuestr + ")"
	_, err := libdb.Exec(sqlstr, params...)
	return err

}

func DeleteExe(dt string, where map[string]any) error {

	if len(where) == 0 {
		return errors.New("data empty")
	}
	sqlstr := "DELETE  from `" + dt + "` "
	params := []any{}
	if len(where) > 0 {
		sqlstr += " where "
		isfirst := true
		for key, value := range where {
			valuet := getType(value)
			if valuet == 1 {
				if isfirst {
					sqlstr += " `" + key + "`=? "
					isfirst = false
				} else {
					sqlstr += " and `" + key + "`=? "
				}

				params = append(params, value)
			} else if valuet == 2 {
				num := len(value.([]string))
				if num > 0 {

					if isfirst {
						sqlstr += " `" + key + "` in( "
						isfirst = false
					} else {
						sqlstr += " and `" + key + "` in( "
					}
					for _, value := range value.([]string) {
						params = append(params, value)
					}
					tmp := strings.Repeat("?,", num)
					sqlstr += strings.Trim(tmp, ",") + ") "
				}
			}

		}
	}
	_, err := libdb.Exec(sqlstr, params...)
	return err

}
func CountRows(dt string, where map[string]any) (int, error) {
	sqlstr := "SELECT count(*) FROM `" + dt + "`"
	params := []any{}
	if len(where) > 0 {
		sqlstr += " where "
		isfirst := true
		for key, value := range where {
			valuet := getType(value)
			if valuet == 1 {
				if isfirst {
					sqlstr += " `" + key + "`=? "
					isfirst = false
				} else {
					sqlstr += " and `" + key + "`=? "
				}

				params = append(params, value)
			} else if valuet == 2 {
				num := len(value.([]string))
				if num > 0 {

					if isfirst {
						sqlstr += " `" + key + "` in( "
						isfirst = false
					} else {
						sqlstr += " and `" + key + "` in( "
					}
					for _, value := range value.([]string) {
						params = append(params, value)
					}
					tmp := strings.Repeat("?,", num)
					sqlstr += strings.Trim(tmp, ",") + ") "
				}
			}

		}
	}
	num := 0
	res := libdb.QueryRow(sqlstr, params...)
	if err := libdb.ScanRow(res, &num); err != nil {
		return num, err
	}
	return num, nil
}
