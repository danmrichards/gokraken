//go:generate stringer -type Currency
package asset

// Code generated by asset
// DO NOT EDIT

const (
{{- range $k, $v := .}}
{{- if eq $k 0 }}
{{$v}} Currency = iota
{{else -}}
{{$v}}
{{end -}}

{{- end}}
)

type Currency int

var validCurrencies = []Currency{
{{- range $k, $v := .}}
{{$v}},
{{- end}}
}