package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"pb/api/v1"
	"pb/global"
	"pb/model"
)

func Tables() []interface{} {
	return []interface{}{
		&model.Topic{},
		&model.Tag{},
		&model.Post{},
		&model.Category{},
		&model.AuthenCode{},
	}
}

func SetRunMode() {
	mode := "release"
	gin.SetMode(mode)
	log.Printf("running in %v mode ", mode)
}

func MigrateTables() {
	global.GVA_DB.SingularTable(true)
	for _, table := range Tables() {
		global.GVA_DB.Debug().AutoMigrate(table)
	}
	log.Println("table migrate success")
}

func DataBaseRun() {
	//var user = model.User{ Age: 99}
	//global.GVA_DB.Debug().Create(&user)
}

func LoadStaticAndHtml(r *gin.Engine) {
	r.LoadHTMLFiles("./dist/index.html")

	r.Static("/static", "./dist/static")

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=3153600")
		c.File("dist/favicon.ico")
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
}

func main() {
	SetRunMode()

	global.Mysql()
	MigrateTables()
	DataBaseRun()

	defer global.GVA_DB.Close()

	Router := api.Routers()
	LoadStaticAndHtml(Router)

	Router.Run(":9933")
}
