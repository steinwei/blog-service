package inits

import (
    "fmt"

    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigName("env")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()

    if err != nil {
        _ = fmt.Errorf("fatal err config file: %w", err)
    }
}
