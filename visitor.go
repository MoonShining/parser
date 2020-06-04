package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

var emptyVar = &mapVars{m: map[string]interface{}{}}

type Vars interface {
	Get(string) interface{}
}

type mapVars struct {
	m map[string]interface{}
}

func (mv *mapVars) Get(key string) interface{} {
	return mv.m[key]
}

type Matcher interface {
	SetVars(Vars)
	Match() bool
}

func NewVisitorMatcher(expr IConditionExprContext) Matcher {
	return &visitor{expr: expr, vars: emptyVar}
}

type visitor struct {
	BasecondVisitor

	expr IConditionExprContext
	vars Vars
}

func (v *visitor) Match() bool {
	return toBool(v.expr.Accept(v))
}

func (v *visitor) SetVars(vars Vars) {
	v.vars = vars
}

// parse tree visitor pattern
func (v *visitor) Visit(tree antlr.ParseTree) interface{} {
	switch ctx := tree.(type) {
	case *ConditionExprContext:
		return ctx.Accept(v)
	default:
		panic("Visit won't happen")
	}
}

func (v *visitor) VisitConditionExpr(ctx *ConditionExprContext) interface{} {
	return ctx.SingleExpr().Accept(v)
}

func (v *visitor) VisitNil(ctx *NilContext) interface{} {
	return nil
}

func (v *visitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	idVal := v.vars.Get(ctx.GetText())
	return idVal
}

func (v *visitor) VisitNumber(ctx *NumberContext) interface{} {
	num, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return num
}

func (v *visitor) VisitParen(ctx *ParenContext) interface{} {
	return ctx.ConditionExpr().Accept(v)
}

func (v *visitor) VisitString(ctx *StringContext) interface{} {
	s, _ := strconv.Unquote(ctx.GetText())
	return s
}

func (v *visitor) VisitBool(ctx *BoolContext) interface{} {
	return ctx.GetText() == "true"
}

func (v *visitor) VisitTest(ctx *TestContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)

	switch ctx.GetOp().GetText() {
	case "==":
		return left == right
	case "!=":
		return left != right
	case ">":
		leftNum, ok1 := toNumber(left)
		rightNum, ok2 := toNumber(right)
		if !ok1 || !ok2 {
			return false
		} else {
			return leftNum > rightNum
		}
	case "<":
		leftNum, ok1 := toNumber(left)
		rightNum, ok2 := toNumber(right)
		if !ok1 || !ok2 {
			return false
		} else {
			return leftNum < rightNum
		}
	case "<=":
		leftNum, ok1 := toNumber(left)
		rightNum, ok2 := toNumber(right)
		if !ok1 || !ok2 {
			return false
		} else {
			return leftNum <= rightNum
		}
	case ">=":
		leftNum, ok1 := toNumber(left)
		rightNum, ok2 := toNumber(right)
		if !ok1 || !ok2 {
			return false
		} else {
			return leftNum >= rightNum
		}
	case "in":
		if list, ok := right.([]interface{}); ok {
			isIn := false
			for _, ele := range list {
				if left == ele {
					isIn = true
					break
				}
			}
			return isIn
		} else {
			return false
		}
	case "not in":
		if list, ok := right.([]interface{}); ok {
			notIn := true
			for _, ele := range list {
				if left == ele {
					notIn = false
					break
				}
			}
			return notIn
		} else {
			return false
		}
	}

	panic("VisitTest wont't happen")
}

func (v *visitor) VisitPrefix(ctx *PrefixContext) interface{} {
	b := ctx.expr.Accept(v)
	ok := toBool(b)
	return !ok
}

func (v *visitor) VisitLogic(ctx *LogicContext) interface{} {
	left := ctx.GetLeft().Accept(v)
	right := ctx.GetRight().Accept(v)

	switch ctx.GetOp().GetText() {
	case "&&":
		return toBool(left) && toBool(right)
	case "||":
		return toBool(left) || toBool(right)
	}

	panic("logic won't happen")
}

func toNumber(val interface{}) (int64, bool) {
	n, ok := val.(int64)
	return n, ok
}

func toBool(val interface{}) bool {
	switch v := val.(type) {
	case nil:
		return false
	case bool:
		return v
	case string:
		return v != ""
	case int64:
		return v != 0
	case []interface{}:
		return true
	default:
		panic("toBool won't happen")
	}
}
