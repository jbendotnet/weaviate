//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// +build ignore

// The following directive is necessary to make the package coherent:
// This program generates data.go.
//
//

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Generating deprecations code")

	fd, err := os.Open("deprecations.yml")
	fatal(err)

	var deprecations struct {
		// yaml tags not working on the go-swagger model, so we need to do the
		// map[string]interface{} workaround
		Deprecations []map[string]interface{} `yaml:"deprecations"`
	}
	err = yaml.NewDecoder(fd).Decode(&deprecations)
	fatal(err)

	parsed, err := parseDeprecations(deprecations.Deprecations)
	fatal(err)

	f, err := os.Create("data.go")
	fatal(err)
	defer f.Close()
	now := time.Now()
	err = packageTemplate.Execute(f, struct {
		Timestamp    string
		Deprecations []models.Deprecation
	}{
		Timestamp:    fmt.Sprintf("%04d-%02d-%02d", now.Year(), now.Month(), now.Day()),
		Deprecations: parsed,
	})
	fatal(err)
}

func parseDeprecations(in []map[string]interface{}) ([]models.Deprecation, error) {
	out := make([]models.Deprecation, len(in))

	for i, d := range in {
		out[i] = models.Deprecation{
			ID:                    d["id"].(string),
			Status:                d["status"].(string),
			APIType:               d["apiType"].(string),
			Msg:                   d["msg"].(string),
			Mitigation:            d["mitigation"].(string),
			SinceVersion:          d["sinceVersion"].(string),
			PlannedRemovalVersion: d["plannedRemovalVersion"].(string),
			Locations:             parseStringSlice(d["locations"].([]interface{})),
			SinceTime:             timeMust(time.Parse(time.RFC3339, d["sinceTime"].(string))),
		}

		if t, ok := d["removedTime"]; ok && t != nil {
			parsed := timeMust(time.Parse(time.RFC3339, t.(string)))
			out[i].RemovedTime = &parsed
		}

		if v, ok := d["removedIn"]; ok && v != nil {
			parsed := v.(string)
			out[i].RemovedIn = &parsed
		}
	}

	return out, nil
}

func timeMust(t time.Time, err error) strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	return strfmt.DateTime(t)
}

func parseStringSlice(in []interface{}) []string {
	out := make([]string, len(in))
	for i, elem := range in {
		out[i] = elem.(string)
	}

	return out
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Funcs(
	template.FuncMap{
		"DerefString": func(i *string) string { return *i },
	},
).Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by go generate ./deprecations at {{ .Timestamp }}
package deprecations

import (
	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"time"
)

func timeMust(t time.Time, err error) strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	return strfmt.DateTime(t)
}

func timeMustPtr(t time.Time, err error) *strfmt.DateTime {
	if err != nil {
		panic(err)
	}

	parsed := strfmt.DateTime(t)
	return &parsed
}

func ptString(in string) *string {
	return &in
}

var ByID = map[string]models.Deprecation{
{{- range .Deprecations }}
	{{ printf "%q" .ID }}: models.Deprecation{ 
		ID: {{ printf "%q" .ID }},
		Status: {{ printf "%q" .Status }},
		APIType: {{ printf "%q" .APIType }},
		Mitigation: {{ printf "%q" .Mitigation }},
		Msg: {{ printf "%q" .Msg }},
		SinceVersion: {{ printf "%q" .SinceVersion }},
		SinceTime: timeMust(time.Parse(time.RFC3339, {{ printf "%q" .SinceTime }})),
		{{ if .RemovedIn -}}
		RemovedIn: ptString({{ printf "%q" (DerefString .RemovedIn) }}),
		{{- end }}
		{{ if .RemovedTime -}}
		RemovedTime: timeMustPtr(time.Parse(time.RFC3339, {{ printf "%q" .SinceTime }})),
		{{- end }}
	},
{{- end }}
}
`))
