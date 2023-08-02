package server

import (
	"github.com/spf13/viper"
)

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

func TarantoolConfig() (string, string, string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", "", "", err
	}

	user := viper.GetString("tarantool.user")
	pass := viper.GetString("tarantool.pass")
	addr := viper.GetString("tarantool.addr")
	return user, pass, addr, nil
}

func DbConfig() (string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	user := viper.GetString("postgres.user")
	dbname := viper.GetString("postgres.dbname")
	password := viper.GetString("postgres.password")
	host := viper.GetString("postgres.host")
	sslmode := viper.GetString("postgres.sslmode")

	return "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " sslmode=" + sslmode, nil
}
