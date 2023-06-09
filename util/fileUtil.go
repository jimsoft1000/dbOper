package util

import (
	"fmt"
	"github.com/spf13/viper"
)

var ApplicationConfig *viper.Viper

// 读取application配置文件
func init() {
	ApplicationConfig = viper.New()
	ApplicationConfig.SetConfigName("application") //设置配置文件的名字
	ApplicationConfig.SetConfigType("toml")        //设置配置文件类型，可选
	ApplicationConfig.AddConfigPath("./config")    //添加配置文件所在的路径

	err := ApplicationConfig.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
