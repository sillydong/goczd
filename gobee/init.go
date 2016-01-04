package gobee
import (
	"git.sillydong.com/chenzhidong/goczd/gofile"
	"path/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/astaxie/beego/cache/memcache"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"runtime"
	"time"
)

func InitLog(){
	beego.Info("initiating log")
	logconf, err := gofile.IoUtilReadFile(filepath.Join(beego.AppPath, "conf", "log.json"))
	if err != nil {
		panic(err)
	} else {
		beego.SetLogger("file", string(logconf))
	}
}

func InitMemcache() {
	beego.Info("initiating memcache")
	memconf := beego.AppConfig.DefaultString("memcache", "")
	if len(memconf) >= 0 {
		_, err := cache.NewCache("memcache", memconf)
		if err != nil {
			panic(err)
		}
	}else{
		panic("missing memcache config")
	}
}

func InitDB() {
	beego.Info("initiating db")
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	//db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DR_Postgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
//	case "sqlite3":
//		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
//		if db_path == "" {
//			db_path = "./"
//		}
//		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
//		break
	default:
		panic("Database driver is not allowed:"+db_type)
	}
	orm.RegisterDataBase("default", db_type, dns)
}

func InitRedis() *redis.Pool {
	beego.Info("initiating redis")
	redishost := beego.AppConfig.DefaultString("redis_host", "127.0.0.1:6379")
	redispass := beego.AppConfig.DefaultString("redis_pass", "")
	redisdb := beego.AppConfig.DefaultString("redis_db", "")

	return &redis.Pool{
		MaxIdle: runtime.NumCPU(),
		IdleTimeout: 180 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				beego.Critical(err)
			}
			return err
		},
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialTimeout("tcp", redishost, 1*time.Second, 1*time.Second, 1*time.Second)
			if err != nil {
				return nil, err
			}
			if redispass != "" {
				if _, err := conn.Do("AUTH", redispass); err != nil {
					return nil, err
				}
			}
			if redisdb != "" {
				if _, err := conn.Do("SELECT", redisdb); err != nil {
					return nil, err
				}
			}
			return conn, nil
		},
	}
}
