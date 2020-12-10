{{ .Greeting }}
{{ .Time }} giờ hôm nay có:
{{- if .Content.Maintenance }}
- Bảo trì hàng tuần.
{{- end }}
- Rắn 10:
{{- $tmp := index .Content.SoulZone10 0 }}
{{- if eq $tmp "all" }}
  - Tất cả các ngự.
{{- else }}
{{- range $index, $soul := .Content.SoulZone10 }}
  - {{ $soul -}}
{{- end }}
{{- end }}
- Rắn 11: {{ .Content.SoulZone11 }}
{{- if eq .Content.TotemZone "false" }}
{{- else if eq .Content.TotemZone "all" }}
- Totem zone: Tất cả.
{{- else }}
- Totem zone: {{ .Content.TotemZone }}
{{- end }}
