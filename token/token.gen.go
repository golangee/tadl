// Code generated by go generate; DO NOT EDIT.
package token

const (
	TokenCharData        TokenType = "TokenCharData"
	TokenIdentifier      TokenType = "TokenIdentifier"
	TokenBlockStart      TokenType = "TokenBlockStart"
	TokenBlockEnd        TokenType = "TokenBlockEnd"
	TokenGroupStart      TokenType = "TokenGroupStart"
	TokenGroupEnd        TokenType = "TokenGroupEnd"
	TokenGenericStart    TokenType = "TokenGenericStart"
	TokenGenericEnd      TokenType = "TokenGenericEnd"
	TokenG2Preamble      TokenType = "TokenG2Preamble"
	TokenDefineElement   TokenType = "TokenDefineElement"
	TokenDefineAttribute TokenType = "TokenDefineAttribute"
	TokenAssign          TokenType = "TokenAssign"
	TokenG1LineEnd       TokenType = "TokenG1LineEnd"
	TokenComma           TokenType = "TokenComma"
	TokenSemicolon       TokenType = "TokenSemicolon"
	TokenG1Comment       TokenType = "TokenG1Comment"
	TokenG2Comment       TokenType = "TokenG2Comment"
	TokenG2Arrow         TokenType = "TokenG2Arrow"
)

func (t *CharData) TokenType() TokenType {
	return TokenCharData
}

func (t *CharData) Pos() *Position {
	return &t.Position
}

func (t *Identifier) TokenType() TokenType {
	return TokenIdentifier
}

func (t *Identifier) Pos() *Position {
	return &t.Position
}

func (t *BlockStart) TokenType() TokenType {
	return TokenBlockStart
}

func (t *BlockStart) Pos() *Position {
	return &t.Position
}

func (t *BlockEnd) TokenType() TokenType {
	return TokenBlockEnd
}

func (t *BlockEnd) Pos() *Position {
	return &t.Position
}

func (t *GroupStart) TokenType() TokenType {
	return TokenGroupStart
}

func (t *GroupStart) Pos() *Position {
	return &t.Position
}

func (t *GroupEnd) TokenType() TokenType {
	return TokenGroupEnd
}

func (t *GroupEnd) Pos() *Position {
	return &t.Position
}

func (t *GenericStart) TokenType() TokenType {
	return TokenGenericStart
}

func (t *GenericStart) Pos() *Position {
	return &t.Position
}

func (t *GenericEnd) TokenType() TokenType {
	return TokenGenericEnd
}

func (t *GenericEnd) Pos() *Position {
	return &t.Position
}

func (t *G2Preamble) TokenType() TokenType {
	return TokenG2Preamble
}

func (t *G2Preamble) Pos() *Position {
	return &t.Position
}

func (t *DefineElement) TokenType() TokenType {
	return TokenDefineElement
}

func (t *DefineElement) Pos() *Position {
	return &t.Position
}

func (t *DefineAttribute) TokenType() TokenType {
	return TokenDefineAttribute
}

func (t *DefineAttribute) Pos() *Position {
	return &t.Position
}

func (t *Assign) TokenType() TokenType {
	return TokenAssign
}

func (t *Assign) Pos() *Position {
	return &t.Position
}

func (t *G1LineEnd) TokenType() TokenType {
	return TokenG1LineEnd
}

func (t *G1LineEnd) Pos() *Position {
	return &t.Position
}

func (t *Comma) TokenType() TokenType {
	return TokenComma
}

func (t *Comma) Pos() *Position {
	return &t.Position
}

func (t *Semicolon) TokenType() TokenType {
	return TokenSemicolon
}

func (t *Semicolon) Pos() *Position {
	return &t.Position
}

func (t *G1Comment) TokenType() TokenType {
	return TokenG1Comment
}

func (t *G1Comment) Pos() *Position {
	return &t.Position
}

func (t *G2Comment) TokenType() TokenType {
	return TokenG2Comment
}

func (t *G2Comment) Pos() *Position {
	return &t.Position
}

func (t *G2Arrow) TokenType() TokenType {
	return TokenG2Arrow
}

func (t *G2Arrow) Pos() *Position {
	return &t.Position
}
