/*
@Time : 2020/4/28 6:12 PM
*/
package models

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/xormplus/core"
	"github.com/xormplus/xorm"
	"io"
	"net/url"
	"path"
	"time"
)

var X *xorm.Engine
var DBErr = errors.New("database error")

func InitDBConn(logWriter io.Writer) {
	var (
		err      error
		location *time.Location
	)
	driver := viper.GetString(`database.driver`)
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.username`)
	dbPass := viper.GetString(`database.password`)
	dbName := viper.GetString(`database.dbname`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "true")
	val.Add("loc", "Asia/Shanghai")
	val.Add("charset", "utf8mb4")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	X, err = xorm.NewEngine(driver, dsn)
	if err != nil {
		panic(err)
	}
	if err = X.Ping(); err != nil {
		panic(err)
	}
	X.SetMaxIdleConns(10)
	X.SetMaxOpenConns(100)
	showSql := viper.GetBool("xorm.show-sql")
	templatePath := path.Join(viper.GetString("app-dir"), viper.GetString("xorm.template-path"))
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "")
	X.SetTableMapper(tbMapper)
	X.SetColumnMapper(core.GonicMapper{})
	location, _ = time.LoadLocation("Asia/Shanghai")
	X.TZLocation = location
	if showSql {
		if logWriter == nil {
			panic("init logWriter first!")
		}
		X.SetLogger(xorm.NewSimpleLogger(logWriter))
	}
	X.ShowSQL(showSql)
	X.ShowExecTime(true)
	X.Logger().SetLevel(core.LOG_INFO)
	//注册动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	//此处注册动态SQL模板配置，使用Pongo2模板引擎，配置文件根目录为"sql"，配置文件后缀为".stpl"
	err = X.RegisterSqlTemplate(xorm.Pongo2(templatePath, ".stpl"))
	if err != nil {
		panic(err)
	}
	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	err = X.StartFSWatcher()
	if err != nil {
		panic(err)
	}
	//cache := xorm.NewLRUCacher(xorm.NewMemoryStore(), 100000)
	//X.SetDefaultCacher(cache)
	// 禁用某个表缓存
	//X.MapCacher(&user, nil)
}

func Transact(txFunc func(*xorm.Session) error) (err error) {
	tx := X.NewSession()
	if err = tx.Begin(); err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil { // 防止txFunc代码出现runtime exception
			_ = tx.Rollback()
			//panic(p) 不Panic出去
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx)
	return
}
