package plugins

import (
	"SafeDp/scanner/password_crack/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ScanMysql(service models.Service) (result models.ScanResult, err error) {
	result.Service = service
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", service.Username, service.Password, service.Ip, service.Port, "mysql")
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return result, err
	}
	err = db.Ping()
	if err != nil {
		return result, err
	}
	result.Result = true
	defer func() {
		if db != nil {
			_ = db.Close()
		}
	}()
	return result, err
}