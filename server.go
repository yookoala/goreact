package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func mustStr(str string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	return str
}

func requireReact() (js string, err error) {
	react, err := requireJs("static/js/react-0.10.0.js")
	if err != nil {
		return
	}
	js = "var self = {};\n" + react + "\nvar React = self.React;"
	return
}

func requireJs(fn string) (js string, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	js = string(b)
	return
}

func renderReact(fns ...string) (output string) {
	output += mustStr(requireReact()) + "\n"
	for _, fn := range fns {
		output += mustStr(requireJs(fn)) + "\n"
	}
	return
}

func renderReactComp(elm, out string, fns ...string) (output string) {
	output += renderReact(fns...)
	output += fmt.Sprintf("var %s = React.renderComponentToString(%s({}))",
		out, elm)
	return
}

func main() {
}
