package godb

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/olebedev/config"
	"git.sillydong.com/chenzhidong/goczd/gofile"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"git.sillydong.com/chenzhidong/goczd/goxorm"
)

var (
	db        *xorm.Engine
	HasEngine bool
	DbCfg     struct {
		Host, DbName, UserName, PassWord, LogDir string
		MaxConn                                  int
	}
)

func LoadConfig(conf string) error {
	if gofile.FileExists(conf) {
		cfg, err := config.ParseJsonFile(conf)
		if err != nil {
			return err
		}else {
			DbCfg.Host=cfg.UString("host", "localhost:3306")
			DbCfg.DbName=cfg.UString("dbname", "")
			DbCfg.UserName=cfg.UString("username", "")
			DbCfg.PassWord=cfg.UString("password", "")
			DbCfg.LogDir=cfg.UString("logdir", "")
			DbCfg.MaxConn=cfg.UInt("maxconn", 0)

			return nil
		}
	}else {
		return fmt.Errorf("file not exists: %s", conf)
	}
}

func InitEngine() error {
	cnnstr := ""
	if DbCfg.Host[0] == '/' { // looks like a unix socket
		cnnstr = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8",
			DbCfg.UserName, DbCfg.PassWord, DbCfg.Host, DbCfg.DbName)
	} else {
		cnnstr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
			DbCfg.UserName, DbCfg.PassWord, DbCfg.Host, DbCfg.DbName)
	}

	db, err := xorm.NewEngine("mysql", cnnstr)
	if err != nil {
		return err
	}

	db.SetDefaultCacher(xorm.NewLRUCacher2(xorm.NewMemoryStore(), 3600, 1000))
	db.SetMapper(core.NewCacheMapper(core.GonicMapper{}))
	db.SetMaxConns(DbCfg.MaxConn)
	db.SetLogger(goxorm.NewXormLogger(DbCfg.LogDir, "db.log", core.LOG_WARNING))

	db.ShowSQL = true
	db.ShowInfo = false
	db.ShowDebug = false
	db.ShowErr = true
	db.ShowWarn = true
	
	if err=Ping();err!=nil{
		return err
	}

	HasEngine=true
	return nil
}

func GetDb()(*xorm.Engine,error){
	if HasEngine && db!=nil{
		return db,nil
	}
	return nil,fmt.Errorf("db not inited, exec LoadConfig & InitEngine first")
}

func Ping() error {
	return db.Ping()
}

func DumpDatabase(filePath string) error {
	return db.DumpAllToFile(filePath)
}
