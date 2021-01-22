package main

import (
	"log"
	"strconv"
	"time"
	"zlog-fun/simple-get-tools/internal/myhttp"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config") // 读取json配置文件
	viper.AddConfigPath(".")      // 设置配置文件和可执行二进制文件在用一个目录
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	currentTime := strconv.FormatInt(time.Now().Unix(), 10)
	domain := viper.GetStringMap("domain")
	urlInfo := &myhttp.URLStructInfo{}
	urlInfo.Url = viper.GetString("host")
	urlInfo.Port = viper.GetString("port")
	for k, v := range domain {
		urlInfo.Host = k
		urlInfo.Path = v.(string) + "?" + currentTime
		if !urlInfo.Get() {
			log.Println(">>>>>test failure ", urlInfo.Host)
		} else {
			log.Println("success", urlInfo.Host)
		}
	}
}
