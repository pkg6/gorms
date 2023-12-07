package gorms

import (
	"github.com/spf13/viper"
)

func Viper(configFile string, cfg any) error {
	vip := viper.New()
	vip.SetConfigFile(configFile)
	if err := vip.ReadInConfig(); err != nil {
		return err
	}
	return vip.Unmarshal(cfg)
}
