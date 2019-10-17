module github.com/asciifaceman/biglybot

go 1.13

replace github.com/mb-14/gomarkov => ../gomarkov

require (
	github.com/mb-14/gomarkov v0.0.0-20190125094512-044dd0dcb5e7
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
)
