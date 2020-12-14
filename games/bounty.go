package games

import (
	"github.com/spf13/viper"
)

type BOUNTY_DATA struct {
	Hint string `yaml:"hint"`
	Shiki string `yaml:"shiki"`
}

type BOUNTY_LOCATION struct {
	Name string
}

func (this *BOUNTY_LOCATION) GetData() []BOUNTY_DATA {
	vp := viper.New()
	vp.SetConfigName("bounty.yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("data/dics/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	datas := vp.AllSettings()
	if err != nil {
		panic(err)
	}
	var result []BOUNTY_DATA
	for _, data := range datas {
		tmp := data.(map[string]interface{})
		result = append(result, BOUNTY_DATA{
			Hint: tmp["hint"].(string),
			Shiki: tmp["shiki"].(string),
		})
	}
	return result
}

func (this *BOUNTY_LOCATION) GetVName(name string) string {
	return translateShikigami(name, "vsub")
}

func (this *BOUNTY_LOCATION) GetEName(name string) string {
	return translateShikigami(name, "esub")
}
