// Code generated from cond.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // cond

import "github.com/antlr/antlr4/runtime/Go/antlr"

// condListener is a complete listener for a parse tree produced by condParser.
type condListener interface {
	antlr.ParseTreeListener

	// EnterConditionExpr is called when entering the conditionExpr production.
	EnterConditionExpr(c *ConditionExprContext)

	// EnterNil is called when entering the nil production.
	EnterNil(c *NilContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterParen is called when entering the paren production.
	EnterParen(c *ParenContext)

	// EnterBool is called when entering the bool production.
	EnterBool(c *BoolContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterLogic is called when entering the logic production.
	EnterLogic(c *LogicContext)

	// ExitConditionExpr is called when exiting the conditionExpr production.
	ExitConditionExpr(c *ConditionExprContext)

	// ExitNil is called when exiting the nil production.
	ExitNil(c *NilContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitParen is called when exiting the paren production.
	ExitParen(c *ParenContext)

	// ExitBool is called when exiting the bool production.
	ExitBool(c *BoolContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitLogic is called when exiting the logic production.
	ExitLogic(c *LogicContext)
}
