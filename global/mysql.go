package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func EnvDefaultVal(k1, k2 string) string {
	res := os.Getenv(k1)
	if res == "" {
		return k2
	} else {
		return res
	}
}

func Mysql() {
	mysqlHost := EnvDefaultVal("MYSQL_HOST", "127.0.0.1")
	if db, err := gorm.Open("mysql", fmt.Sprintf("root:123@tcp(%s:3306)/p1?charset=utf8mb4&parseTime=True&loc=Local", mysqlHost)); err != nil {
		log.Fatalf("Mysql link failure %v", err)
	} else {
		GVA_DB = db
		GVA_DB.DB().SetMaxIdleConns(10)
		GVA_DB.DB().SetMaxOpenConns(100)
		GVA_DB.LogMode(false)
		log.Println("Mysql link success")
	}

}
