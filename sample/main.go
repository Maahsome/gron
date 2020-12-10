package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/maahsome/gron"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

func main() {

	// Gron the data
	jsonData := `{
"book":[
	{
		"id":"444",
		"language":"C",
		"edition":"First",
		"author":"Dennis Ritchie"
	},
	{
		"id":"555",
		"language":"C++",
		"edition":"Second",
		"author":"Bjarne Stroustrup"
	}
]
}
`
	reader := strings.NewReader(jsonData)
	out := &bytes.Buffer{}
	// ge := gron.NewGron(reader, out)
	ge := gron.NewGron(reader, colorable.NewColorableStdout())
	err := ge.ToGron()
	if err != nil {
		logrus.Error("Problem generating gron syntax", err)
	}
	//fmt.Println(out)

	// Ungron the data
	gronData := `json = {};
json.book = [];
json.book[0] = {};
json.book[0].author = "Dennis Ritchie";
json.book[0].edition = "First";
json.book[0].id = "444";
json.book[0].language = "C";
json.book[1] = {};
json.book[1].author = "Bjarne Stroustrup";
json.book[1].edition = "second";
json.book[1].id = "555";
json.book[1].language = "C++";
`

	reader = strings.NewReader(gronData)
	out = &bytes.Buffer{}
	ge = gron.NewGron(reader, out)
	ge.SetMonochrome(true)
	err = ge.ToJSON()
	if err != nil {
		logrus.Error("Problem generating gron syntax", err)
	}

	fmt.Println(out)

	ydata, err := yaml.JSONToYAML(out.Bytes())
	if err != nil {
		logrus.Error("YAML Conversion Error", err)
	}
	fmt.Println(string(ydata[:]))

}
