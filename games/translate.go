package games

import (
	"fmt"
	"github.com/spf13/viper"
)

func translate(filename string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigName(filename + ".yaml")
        vp.SetConfigType("yaml")
        vp.AddConfigPath("data/dics/")

        err := vp.ReadInConfig()
        if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
        }

	return vp
}

func translateSoul(name string) string {
	vp := translate("soul")
	return vp.Sub(name).GetString("vsub")
}

func translateKirin(name string) string {
	vp := translate("evol")
	return vp.GetString(name)
}
