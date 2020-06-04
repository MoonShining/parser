// Code generated from cond.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // cond

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 34, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 21, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 2, 3, 4,
	4, 2, 4, 2, 4, 3, 2, 4, 11, 3, 2, 12, 13, 2, 39, 2, 6, 3, 2, 2, 2, 4, 20,
	3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 3, 3, 2, 2, 2, 8, 9, 8, 3, 1, 2, 9, 10,
	7, 3, 2, 2, 10, 21, 5, 4, 3, 11, 11, 21, 7, 18, 2, 2, 12, 21, 7, 16, 2,
	2, 13, 21, 7, 17, 2, 2, 14, 21, 7, 19, 2, 2, 15, 21, 7, 20, 2, 2, 16, 17,
	7, 14, 2, 2, 17, 18, 5, 2, 2, 2, 18, 19, 7, 15, 2, 2, 19, 21, 3, 2, 2,
	2, 20, 8, 3, 2, 2, 2, 20, 11, 3, 2, 2, 2, 20, 12, 3, 2, 2, 2, 20, 13, 3,
	2, 2, 2, 20, 14, 3, 2, 2, 2, 20, 15, 3, 2, 2, 2, 20, 16, 3, 2, 2, 2, 21,
	30, 3, 2, 2, 2, 22, 23, 12, 10, 2, 2, 23, 24, 9, 2, 2, 2, 24, 29, 5, 4,
	3, 11, 25, 26, 12, 9, 2, 2, 26, 27, 9, 3, 2, 2, 27, 29, 5, 4, 3, 10, 28,
	22, 3, 2, 2, 2, 28, 25, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28, 3, 2, 2,
	2, 30, 31, 3, 2, 2, 2, 31, 5, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 5, 20, 28,
	30,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'!'", "'<'", "'>'", "'<='", "'>='", "'=='", "'!='", "'in'", "'not in'",
	"'&&'", "'||'", "'('", "')'", "", "'null'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "BOOL", "NULL",
	"NUMBER", "ID", "STRING", "WS",
}

var ruleNames = []string{
	"conditionExpr", "singleExpr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type condParser struct {
	*antlr.BaseParser
}

func NewcondParser(input antlr.TokenStream) *condParser {
	this := new(condParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "cond.g4"

	return this
}

// condParser tokens.
const (
	condParserEOF    = antlr.TokenEOF
	condParserT__0   = 1
	condParserT__1   = 2
	condParserT__2   = 3
	condParserT__3   = 4
	condParserT__4   = 5
	condParserT__5   = 6
	condParserT__6   = 7
	condParserT__7   = 8
	condParserT__8   = 9
	condParserT__9   = 10
	condParserT__10  = 11
	condParserT__11  = 12
	condParserT__12  = 13
	condParserBOOL   = 14
	condParserNULL   = 15
	condParserNUMBER = 16
	condParserID     = 17
	condParserSTRING = 18
	condParserWS     = 19
)

// condParser rules.
const (
	condParserRULE_conditionExpr = 0
	condParserRULE_singleExpr    = 1
)

// IConditionExprContext is an interface to support dynamic dispatch.
type IConditionExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionExprContext differentiates from other interfaces.
	IsConditionExprContext()
}

type ConditionExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionExprContext() *ConditionExprContext {
	var p = new(ConditionExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = condParserRULE_conditionExpr
	return p
}

func (*ConditionExprContext) IsConditionExprContext() {}

func NewConditionExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionExprContext {
	var p = new(ConditionExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = condParserRULE_conditionExpr

	return p
}

func (s *ConditionExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionExprContext) SingleExpr() ISingleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *ConditionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterConditionExpr(s)
	}
}

func (s *ConditionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitConditionExpr(s)
	}
}

func (s *ConditionExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitConditionExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *condParser) ConditionExpr() (localctx IConditionExprContext) {
	localctx = NewConditionExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, condParserRULE_conditionExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.singleExpr(0)
	}

	return localctx
}

// ISingleExprContext is an interface to support dynamic dispatch.
type ISingleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExprContext differentiates from other interfaces.
	IsSingleExprContext()
}

type SingleExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExprContext() *SingleExprContext {
	var p = new(SingleExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = condParserRULE_singleExpr
	return p
}

func (*SingleExprContext) IsSingleExprContext() {}

func NewSingleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExprContext {
	var p = new(SingleExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = condParserRULE_singleExpr

	return p
}

func (s *SingleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExprContext) CopyFrom(ctx *SingleExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NilContext struct {
	*SingleExprContext
}

func NewNilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilContext {
	var p = new(NilContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) NULL() antlr.TerminalNode {
	return s.GetToken(condParserNULL, 0)
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*SingleExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(condParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierContext struct {
	*SingleExprContext
}

func NewIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierContext {
	var p = new(IdentifierContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(condParserID, 0)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	*SingleExprContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) ConditionExpr() IConditionExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionExprContext)
}

func (s *ParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterParen(s)
	}
}

func (s *ParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitParen(s)
	}
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolContext struct {
	*SingleExprContext
}

func NewBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolContext {
	var p = new(BoolContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(condParserBOOL, 0)
}

func (s *BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterBool(s)
	}
}

func (s *BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitBool(s)
	}
}

func (s *BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	*SingleExprContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(condParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type TestContext struct {
	*SingleExprContext
	left  ISingleExprContext
	op    antlr.Token
	right ISingleExprContext
}

func NewTestContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TestContext {
	var p = new(TestContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *TestContext) GetOp() antlr.Token { return s.op }

func (s *TestContext) SetOp(v antlr.Token) { s.op = v }

func (s *TestContext) GetLeft() ISingleExprContext { return s.left }

func (s *TestContext) GetRight() ISingleExprContext { return s.right }

func (s *TestContext) SetLeft(v ISingleExprContext) { s.left = v }

func (s *TestContext) SetRight(v ISingleExprContext) { s.right = v }

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) AllSingleExpr() []ISingleExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExprContext)(nil)).Elem())
	var tst = make([]ISingleExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExprContext)
		}
	}

	return tst
}

func (s *TestContext) SingleExpr(i int) ISingleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *TestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterTest(s)
	}
}

func (s *TestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitTest(s)
	}
}

func (s *TestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitTest(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixContext struct {
	*SingleExprContext
	prefix antlr.Token
	expr   ISingleExprContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *PrefixContext) GetPrefix() antlr.Token { return s.prefix }

func (s *PrefixContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *PrefixContext) GetExpr() ISingleExprContext { return s.expr }

func (s *PrefixContext) SetExpr(v ISingleExprContext) { s.expr = v }

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) SingleExpr() ISingleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (s *PrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type LogicContext struct {
	*SingleExprContext
	left  ISingleExprContext
	op    antlr.Token
	right ISingleExprContext
}

func NewLogicContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicContext {
	var p = new(LogicContext)

	p.SingleExprContext = NewEmptySingleExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExprContext))

	return p
}

func (s *LogicContext) GetOp() antlr.Token { return s.op }

func (s *LogicContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogicContext) GetLeft() ISingleExprContext { return s.left }

func (s *LogicContext) GetRight() ISingleExprContext { return s.right }

func (s *LogicContext) SetLeft(v ISingleExprContext) { s.left = v }

func (s *LogicContext) SetRight(v ISingleExprContext) { s.right = v }

func (s *LogicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicContext) AllSingleExpr() []ISingleExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExprContext)(nil)).Elem())
	var tst = make([]ISingleExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExprContext)
		}
	}

	return tst
}

func (s *LogicContext) SingleExpr(i int) ISingleExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExprContext)
}

func (s *LogicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.EnterLogic(s)
	}
}

func (s *LogicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(condListener); ok {
		listenerT.ExitLogic(s)
	}
}

func (s *LogicContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case condVisitor:
		return t.VisitLogic(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *condParser) SingleExpr() (localctx ISingleExprContext) {
	return p.singleExpr(0)
}

func (p *condParser) singleExpr(_p int) (localctx ISingleExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, condParserRULE_singleExpr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case condParserT__0:
		localctx = NewPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(7)

			var _m = p.Match(condParserT__0)

			localctx.(*PrefixContext).prefix = _m
		}
		{
			p.SetState(8)

			var _x = p.singleExpr(9)

			localctx.(*PrefixContext).expr = _x
		}

	case condParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(condParserNUMBER)
		}

	case condParserBOOL:
		localctx = NewBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(condParserBOOL)
		}

	case condParserNULL:
		localctx = NewNilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(condParserNULL)
		}

	case condParserID:
		localctx = NewIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(condParserID)
		}

	case condParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.Match(condParserSTRING)
		}

	case condParserT__11:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(condParserT__11)
		}
		{
			p.SetState(15)
			p.ConditionExpr()
		}
		{
			p.SetState(16)
			p.Match(condParserT__12)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewTestContext(p, NewSingleExprContext(p, _parentctx, _parentState))
				localctx.(*TestContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, condParserRULE_singleExpr)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(21)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*TestContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<condParserT__1)|(1<<condParserT__2)|(1<<condParserT__3)|(1<<condParserT__4)|(1<<condParserT__5)|(1<<condParserT__6)|(1<<condParserT__7)|(1<<condParserT__8))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*TestContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(22)

					var _x = p.singleExpr(9)

					localctx.(*TestContext).right = _x
				}

			case 2:
				localctx = NewLogicContext(p, NewSingleExprContext(p, _parentctx, _parentState))
				localctx.(*LogicContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, condParserRULE_singleExpr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(24)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LogicContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == condParserT__9 || _la == condParserT__10) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LogicContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(25)

					var _x = p.singleExpr(8)

					localctx.(*LogicContext).right = _x
				}

			}

		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *condParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *SingleExprContext = nil
		if localctx != nil {
			t = localctx.(*SingleExprContext)
		}
		return p.SingleExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *condParser) SingleExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
