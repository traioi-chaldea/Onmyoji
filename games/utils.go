package games


func getGuildRaid(gRaid string) string {
	result := ""
	if gRaid != "false" {
		result = "- Kỳ Lân: " + gRaid
	}
	return result
}

func getNetherworldGate(nGate bool) string {
	result := ""
	if nGate {
		result = "- Cổng ma"
	}
	return result
}

func getDemonEncounter(dEnco string) string {
	result := ""
	if dEnco != "" {
		result = "- Boss thế giới: " + translateSoul(dEnco)
	}
	return result
}

func getMaintenance(m bool) string {
	result := ""
	if m {
		result = "- Bảo trì hàng tuần."
	}
	return result
}

func getSoulZone10(sZone10 []string) string {
	result := ""
	result += "- Rắn 10: "
	if sZone10[0] == "all" {
		result += "Tất cả các ngự."
	} else {
		for _, soul := range sZone10 {
			result += "\n  - " + translateSoul(soul)
		}
	}
	return result
}

func getSoulZone11(sZone11 []string) string {
	result := ""
	result += "- Rắn 11: "
	for _, soul := range sZone11 {
		result += "\n  - " + translateSoul(soul)
	}
	return result
}

func getTotemZone(tZone []string) string {
	result := ""
	if len(tZone) > 0 {
		result += "- Totem zone: " + tZone[0]
	}
	return result
}
