package gobee

import (
	"git.sillydong.com/chenzhidong/goczd/gofile"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"path/filepath"
	//_ "github.com/mattn/go-sqlite3"
	"fmt"
	_ "github.com/astaxie/beego/cache/memcache"
	"github.com/garyburd/redigo/redis"
	"os"
	"runtime"
	"syscall"
	"time"
)

func InitLog() {
	beego.Info("initiating log")
	logconf, err := gofile.IoUtilReadFile(filepath.Join(beego.AppPath, "conf", "log.json"))
	if err != nil {
		panic(err)
	} else {
		beego.SetLogger("file", string(logconf))
	}
}

//初始化memcache缓存，配置参数如下
//
// memcache_host memcache主机(127.0.0.1)
//
// memcache_port memcache端口(11211)  
func InitMemcache() (cache.Cache, error) {
	beego.Info("initiating memcache")
	memcache_host := beego.AppConfig.DefaultString("memcache_host", "127.0.0.1")
	memcache_port := beego.AppConfig.DefaultString("memcache_port", "11211")
	return cache.NewCache("memcache", fmt.Sprintf(`{"conn":"%s:%s"`, memcache_host, memcache_port))
}

//初始化文件缓存，配置参数如下
//
// filecache_dir 缓存文件目录
//
// filecache_suffix 缓存文件后缀(.cache)
//
// filecache_level 目录层级(2)
//
// filecache_expire 过期时间(3600秒)
func InitFilecache() (cache.Cache, error) {
	beego.Info("initiating file cache")
	filecache_dir := beego.AppConfig.DefaultString("filecache_dir", "")
	filecache_suffix := beego.AppConfig.DefaultString("filecache_suffix", ".cache")
	filecache_level := beego.AppConfig.DefaultInt("filecache_level", 2)
	filecache_expire := beego.AppConfig.DefaultInt("filecache_expire", 3600)
	if filecache_dir == "" {
		panic("filecache_dir is not set")
	}
	info, err := os.Stat(filecache_dir)
	if err != nil && !os.IsExist(err) {
		if err := os.MkdirAll(filecache_dir, 777); err != nil {
			panic(fmt.Sprintf("%s not exist and can not be created\n%v", filecache_dir, err))
		}
	}
	if !info.IsDir() {
		panic(fmt.Sprintf("%s is not a directory", filecache_dir))
	}
	if err := syscall.Access(filecache_dir, syscall.O_RDWR); err != nil {
		panic(fmt.Sprintf("%s is not accessable\n%v", filecache_dir, err))
	}
	return cache.NewCache("file", fmt.Sprintf(`{"CachePath":"%s","FileSuffix":"%s","DirectoryLevel":%d,"EmbedExpiry":%d}`, filecache_suffix, filecache_suffix, filecache_level, filecache_expire))
}

//初始化内存缓存，配置参数如下
//
// memory_gc_interval 内存回收周期(60秒)
func InitMemorycache() (cache.Cache, error) {
	beego.Info("initiating memory cache")
	memory_interval := beego.AppConfig.DefaultInt("memory_gc_interval", 60)
	return cache.NewCache("memory", fmt.Sprintf(`{"interval":%d}`, memory_interval))
}

//初始化redis缓存，配置参数如下
//
// rediscache_host redis主机(127.0.0.1)
//
// rediscache_port redis端口(6379)
func InitRediscache() (cache.Cache, error) {
	beego.Info("initiating redis cache")
	rediscache_host := beego.AppConfig.DefaultString("rediscache_host", "127.0.0.1")
	rediscache_port := beego.AppConfig.DefaultString("rediscache_port", "6379")
	return cache.NewCache("redis", fmt.Sprintf(`{"conn":"%s:%s"}`, rediscache_host, rediscache_port))
}

//初始化mysql或postgresql数据库，配置参数如下
//
// db_type 数据库类型，mysql或postgres(mysql)
//
// db_host 数据库主机地址(127.0.0.1)
//
// db_port 数据库端口(3306)
//
// db_user 数据库用户(root)
//
// db_pass 数据库密码
//
// db_name 数据库名
//
// db_sslmode 是否SSL模式连接，仅用于postgresql(disable)
func InitDB() {
	beego.Info("initiating db")
	var dsn string
	db_type := beego.AppConfig.DefaultString("db_type", "mysql")
	db_host := beego.AppConfig.DefaultString("db_host", "127.0.0.1")
	db_port := beego.AppConfig.DefaultString("db_port", 3306)
	db_user := beego.AppConfig.DefaultString("db_user", "root")
	db_pass := beego.AppConfig.DefaultString("db_pass", "")
	db_name := beego.AppConfig.String("db_name")
	//db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.DefaultString("db_sslmode", "disable")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DR_Postgres)
		dsn = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
		//	case "sqlite3":
		//		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
		//		if db_path == "" {
		//			db_path = "./"
		//		}
		//		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		//		break
	default:
		panic("Database driver is not allowed:" + db_type)
	}
	orm.RegisterDataBase("default", db_type, dsn)
}

//初始化redis数据库，配置参数如下
//
// redis_host redis主机(127.0.0.1)
//
// redis_port redis端口(6379)
//
// redis_pass redis密码
//
// redis_db redis使用的数据库(0)
func InitRedis() *redis.Pool {
	beego.Info("initiating redis db")
	redis_host := beego.AppConfig.DefaultString("redis_host", "127.0.0.1")
	redis_port := beego.AppConfig.DefaultString("redis_port", "6379")
	redis_pass := beego.AppConfig.DefaultString("redis_pass", "")
	redis_db := beego.AppConfig.DefaultString("redis_db", "")

	return &redis.Pool{
		MaxIdle:     runtime.NumCPU(),
		IdleTimeout: 180 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				beego.Critical(err)
			}
			return err
		},
		Dial: func() (redis.Conn, error) {
			conn, err := redis.DialTimeout("tcp", fmt.Sprintf("%s:%s", redis_host, redis_port), 1*time.Second, 1*time.Second, 1*time.Second)
			if err != nil {
				return nil, err
			}
			if redis_pass != "" {
				if _, err := conn.Do("AUTH", redis_pass); err != nil {
					return nil, err
				}
			}
			if redis_db != "" {
				if _, err := conn.Do("SELECT", redis_db); err != nil {
					return nil, err
				}
			}
			return conn, nil
		},
	}
}
