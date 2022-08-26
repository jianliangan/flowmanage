package libdb

import (
	"database/sql"
	"runtime/debug"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jianliangan/flowmanage/libs/tools/jrlog"
)

type Mysql struct {
	Dburi   string
	Maxopen int
	Maxidle int
}

//连接句柄
var db *sql.DB

/*func Default() *sql.DB {
	return db
}*/
func ScanRow(res *sql.Row, dest ...any) error {
	err := res.Scan(dest...)
	if err != nil {
		jrlog.Logprintln("[ERR]", "mysql:", err)
	}
	var reerr error
	if err != nil {
		reerr = err
	}
	return reerr
}
func Scan(res *sql.Rows, dest ...any) error {
	err := res.Scan(dest...)
	if err != nil {
		jrlog.Logprintln("[ERR]", "mysql:", err)
	}
	var reerr error
	if err != nil {
		reerr = err
	}
	return reerr
}
func QueryRow(query string, args ...any) *sql.Row {
	row := db.QueryRow(query, args...)
	return row
}
func Query(query string, args ...any) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	rows, err = db.Query(query, args...)
	if err != nil {

		jrlog.Logprintln("[ERR]", "mysql:", query, args, err, string(debug.Stack()))
	}
	var reerr error
	if err != nil {
		reerr = err
	}
	return rows, reerr
}
func Exec(query string, args ...any) (sql.Result, error) {

	result, err := db.Exec(query, args...)
	if err != nil {
		jrlog.Logprintln("[ERR]", "mysql:", query, args, err, string(debug.Stack()))
	}
	var reerr error
	if err != nil {
		reerr = err
	}

	return result, reerr
}

func NewMySQL(env *Mysql) error {
	var err error
	db, err = sql.Open("mysql", env.Dburi)
	if err != nil {
		jrlog.Logprintln("[ERR]", "mysql:", err)
		var reerr error
		if err != nil {
			reerr = err
		}
		return reerr
	}
	db.SetMaxOpenConns(env.Maxopen)
	db.SetMaxIdleConns(env.Maxidle)

	return nil
}
