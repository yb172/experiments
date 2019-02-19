package cfg

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Conf contains actual config
var Conf Config

// InitConfig initializes config
func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	//Substitute the _ to .
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while reading config: %s", err)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return fmt.Errorf("error while unmarshaling config: %s", err)
	}

	return nil
}
