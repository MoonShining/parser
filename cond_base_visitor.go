// Code generated from cond.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // cond

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasecondVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasecondVisitor) VisitConditionExpr(ctx *ConditionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitNil(ctx *NilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitBool(ctx *BoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitPrefix(ctx *PrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasecondVisitor) VisitLogic(ctx *LogicContext) interface{} {
	return v.VisitChildren(ctx)
}
