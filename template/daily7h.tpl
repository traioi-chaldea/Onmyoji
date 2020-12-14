{{ .Greeting }}
{{ .Time }} giờ hôm nay có:
{{- if ne .Content.GuildRaid "" }}
- Kỳ Lân: {{ .Content.GuildRaid }}
{{- end }}
{{- if .Content.NetherworldGate }}
- Cổng ma
{{- end }}
{{- if ne .Content.DemonEncounter "" }}
- Boss thế giới: {{ .Content.DemonEncounter }}
{{- end }}
