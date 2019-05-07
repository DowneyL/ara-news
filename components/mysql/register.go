package mysql

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils"
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

type ConnConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

var (
	configs = make(map[string]ConnConfig)
	ormers  = make(map[string]orm.Ormer)
)

func InitMysql() {
	initMysqlConnConfigs()
	registerDatabases()
}

func initMysqlConnConfigs() {
	dbIniPath := getConfigPath()
	b := utils.FileExists(dbIniPath)
	if !b {
		panic("Database config not exists")
	}

	cfg, err := ini.Load(dbIniPath)

	if err != nil {
		panic(fmt.Sprintf("Config File Read Error!, the err is %v", err))
	}

	sections := cfg.Sections()

	for _, section := range sections {
		name := section.Name()
		if name == "DEFAULT" {
			continue
		}
		sec := cfg.Section(name)
		host, _ := sec.GetKey("HOST")
		port, _ := sec.GetKey("PORT")
		username, _ := sec.GetKey("USERNAME")
		password, _ := sec.GetKey("PASSWORD")
		dbname, _ := sec.GetKey("DBNAME")
		configs[section.Name()] = ConnConfig{
			Host:     host.String(),
			Port:     port.String(),
			Username: username.String(),
			Password: password.String(),
			DbName:   dbname.String(),
		}
	}
}

func getConfigPath() string {
	return "conf/conn/db.ini"
}

func getDsn(config ConnConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		config.Username, config.Password, config.Host, config.Port, config.DbName,
	)
}

func registerDatabases() {
	var err error
	var dsn string
	if beego.BConfig.RunMode != beego.PROD {
		logs.Trace(configs)
	}

	for dbAttr, config := range configs {
		dsn = getDsn(config)
		if beego.BConfig.RunMode != beego.PROD {
			logs.Trace("Loading orm: " + dbAttr + " with dsn " + dsn)
		}
		switch dbAttr {
		case "slave":
			err = orm.RegisterDataBase("default", "mysql", dsn)
			orm.SetMaxIdleConns("default", 200)
			orm.SetMaxOpenConns("default", 200)
			break
		default:
			err = orm.RegisterDataBase(dbAttr, "mysql", dsn)
			orm.SetMaxIdleConns(dbAttr, 200)
			orm.SetMaxOpenConns(dbAttr, 200)
		}

		if err != nil {
			panic(err)
		}

		if beego.BConfig.RunMode == "dev" {
			orm.Debug = true
		}
	}
}
