package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"webGo/config"
)
var Engine = Database()
func Database()*xorm.Engine  {
	driver := config.Conf.Get("database.driver").(string)
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Conf.Get("mysql.databaseUserName").(string),
		config.Conf.Get("mysql.databasePassword").(string),
		config.Conf.Get("mysql.host").(string),
		config.Conf.Get("mysql.port").(int64),
		config.Conf.Get("mysql.databaseName").(string),
		config.Conf.Get("mysql.charset").(string))
	engine, err :=xorm.NewEngine(driver,dataSource)
	if err!=nil{
		log.Fatal(err.Error())
	}
	engine.ShowSQL(true)
	return engine
}