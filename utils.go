package goreact

import (
	"github.com/mamaar/risotto/generator"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func mustStr(str string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	return str
}

func requireReact() (js string, err error) {
	react, err := requireJs("_test/js/react-0.13.3.js")
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

func requireJsx(fn string) (js string, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return
	}
	gen, err := generator.ParseAndGenerate(f)
	if err != nil {
		return
	}
	b, err := ioutil.ReadAll(gen)
	if err != nil {
		return
	}
	js = string(b) + "\n"
	return
}

func renderReact(fns ...string) (output string) {
	output += mustStr(requireReact()) + "\n"
	for _, fn := range fns {
		ext := filepath.Ext(fn)
		if ext == ".js" {
			output += mustStr(requireJs(fn)) + "\n"
		} else if ext == ".jsx" {
			output += mustStr(requireJsx(fn)) + "\n"
		} else {
			log.Fatal("Unknown script file extension " + ext)
		}
	}
	return
}

func renderReactElem(elm, out string, fns ...string) (output string) {
	output += renderReact(fns...)
	output += fmt.Sprintf("var %s = React.renderToString(React.createFactory(%s)({}))",
		out, elm)
	return
}
