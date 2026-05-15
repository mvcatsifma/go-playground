package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Loads config.yml, applies a default, and auto-reads env vars prefixed with CRAWLER_.
func main() {
	viper.AddConfigPath(".")
	viper.SetDefault("max_num_goroutine", 16)
	viper.SetEnvPrefix("crawler")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	maxNumGoroutine := viper.GetInt("max_num_goroutine")
	databaseUrl := viper.GetString("database_url")
	indexUrl := viper.GetString("index_url")
	shares := viper.GetStringMapString("shares")

	fmt.Println(maxNumGoroutine)
	fmt.Println(databaseUrl)
	fmt.Println(indexUrl)
	fmt.Println(shares)
}
