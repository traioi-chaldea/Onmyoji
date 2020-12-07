package games

import(
	"fmt"
	"encoding/json"
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
	SoulZone11 []string `json:"soul_zone_11"`
	TotemZone []string  `json:"totem_zone"`
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
	this.Second.SoulZone11 = info.GetStringSlice("soul_zone_11")
	this.Second.TotemZone = info.GetStringSlice("totem_zone")
}

func (this *DAILY_INFO) getFrom18H(info *vp.Viper) {
	this.Third.DemonEncounter = info.GetString("demon_encounter")
	this.Third.NetherworldGate = info.GetBool("netherworld_gate")
}

func (this *DAILY_INFO) RunJSON() {
	result, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}

func (this *DAILY_INFO) parseFrom7H() string {
	var result string
	gRaid := getGuildRaid(this.First.GuildRaid)
	nGate := getNetherworldGate(this.First.NetherworldGate)
	dEnco := getDemonEncounter(this.First.DemonEncounter)

	result += fmt.Sprintf("7h có:\n")
	if gRaid != "" {
		result += fmt.Sprintf(gRaid) + "\n"
	}
	if nGate != "" {
		result += fmt.Sprintf(nGate) + "\n"
	}
	if dEnco != "" {
		result += fmt.Sprintf(dEnco) + "\n"
	}
	return result
}

func (this *DAILY_INFO) parseFrom12H() string {
	var result string
	m       := getMaintenance(this.Second.Maintenance)
	sZone10 := getSoulZone10(this.Second.SoulZone10)
	sZone11 := getSoulZone11(this.Second.SoulZone11)
	tZone   := getTotemZone(this.Second.TotemZone)

	result += fmt.Sprintf("12h có:\n")
	if m != "" {
		result += fmt.Sprintf(m) + "\n"
	}
	if sZone10 != "" {
		result += fmt.Sprintf(sZone10) + "\n"
	}
	if sZone11 != "" {
		result += fmt.Sprintf(sZone11) + "\n"
	}
	if tZone != "" {
		result += fmt.Sprintf(tZone) + "\n"
	}
	return result
}

func (this *DAILY_INFO) parseFrom18H() string {
	var result string
	dEnco := getDemonEncounter(this.Third.DemonEncounter)
	nGate := getNetherworldGate(this.Third.NetherworldGate)

	result += fmt.Sprintf("18h có:\n")
	if dEnco != "" {
		result += fmt.Sprintf(dEnco) + "\n"
	}
	if nGate != "" {
		result += fmt.Sprintf(nGate) + "\n"
	}
	return result
}

func (this *DAILY_INFO) Run(choice string) {
	switch choice {
	case "7":
		fmt.Print(this.parseFrom7H())
	case "12":
		fmt.Print(this.parseFrom12H())
	case "18":
		fmt.Print(this.parseFrom18H())
	}
}
