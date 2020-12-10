package gron

import (
	"io"
)

var gronVersion = "dev"

// Exit codes
const (
	exitOK = iota
	exitOpenFile
	exitReadInput
	exitFormStatements
	exitFetchURL
	exitParseStatements
	exitJSONEncode
)

// Option bitfields
const (
	optMonochrome = 1 << iota
	optNoSort
	optJSON
	optOnlyData
)

// Gron Interface Parameters
type Gron struct {
	input        io.Reader
	output       io.Writer
	noSort       bool
	monochrome   bool
	asJSONStream bool
	onlyData     bool
}

// NewGron - Create a new processor
func NewGron(reader io.Reader, writer io.Writer) *Gron {
	g := &Gron{
		input:        reader,
		output:       writer,
		noSort:       false,
		monochrome:   false,
		asJSONStream: false,
		onlyData:     false,
	}
	return g
}

// // SetWriter - Set the io.Writer interface
// func (g *Gron) SetWriter(writer io.Writer) {
// 	g.output = writer
// }

// SetNoSort - Set to suppress sorting
func (g *Gron) SetNoSort(sort bool) {
	g.noSort = sort
}

// SetMonochrome - Set output to no color
func (g *Gron) SetMonochrome(mono bool) {
	g.monochrome = mono
}

// SetJSONStream - Set output to be wrapped as JSON
func (g *Gron) SetJSONStream(json bool) {
	g.asJSONStream = json
}

// SetOnlyData - Set output to be wrapped as JSON
func (g *Gron) SetOnlyData(onlydata bool) {
	g.onlyData = onlydata
}

func calculateOptions(g *Gron) int {
	var opts int
	if g.monochrome {
		opts = opts | optMonochrome
	}
	if g.noSort {
		opts = opts | optNoSort
	}
	if g.asJSONStream {
		opts = opts | optJSON
	}
	if g.onlyData {
		opts = opts | optOnlyData
	}
	return opts
}

// ToGron - Output JSON/YAML in Gron Notation
func (g *Gron) ToGron() error {
	_, err := gron(g.input, g.output, calculateOptions(g))
	return err
}

// ToJSON - Output Gron in JSON
func (g *Gron) ToJSON() error {
	_, err := ungron(g.input, g.output, calculateOptions(g))
	return err
}

// // ToYaml - Output Gron in YAML
// func (g *Gron) ToYaml() string {
// 	_, jdata, err := ungron(g.input, colorable.NewColorableStdout(), calculateOptions(g))
// 	if err != nil {
// 		logrus.Error("Return error", err)
// 	}

// 	ydata, err := yaml.JSONToYAML([]byte(jdata))
// 	if err != nil {
// 		logrus.Error("YAML Conversion Error", err)
// 	}

// 	return string(ydata[:])

// }
