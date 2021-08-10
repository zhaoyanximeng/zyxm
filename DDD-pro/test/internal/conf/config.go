package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Cfg struct {
	DBUser string `json:"db_user" yaml:"db_user" mapstructure:"db_user"`
	DBPassword string `json:"db_password" yaml:"db_password" mapstructure:"db_password"`
	DBHost string	`json:"db_host" yaml:"db_host" mapstructure:"db_host"`
	DBPort string `json:"db_port" yaml:"db_port" mapstructure:"db_port"`
}

var Config *Cfg

func init() {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./..")
	v.AddConfigPath("./../..")
	v.AddConfigPath("./../../..")

	c := &Cfg{}
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed:%v",err)
	}
	err = v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unmarshal config failed:%v",err)
	}
	Config = c
	fmt.Println(Config.DBPassword)
}