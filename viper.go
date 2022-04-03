package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	vi := viper.New()
	vi.SetConfigName("test")
	vi.SetConfigType("yaml")
	vi.AddConfigPath(".")
	vi.AutomaticEnv()

	err := vi.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Declare var
	env := vi.GetString("app.env")
	producerbroker := vi.GetString("app.producerbroker")
	consumerbroker := vi.GetString("app.consumerbroker")
	linetoken := vi.GetString("app.linetoken")
	if linetoken == "" {
		linetoken = vi.GetString("linetoken")
	}

	// Print
	fmt.Println("---------- Dynamic values ----------")
	fmt.Println("app.env :", env)
	fmt.Println("app.producerbroker :", producerbroker)
	fmt.Println("app.consumerbroker :", consumerbroker)
	fmt.Println("app.linetoken :", linetoken)
}
