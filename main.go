package main

import (
	"github.com/extism/go-pdk"
	hostfuncs "github.com/spirefy/go-pdk"
	t "github.com/spirefy/go-pdk/types"
)

func getExtensions() []t.Extension {
	pluginGeneratorGoExtension := t.CreateExtension(
		"spirefy.plugingenerators.go",
		"Plugin Generator For Go",
		"spirefy.plugingenerators.generators",
		"Adds a Plugin Generator for Golang",
		"pluginGenerate",
		nil,
		nil)

	return []t.Extension{*pluginGeneratorGoExtension}
}

//export pluginGenerate
func gen() int32 {
	pdk.Log(pdk.LogDebug, "Go Plugin Generator generating output")

	return 0
}

//export start
func start() int32 {
	hostfuncs.RegisterPlugin("spirefy.plugin-generator-go", "Spirefy Plugin Golang Generator", "1.0.0", "1.0.0", "A plugin that provides an extension point for other plugins to extend that add plugin template generation which adhere to the Extism + Plugin Engine requirements.", nil, getExtensions())
	return 0
}

func main() {}
