package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
	"time"
)

var (
	DBHost      string
	DBPort      int
	DBUser      string
	DBPawd      string
	DBName      string
	DBCharset   string
	DBTimeLoc   string
	DBMaxIdle   int
	DBMaxConn   int
	DBDebug     bool
	RedisClient *redis.Pool
	REDIS_HOST  string
	REDIS_DB    int
)

func initDB() {
	var (
		err error
	)
	DBHost = strings.TrimSpace(beego.AppConfig.String("db::host"))
	if "" == DBHost {
		panic("app parameter `db::host` empty")
	}

	DBPort, err = beego.AppConfig.Int("db::port")
	if err != nil {
		panic("app parameter `db::port` error")
	}
	DBUser = strings.TrimSpace(beego.AppConfig.String("db::user"))
	if "" == DBUser {
		panic("app parameter `db::user` empty")
	}

	DBPawd = strings.TrimSpace(beego.AppConfig.String("db::pawd"))
	if "" == DBPawd {
		panic("app parameter `db::pawd` empty")
	}

	DBName = strings.TrimSpace(beego.AppConfig.String("db::name"))
	if "" == DBName {
		panic("app parameter `db::name` empty")
	}

	DBCharset = strings.TrimSpace(beego.AppConfig.String("db::charset"))
	if "" == DBCharset {
		panic("app parameter `db::charset` empty")
	}

	DBTimeLoc = strings.TrimSpace(beego.AppConfig.String("db::time_loc"))
	if "" == DBTimeLoc {
		panic("app parameter `db::time_loc` empty")
	}

	DBMaxIdle, err = beego.AppConfig.Int("db::max_idle")
	if err != nil {
		panic("app parameter `db::max_idle` error")
	}

	DBMaxConn, err = beego.AppConfig.Int("db::max_conn")
	if err != nil {
		panic("app parameter `db::max_conn` error")
	}

	setDb()

	return
}

func setDb(){
	
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s", DBUser, DBPawd, DBHost, DBPort, DBName, DBCharset, url.QueryEscape(DBTimeLoc))
	
	err := orm.RegisterDataBase("default", "mysql", dataSourceName, DBMaxIdle, DBMaxConn)

	if err != nil {
		
				panic("err: " + err.Error())
		
			}
	//此处需要优化不同DB账号的引入(查询与修改操作账号需要分开)
	//1.修改conf，添加多账号
	//2.修改读取与DB初始化
	/*dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s", DBUser, DBPawd, DBHost, DBPort, DBName, DBCharset, url.QueryEscape(DBTimeLoc))

	err = orm.RegisterDataBase("test", "mysql", dataSourceName, DBMaxIdle, DBMaxConn)

	if err != nil {

		panic("err: " + err.Error())

	}*/

	fmt.Println("MySQL初始化成功")
}

func initRedis() {

	REDIS_HOST = beego.AppConfig.String("redis::host")

	REDIS_DB, _ = beego.AppConfig.Int("redis::db")

	RedisClient = &redis.Pool{

		MaxIdle: beego.AppConfig.DefaultInt("redis.maxidle", 1),

		MaxActive: beego.AppConfig.DefaultInt("redis.maxactive", 10),

		IdleTimeout: 180 * time.Second,

		Dial: func() (redis.Conn, error) {

			c, err := redis.Dial("tcp", REDIS_HOST)

			if err != nil {

				return nil, err

			}

			c.Do("SELECT", REDIS_DB)

			return c, nil

		},
	}
}

func Init() {
	initDB()

	initRedis()

	// orm debug
	DBDebug, err := beego.AppConfig.Bool("dev::debug")
	if err != nil {
		panic("app parameter `dev::debug` error:" + err.Error())
	}
	if DBDebug {
		orm.Debug = true
	}

	fmt.Println("框架初始化成功")
}
