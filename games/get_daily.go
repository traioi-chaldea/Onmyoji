package games

import(
	"fmt"
	vp "github.com/spf13/viper"
)

type FIRST_ZONE struct {
	GuildRaid string      `json:"guild_raid"`
	NetherworldGate bool  `json:"netherworld_gate"`
	DemonEncounter string `json:"demon_encounter"`
}

type SECOND_ZONE struct {
	Maintenance bool    `json:"maintenance"`
	SoulZone10 []string `json:"soul_zone_10"`
	SoulZone11 string `json:"soul_zone_11"`
	TotemZone string  `json:"totem_zone"`
}

type THIRD_ZONE struct {
	DemonEncounter string `json:"demon_encounter"`
	NetherworldGate bool  `json:"netherworld_gate"`
}

type DAILY_INFO struct {
	First FIRST_ZONE   `json:"7"`
	Second SECOND_ZONE `json:"12"`
	Third THIRD_ZONE   `json:"18"`
}

func (this *DAILY_INFO) GetTime(filename string) {
	vp.SetConfigName(filename + ".yaml")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("data/daily/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	this.getFrom7H(vp.Sub("7"))
	this.getFrom12H(vp.Sub("12"))
	this.getFrom18H(vp.Sub("18"))

}

func (this *DAILY_INFO) getFrom7H(info *vp.Viper) {
	this.First.GuildRaid = info.GetString("guild_raid")
	this.First.NetherworldGate = info.GetBool("netherworld_gate")
	this.First.DemonEncounter = info.GetString("demon_encounter")
}

func (this *DAILY_INFO) getFrom12H(info *vp.Viper) {
	this.Second.Maintenance = info.GetBool("maintenance")
	this.Second.SoulZone10 = info.GetStringSlice("soul_zone_10")
	this.Second.SoulZone11 = info.GetString("soul_zone_11")
	this.Second.TotemZone = info.GetString("totem_zone")
}

func (this *DAILY_INFO) getFrom18H(info *vp.Viper) {
	this.Third.DemonEncounter = info.GetString("demon_encounter")
	this.Third.NetherworldGate = info.GetBool("netherworld_gate")
}

func (this *DAILY_INFO) parse7H() {
	this.First = FIRST_ZONE{
		this.First.GuildRaid,
		this.First.NetherworldGate,
		translateSoul(this.First.DemonEncounter),
	}
}

func (this *DAILY_INFO) parse12H() {
        for index, _ := range this.Second.SoulZone10 {
                 tmp := this.Second.SoulZone10[index]
                 this.Second.SoulZone10[index] = translateSoul(tmp)
        }
	this.Second = SECOND_ZONE{
		this.Second.Maintenance,
		this.Second.SoulZone10,
		translateSoul(this.Second.SoulZone11),
		this.Second.TotemZone,
	}
}

func (this *DAILY_INFO) parse18H() {
	this.Third = THIRD_ZONE{
		translateSoul(this.Third.DemonEncounter),
		this.Third.NetherworldGate,
	}
}

func (this *DAILY_INFO) print7H() FIRST_ZONE {
	this.parse7H()
	return this.First
}

func (this *DAILY_INFO) print12H() SECOND_ZONE {
	this.parse12H()
	return this.Second
}

func (this *DAILY_INFO) print18H() THIRD_ZONE {
	this.parse18H()
	return this.Third
}

func (this *DAILY_INFO) Run(choice string) interface{} {
	var result interface{}
	switch choice {
	case "7":
		result = this.print7H()
	case "12":
		result = this.print12H()
	case "18":
		result = this.print18H()
	}
	return result
}
