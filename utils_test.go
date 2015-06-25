package goreact

import (
	"testing"
)

func testData() map[string]interface{} {
	return map[string]interface{}{
		"data": []interface{}{
			map[string]interface{}{
				"hello": "data world",
			},
		},
	}
}

func TestJs(t *testing.T) {

	// render parameters
	params := MustStr(JSON(testData()))
	t.Logf("params: %#v", params)

	// render to string
	html, err := RenderElemToString("CommentBox",
		params,
		MustStr(RequireReact("_test/js/react-0.13.3.js")),
		MustStr(Require("_test/js/main.js")))
	if err != nil {
		t.Error(err.Error())
	}
	if html == "" {
		t.Error("Empty html result")
	}
	t.Log("result: \n", html)
}

func TestJsx(t *testing.T) {

	// render parameters
	params := MustStr(JSON(testData()))
	t.Logf("params: %#v", params)

	// render to string
	html, err := RenderElemToString("CommentBox",
		params,
		MustStr(RequireReact("_test/js/react-0.13.3.js")),
		MustStr(Require("_test/js/main.jsx")))
	if err != nil {
		t.Error(err.Error())
	}
	if html == "" {
		t.Error("Empty html result")
	}
	t.Log("result: \n", html)
}
