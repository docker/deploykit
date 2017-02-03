package template

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunTemplateWithJMESPath(t *testing.T) {

	// Example from http://jmespath.org/
	str := `{{ q "locations[?state == 'WA'].name | sort(@) | {WashingtonCities: join(', ', @)}" . | to_json}}`

	tpl, err := NewTemplate("str://"+str, Options{})
	require.NoError(t, err)

	view, err := tpl.Render(map[string]interface{}{
		"locations": []map[string]interface{}{
			{"name": "Seattle", "state": "WA"},
			{"name": "New York", "state": "NY"},
			{"name": "Bellevue", "state": "WA"},
			{"name": "Olympia", "state": "WA"},
		},
	})

	require.NoError(t, err)
	expected := `{
  "WashingtonCities": "Bellevue, Olympia, Seattle"
}`
	require.Equal(t, expected, view)
}

func TestVarAndGlobal(t *testing.T) {
	str := `{{ q "locations[?state == 'WA'].name | sort(@) | {WashingtonCities: join(', ', @)}" . | global "washington-cities"}}

{{/* The query above is exported and referenced somewhere else */}}
{{ from_json "[\"SF\",\"LA\"]" | def "california-cities" "Default value for California cities" }}
{{ from_json "{\"SF\":\"94109\",\"LA\":\"90210\"}" | def "zip-codes" "Default value for zip codes" }}

{
  "test" : "hello",
  "val"  : true,
  "result" : {{ref "washington-cities" | to_json}},
  "california" : {{ ref "california-cities" | to_json}},
  "sf_zip" : {{ ref "zip-codes" | q "SF" | to_json }}
}
`

	tpl, err := NewTemplate("str://"+str, Options{})
	require.NoError(t, err)

	view, err := tpl.Render(map[string]interface{}{
		"locations": []map[string]interface{}{
			{"name": "Seattle", "state": "WA"},
			{"name": "New York", "state": "NY"},
			{"name": "Bellevue", "state": "WA"},
			{"name": "Olympia", "state": "WA"},
		},
	})

	require.NoError(t, err)

	// Note the extra newlines because of comments, etc.
	expected := `





{
  "test" : "hello",
  "val"  : true,
  "result" : {
  "WashingtonCities": "Bellevue, Olympia, Seattle"
},
  "california" : [
  "SF",
  "LA"
],
  "sf_zip" : "94109"
}
`
	require.Equal(t, expected, view)

}

type context struct {
	Count  int
	Bool   bool
	String string

	invokes int
}

func (s *context) SetBool(b bool) {
	s.invokes++
	s.Bool = b
}

func (s *context) Funcs() []Function {
	return []Function{
		{
			Name:        "inc",
			Description: "increments the context counter when called",
			Func: func(c Context) interface{} {
				c.(*context).invokes++
				c.(*context).Count++
				return ""
			},
		},
		{
			Name:        "dec",
			Description: "decrements the context counter when called",
			Func: func(s Context) interface{} {
				s.(*context).invokes++
				s.(*context).Count--
				return ""
			},
		},
		{
			Name:        "str",
			Description: "prints the string",
			Func: func(c Context) string {
				c.(*context).invokes++
				return s.String
			},
		},
		{
			Name:        "count",
			Description: "prints the count",
			Func: func(c Context) int {
				c.(*context).invokes++
				return s.Count
			},
		},
		{
			Name:        "setString",
			Description: "sets the string",
			Func: func(c Context, n string) interface{} {
				c.(*context).invokes++
				s.String = n
				return ""
			},
		},
		{
			Name:        "setBool",
			Description: "sets the bool",
			Func: func(c Context, b bool) bool {
				c.(*context).SetBool(b)
				return c.(*context).Bool
			},
		},
		{
			Name:        "invokes",
			Description: "prints the invokes count",
			Func: func() int {
				s.invokes++
				return s.invokes
			},
		},
	}
}

func TestContextFuncs(t *testing.T) {

	_, err := makeTemplateFunc(&context{}, "string")
	require.Error(t, err)

	input := make(chan string, 2)
	func1 := func(a string) string {
		input <- a
		return a
	}

	tf, err := makeTemplateFunc(&context{}, func1)
	require.NoError(t, err)
	require.Equal(t, func1("x"), tf.(func(string) string)("x"))
	require.Equal(t, "x", <-input)
	require.Equal(t, "x", <-input)

	input2 := make(chan string, 2)
	func2 := func(ctx Context, a string) string {
		input2 <- a
		return a
	}
	tf, err = makeTemplateFunc(&context{}, func2)
	require.NoError(t, err)
	require.Equal(t, func2(&context{}, "x"), tf.(func(string) string)("x"))
	require.Equal(t, "x", <-input2)
	require.Equal(t, "x", <-input2)

	input3 := make(chan string, 2)
	input4 := make(chan bool, 2)
	func3 := func(ctx Context, a string, opt ...bool) (string, error) {
		input3 <- a
		input4 <- len(opt) > 0 && opt[0]
		return a, nil
	}
	tf, err = makeTemplateFunc(&context{}, func3)
	require.NoError(t, err)

	o, e := func3(&context{}, "x")
	require.NoError(t, e)
	oo, ee := tf.(func(string, ...bool) (string, error))("x")
	require.NoError(t, ee)

	require.Equal(t, o, oo)

	require.Equal(t, "x", <-input3)
	require.Equal(t, "x", <-input3)
	require.False(t, <-input4)
	require.False(t, <-input4)

	input5 := make(chan string, 1)
	input6 := make(chan bool, 1)
	func4 := func(ctx Context, a, b string, opt ...bool) (string, error) {
		input5 <- a
		input6 <- len(opt) > 0 && opt[0]
		ctx.(*context).Count++
		if a == b {
			ctx.(*context).Bool = true
		}
		return a, nil
	}

	s := &context{}
	tf, err = makeTemplateFunc(s, func4)
	require.NoError(t, err)

	oo, ee = tf.(func(string, string, ...bool) (string, error))("x", "x")
	require.NoError(t, ee)

	require.Equal(t, "x", <-input5)
	require.False(t, <-input6)

	require.True(t, s.Bool)
	require.Equal(t, 1, s.Count)

	s = &context{}
	tf, err = makeTemplateFunc(s, s.Funcs()[0].Func)
	require.NoError(t, err)

	for range []int{1, 1, 1, 1} {
		tf.(func() interface{})()
	}

	require.Equal(t, 4, s.Count)
}
