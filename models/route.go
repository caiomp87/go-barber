package models

type Route struct {
	Method string
	Name   string
	Path   string
	Params []string
}
