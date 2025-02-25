package main

import (
	"context"
	"fmt"
	"microservice/gateway/authentication"
	"microservice/gateway/plugin"
	"net/http"
)

var pluginName = "auth-handler"
var HandlerRegisterer = registerer(pluginName)
var ModifierRegisterer = registerer("custom-modifier")

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r)+"-value", authentication.RegisterHandlers)
}

func (r registerer) RegisterModifiers(f func(
	name string,
	factoryFunc func(map[string]interface{}) func(interface{}) (interface{}, error),
	appliesToRequest bool,
	appliesToResponse bool,
)) {
	f(string(r)+"-request", plugin.RequestModifier, true, false)
	f(string(r)+"-response", plugin.ResponseModifier, false, true)
	f(string(r)+"-new-modifier", plugin.NewCustomModifier, true, true)
	fmt.Println("Modifiers registered!")
}
