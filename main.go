package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gostudy/config"
	"gostudy/router"
	"log"
	"net/http"
)
var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)
func main() {
	pflag.Parse()
	if err:=config.Init(*cfg); err!=nil{
		panic(err)
	}
	g :=gin.New()

	middlewares :=[]gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...,
		)
	log.Printf("开启服务")
	log.Printf(http.ListenAndServe(viper.GetString("addr"),g).Error())



}
