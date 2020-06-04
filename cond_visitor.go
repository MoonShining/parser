// Code generated from cond.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // cond

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by condParser.
type condVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by condParser#conditionExpr.
	VisitConditionExpr(ctx *ConditionExprContext) interface{}

	// Visit a parse tree produced by condParser#nil.
	VisitNil(ctx *NilContext) interface{}

	// Visit a parse tree produced by condParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by condParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by condParser#paren.
	VisitParen(ctx *ParenContext) interface{}

	// Visit a parse tree produced by condParser#bool.
	VisitBool(ctx *BoolContext) interface{}

	// Visit a parse tree produced by condParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by condParser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by condParser#prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by condParser#logic.
	VisitLogic(ctx *LogicContext) interface{}
}
