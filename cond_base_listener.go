// Code generated from cond.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // cond

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecondListener is a complete listener for a parse tree produced by condParser.
type BasecondListener struct{}

var _ condListener = &BasecondListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecondListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecondListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecondListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecondListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConditionExpr is called when production conditionExpr is entered.
func (s *BasecondListener) EnterConditionExpr(ctx *ConditionExprContext) {}

// ExitConditionExpr is called when production conditionExpr is exited.
func (s *BasecondListener) ExitConditionExpr(ctx *ConditionExprContext) {}

// EnterNil is called when production nil is entered.
func (s *BasecondListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production nil is exited.
func (s *BasecondListener) ExitNil(ctx *NilContext) {}

// EnterNumber is called when production number is entered.
func (s *BasecondListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasecondListener) ExitNumber(ctx *NumberContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasecondListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasecondListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterParen is called when production paren is entered.
func (s *BasecondListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production paren is exited.
func (s *BasecondListener) ExitParen(ctx *ParenContext) {}

// EnterBool is called when production bool is entered.
func (s *BasecondListener) EnterBool(ctx *BoolContext) {}

// ExitBool is called when production bool is exited.
func (s *BasecondListener) ExitBool(ctx *BoolContext) {}

// EnterString is called when production string is entered.
func (s *BasecondListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BasecondListener) ExitString(ctx *StringContext) {}

// EnterTest is called when production test is entered.
func (s *BasecondListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasecondListener) ExitTest(ctx *TestContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BasecondListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BasecondListener) ExitPrefix(ctx *PrefixContext) {}

// EnterLogic is called when production logic is entered.
func (s *BasecondListener) EnterLogic(ctx *LogicContext) {}

// ExitLogic is called when production logic is exited.
func (s *BasecondListener) ExitLogic(ctx *LogicContext) {}
