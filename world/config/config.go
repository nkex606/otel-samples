package config

import "github.com/spf13/viper"

func InitConf() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			panic("Config.yaml file not found.")
		default:
			panic("Failed to load config.yaml.")
		}
	}
	return nil
}

type httpConfig struct {
	Port string `mapstructure:"port"`
}

type mysqlConfig struct {
	Host            string `mapstructure:"host"`
	Port            string `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"db"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type otelConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

func GetHttpConfig() httpConfig {
	return httpConfig{
		Port: viper.GetString("http.port"),
	}
}

func GetMysqlConfig() *mysqlConfig {
	return &mysqlConfig{
		Host:            viper.GetString("mysql.host"),
		Port:            viper.GetString("mysql.port"),
		User:            viper.GetString("mysql.user"),
		Password:        viper.GetString("mysql.password"),
		DBName:          viper.GetString("mysql.db"),
		MaxIdleConns:    viper.GetInt("mysql.maxIdleConns"),
		MaxOpenConns:    viper.GetInt("mysql.maxOpenConns"),
		ConnMaxLifeTime: viper.GetInt("mysql.connMaxLifeTime"),
	}
}

func GetOTELConfig() otelConfig {
	return otelConfig{
		Endpoint: viper.GetString("otel_collector.endpoint"),
	}
}

func GetSvcName() string {
	return viper.GetString("svcName")
}
