{{/*
Helper function to get the proper image prefix
*/}}
{{- define "cray-hms-bss.image-prefix" -}}
    {{ $base := index . "cray-service" }}
    {{- if $base.imagesHost -}}
        {{- printf "%s/" $base.imagesHost -}}
    {{- else -}}
        {{- printf "" -}}
    {{- end -}}
{{- end -}}