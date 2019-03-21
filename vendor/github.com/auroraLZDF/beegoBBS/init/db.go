package init

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var cfg = beego.AppConfig

/*func InitDb() {

	// database
	dbUser := cfg.String("mysql_user")
	dbPass := cfg.String("mysql_pass")
	dbHost := cfg.String("mysql_host")
	dbPort := cfg.String("mysql_port")
	dbName := cfg.String("mysql_db")
	maxIdleConn, _ := cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := cfg.Int("db_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	RunMode := cfg.String("RunMode")
	if RunMode == "dev" {
		orm.Debug = true
	}

	orm.RegisterModel(new(models.Users))
	//RegisterModelWithPrefix
	//orm.RegisterModelWithPrefix("bbs_", new(models.Users))
}*/