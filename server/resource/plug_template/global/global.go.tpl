package global

{{- if .HasGlobal }}

import "github.com/liujianjiang/goadmin/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}