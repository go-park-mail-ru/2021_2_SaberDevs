package server

import "github.com/spf13/viper"

func Config() (string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")
	return address + ":" + port, nil
}
