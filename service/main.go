package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"log"
	"time"
	"audioscoring/common"
	"audioscoring/score"
	"audioscoring/crv"
)

func main() {
	//设置log打印文件名和行号
  	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//初始化时区
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone

	conf:=common.InitConfig()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:true,
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
  	}))
	
	crvClient:=&crv.CRVClient{
		Server:conf.CRV.Server,
		Token:conf.CRV.Token,
	}
	
	scoreController:=score.ScoreController{
		CRVClient:crvClient,
	}
	
	scoreController.Bind(router)

	router.Run(conf.Service.Port)
}