package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {
	var templateText = "{\"script\": " +
		"{\"source\": \"" +
		" if (ctx._source.case_description != {{.CaseDescription}} ||" +
		"  ctx._source.vik_case_number != {{.VikCaseNumber}} ||" +
		"  ctx._source.alien_registration_number != params.alien_registration_number ||" +
		"  ctx._source.bvh_registration_number != params.bvh_registration_number) " +
		"  { ctx._source.updated_datetime = ctx._now } " +
		"  ctx._source.case_description = params.case_description; " +
		"  ctx._source.vik_case_number = params.vik_case_number; " +
		"  ctx._source.alien_registration_number = params.alien_registration_number; " +
		"  ctx._source.bvh_registration_number = params.bvh_registration_number;\"," +
		" \"params\": {" +
		"  \"case_description\": \"{{.CaseDescription}}\"," +
		"  \"vik_case_number\": \"{{.VikCaseNumber}}\"," +
		"  \"share\": \"{{.Share}}\"," +
		"  \"alien_registration_number\": \"vreemdeling registratie nummer\"," +
		"  \"bvh_registration_number\": \"bvh registratie nummer\"}" +
		" }," +
		" \"upsert\": {" +
		"  \"case_description\": \"{{.CaseDescription}}\"," +
		"  \"vik_case_number\": \"{{.VikCaseNumber}}\"" +
		" }" +
		"}"

	t, err := template.New("case_update").Parse(templateText)
	if err != nil {
		panic(err)
	}
	data := make(map[string]any)
	data["CaseDescription"] = "zaak omschrijving"
	data["VikCaseNumber"] = "vik zaak nummer"
	data["Share"] = `\\\\pdczsn0010.digi.intern\\Evidence`

	w := bytes.Buffer{}
	err = t.Execute(&w, data)
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())
}