package goreact

import (
	"github.com/mamaar/risotto/generator"
	"github.com/robertkrimen/otto"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// MustStr takes a string and error as input
// It will log fatal if err is not nil.
func MustStr(str string, err error) string {
	if err != nil {
		log.Fatal(err)
	}
	return str
}

// RequireReact works like Require except it expects only
// the file path to react.js library. It will wrap the script
// file with javascripts to make React suitable for otto to run with.
func RequireReact(fn string) (js string, err error) {
	react, err := requireJs(fn)
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

// Require reads a script file and returns its content.
// If the file is a .jsx file, it will try to render as javascript
// (with `github.com/mamaar/risotto/generator`)
func Require(fn string) (js string, err error) {
	ext := filepath.Ext(fn)
	if ext == ".js" {
		js = MustStr(requireJs(fn)) + "\n"
	} else if ext == ".jsx" {
		js = MustStr(requireJsx(fn)) + "\n"
	} else {
		err = fmt.Errorf("Unknown script file extension \"%s\"", ext)
	}
	return
}

// JSON marshals input parameters map into json string
func JSON(m interface{}) (out string, err error) {
	b, err := json.Marshal(m)
	if err != nil {
		return
	}
	out = string(b)
	return
}

// RenderElemToString renders a ReactElement to string to HTML string
func RenderElemToString(elm string, jsonIn string, scripts ...string) (output string, err error) {
	var script string
	script += strings.Join(scripts, "\n")
	script += fmt.Sprintf(
		"var _result = React.renderToString(React.createFactory(%s)(%s))",
		elm, jsonIn)

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
