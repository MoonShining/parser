package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func BenchmarkParser(b *testing.B) {
	str := `a != 1 && c == "666" && !(d != nil) && (e > 3 || f <5)`
	is := antlr.NewInputStream(str)
	l := NewcondLexer(is)
	stream := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)
	p := NewcondParser(stream)
	matcher := NewVisitorMatcher(p.ConditionExpr())

	vars := &mapVars{m: map[string]interface{}{
		"a": int64(3),
		"c": "666",
		"e": int64(3),
		"f": int64(4),
	}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matcher.SetVars(vars)
		matcher.Match()
	}
}

type onecase struct {
	input  string
	result bool
}

type onevarcase struct {
	input  string
	result bool
	vars   *mapVars
}

func TestBaseParser(t *testing.T) {
	cases := []onecase{
		{"true", true},
		{"false", false},
		{`"hello"`, true},
		{`""`, false},
		{"1", true},
		{"0", false},
		{"null", false},
		{"nil", false}, // not exist id
	}

	for i, c := range cases {
		cond, err := Parse(c.input)
		if err != nil {
			t.Fatal(err)
		}
		matcher := NewVisitorMatcher(cond)
		if matcher.Match() != c.result {
			t.Fatalf("case idx:%d not match", i)
		}
	}
}

func TestLogicParser(t *testing.T) {
	cases := []onecase{
		{"true || false", true},
		{"false && false", false},
		{`"hello" && false`, false},
		{`"" || 1`, true},
		{`null || "1"`, true},
		{"nil || 1234", true},
	}

	for i, c := range cases {
		cond, _ := Parse(c.input)
		matcher := NewVisitorMatcher(cond)
		if matcher.Match() != c.result {
			t.Fatalf("case idx:%d not match", i)
		}
	}
}

func TestComplexParser(t *testing.T) {
	cases := []onevarcase{
		{
			`a == 1 && b == "2" && c == null && (d>5 || e <= 10 && f != "g")`,
			true,
			&mapVars{
				map[string]interface{}{
					"a": int64(1),
					"b": "2",
					"c": nil,
					"d": int64(1),
					"e": int64(9),
					"f": "k",
				},
			},
		},
	}

	for i, c := range cases {
		cond, _ := Parse(c.input)
		matcher := NewVisitorMatcher(cond)
		matcher.SetVars(c.vars)
		if matcher.Match() != c.result {
			t.Fatalf("case idx:%d not match", i)
		}
	}
}

func TestInvalidInputParser(t *testing.T) {
	cases := []onecase{
		{"a & 1", false}, // invalid token
		{"a &&", false},  // invalid grammar
	}

	for _, c := range cases {
		_, err := Parse(c.input)
		if err != nil {
			t.Log(err)
			continue
		}

		t.Fatalf("expect error but not happend")
	}
}
