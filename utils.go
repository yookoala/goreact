package goreact

import (
	"github.com/mamaar/risotto/generator"
	"github.com/robertkrimen/otto"

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

func renderElemToString(elm string, fns ...string) (output string, err error) {
	var script string
	script += renderReact(fns...)
	script += fmt.Sprintf(
		"var _result = React.renderToString(React.createFactory(%s)({}))", elm)

	vm := otto.New()
	_, err = vm.Run(script)
	if err != nil {
		return
	}

	val, err := vm.Get("_result")
	if err != nil {
		return
	}

	return val.ToString()
}
