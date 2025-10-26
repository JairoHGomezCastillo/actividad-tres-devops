{{/*
Define el nombre completo del release (usado por Helm para evitar colisiones)
*/}}
{{- define "razas-api.fullname" -}}
{{- printf "%s-%s" .Release.Name .Chart.Name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
