package appconfig

import (
	"fmt"
	"github.com/spf13/viper"
)

type ViperConfData struct {
	Nacos struct {
		SpaceId string
		Host    string
		Port    uint64
		DataId  string
		Group   string
	}
}

var ConfData ViperConfData

func GetViperConfData() {
	viper.SetConfigFile("../../Common/appconfig/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&ConfData)
	if err != nil {
		return
	}
	fmt.Println(ConfData)
	fmt.Println("Viper connect successfully")
}
