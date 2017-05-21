package main

import (
	"github.com/yudeguang/config"
	"log"
)

func main() {
	defaultConfig()
	nonDefaultConfig()
}
func defaultConfig() {
	conf, err := config.NewConfig("")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("默认配置测试:")
	log.Printf("key1=:%v\n", conf.Get("key1"))
}
func nonDefaultConfig() {
	log.Println("非默认配置测试:")
	conf, err := config.NewConfig("nonDefaultConfig.txt")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("key1=:%v\n", conf.Get("key1"))
}
