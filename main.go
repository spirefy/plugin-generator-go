package main

import (
	"encoding/json"
	"github.com/extism/go-pdk"
	hostfuncs "github.com/spirefy/go-pdk"
	t "github.com/spirefy/go-pdk/types"
	"strings"
)

// This is the name/value of the options provided to the plugin-generator plugin. It is passed in to this
// generator plugin as a slice of this
type inputOptions struct {
	Name  string
	Value string
}

// go:embed templates/**/*.tpl
// var templatesFS embed.FS

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

	input := pdk.Input()

	s := make([]inputOptions, 0)

	err := json.Unmarshal(input, &s)
	if nil != err {
		pdk.Log(pdk.LogDebug, "Error with unmarshal of input: %s"+err.Error())
		return 1
	}

	var outputPath, language string

	for _, d := range s {
		switch d.Name {
		case "lang":
			language = strings.ToLower(d.Value)
		case "out":
			outputPath = strings.ToLower(d.Value)
		}
	}

	// if the language is Go, lets generate
	if language == "go" {
		/*
			template := pcg.LookupTemplate(g.templates, "router.tpl")
			var builder bytes.Buffer

		*/
	}

	pdk.Log(pdk.LogDebug, "Generating for language: "+language)
	pdk.Log(pdk.LogDebug, "Generating to path: "+outputPath)

	return 0
}

//export start
func start() int32 {
	hostfuncs.RegisterPlugin("spirefy.plugin-generator-go", "Spirefy Plugin Golang Generator", "1.0.0", "1.0.0", "A plugin that provides an extension point for other plugins to extend that add plugin template generation which adhere to the Extism + Plugin Engine requirements.", nil, getExtensions())
	return 0
}

func main() {}
