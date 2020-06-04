package parser

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

func Parse(in string) (IConditionExprContext, error) {
	el := &errListener{}
	is := antlr.NewInputStream(in)
	l := NewcondLexer(is)
	l.RemoveErrorListeners()
	l.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)
	p := NewcondParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(el)
	condCtx := p.ConditionExpr()
	if len(el.errMsg) > 0 {
		return nil, errors.New(strings.Join(el.errMsg, "."))
	}
	if el.errors > 0 {
		return nil, errors.New("unknown syntax error")
	}
	return condCtx, nil
}
