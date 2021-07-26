package streamxmlencoder

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/golangee/tadl/parser"
	"github.com/golangee/tadl/token"
)

// XMLEncoder translates tadl-input to corresponding XML
type XMLEncoder struct {
	lexer      *token.Lexer
	tokens     []token.Token
	writer     io.Writer
	buffWriter *bufio.Writer

	tadlText string
	xmlText  string
	prefix   string
	postfix  string

	identifiers     []string
	postfixes       []string
	attributeBuffer []token.Token
}

const (
	lt         = "\x3C"
	equals     = "\x3D"
	gt         = "\x3E"
	dquotes    = "\x22"
	slash      = "\x2F"
	whitespace = "\x20"
)

func (x *XMLEncoder) write(in ...string) {
	for _, text := range in {
		x.buffWriter.Write([]byte(text))
	}
}

// NewEncoder creades a new XMLEncoder
// tadl-input is given as an io.Reader instance
func NewEncoder(filename string, r io.Reader, w io.Writer) XMLEncoder {
	return nil
}

// EncodeToXml uses a parser2.parser to create a syntax tree,
// utilizes the encodeRek method to translate it and returns the result
func (x *XMLEncoder) EncodeToXML() (string, error) {
	return "", errors.New("Encoding not implemented yet")
}

// encodeRek recursively translates the syntax tree
// given by its root Element to the corresponding XML.
func encodeRek(root parser.TreeNode) (string, error) {
	if root.IsComment() {
		return "<!-- " + *root.Comment + " -->", nil
	} else if root.IsText() {
		return *root.Text, nil
	} else if root.IsNode() {
		var outString, postfix string

		if root.BlockType == parser.BlockNormal || root.BlockType == parser.BlockNone {
			outString += "<" + root.Name
			postfix = "</" + root.Name + ">"
		} else if root.BlockType == parser.BlockGroup {
			outString += "<" + root.Name + ` _groupType="()"`
			postfix = "</" + root.Name + ">"
		} else if root.BlockType == parser.BlockGeneric {
			outString += "<" + root.Name + ` _groupType="<>"`
			postfix = "</" + root.Name + ">"
		}

		if len(root.Attributes) != 0 {
			for key, val := range root.Attributes {
				outString += " " + key + `="` + val + `"`
			}
		}

		outString += ">"
		if root.Name == "title" {
		}
		for _, child := range root.Children {

			fmt.Printf("root: %v, child: %v", root, child)
			var text string
			text, err := encodeRek(*child)
			if err != nil {
				return "", err
			}

			outString += text
		}

		return outString + postfix, nil
	} else {
		return "", errors.New("Token not identified, aborting encoding")
	}
}

// Tokenize creates a Slice of consecutive Tokens, representing the tadl-input syntax
func (x *XMLEncoder) Tokenize() error {
	for {
		currentToken, err := x.getNextToken()

		if err != nil {
			return err
		}
		if currentToken == nil {
			break
		}

		x.tokens = append(x.tokens, *currentToken)
	}
	return nil
}

// Next returns the next XML Translation
// to the corresponding TADL token in the input stream.
func (x *XMLEncoder) Next() error {
	token, err := x.getNextToken()
	fmt.Println("Token, err, tokentype ", token, err, (*token).TokenType())
	if err != nil {
		return err
	}

	err = x.tokenToXML(token)
	if err != nil {
		return err
	}
	x.buffWriter.Flush()
	return nil
}

// getNextToken uses a Lexer to read the next consecutive Token
func (x *XMLEncoder) getNextToken() (*token.Token, error) {
	token, err := x.lexer.Token()
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// tokenToXML encodes the given Token and writes the corresponding
// XML translation to the io.Writer in x.writer
func (x *XMLEncoder) tokenToXML(currentToken *token.Token) error {
	x.xmlText = ""
	switch (*currentToken).TokenType() {
	case token.TokenIdentifier:
		ct, _ := (*currentToken).(*token.Identifier)
		fmt.Printf("found identifier, translating to %s %s", lt, ct.Value)
		x.pushToStack(x.identifiers, ct.Value)
		x.write(lt, ct.Value)

		nextToken, err := x.getNextToken()
		if err != nil {
			return err
		}

		switch (*currentToken).TokenType() {
		case token.TokenDefineAttribute:
			_, forward := (*currentToken).(*token.DefineAttribute)

			//TODO: multiple forwarded Attributes
			if forward {
				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenIdentifier {
					return errors.New("Unexpected Token, expected Identifier")
				}

				nextTokenIdent, forward := (*nextToken).(*token.Identifier)
				if forward {
					return errors.New("Unexpected Forward, expected unforwarded Identifier")
				}
				x.write(whitespace, nextTokenIdent.Value, equals)

				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenBlockStart {
					return errors.New("Unexpected Token, expected BlockStart")
				}
				if err != nil {
					return err
				}

				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenCharData {
					return errors.New("Unexpected Token, expected Chardata")
				}
				if err != nil {
					return err
				}
				x.write(dquotes, (*nextToken).(*token.CharData).Value, dquotes)

			} else {
				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenIdentifier {
					return errors.New("Unexpected Token, expected Identifier")
				}
				if err != nil {
					return err
				}
				//identifier := nextToken.(*parser2.Identifier).Value

				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenBlockStart {
					return errors.New("Unexpected Token, expected BlockStart")
				}
				if err != nil {
					return err
				}

				if nextToken, err = x.getNextToken(); (*nextToken).TokenType() != token.TokenCharData {
					return errors.New("Unexpected Token, expected Chardata")
				}
				if err != nil {
					return err
				}

				//x.pushToStack(x.attributeBuffer, nextToken.(*parser2.CharData).Value)
				//x.pushToStack(x.attributeBuffer, identifier)
			}

		case token.TokenBlockStart:
			x.write(lt)
		case token.TokenIdentifier:
			x.write(gt, lt, slash, ct.Value, gt)
		}

	}

	/*

			return "", nil

		case parser2.TokenBlockStart:
			identifier, err := x.getFromStack(x.identifiers)
			if err != nil {
				return "", err
			}
			x.pushToStack(x.postfixes, ">")
			return ("<" + identifier), nil

		case parser2.TokenBlockEnd:
			identifier, err := x.popFromStack(x.identifiers)
			if err != nil {
				return "", err
			}
			return ("</" + identifier + ">"), nil

		case parser2.TokenGroupStart:
			return
		case parser2.TokenGroupEnd:
		case parser2.TokenGenericStart:
		case parser2.TokenGenericEnd:
		case parser2.TokenG2Preamble:
		case parser2.TokenDefineElement:
		case parser2.TokenDefineAttribute:
		case parser2.TokenAssign:
		case parser2.TokenComma:
		case parser2.TokenCharData:
		case parser2.TokenG1Comment:
		case parser2.TokenG2Comment:
			//return ("<!-- " + nextToken.value + "-->"), nil
			return "<!-- Comment -->", nil
		case parser2.TokenG1LineEnd:
		}
		fmt.Print(currentToken.TokenType())*/
	return nil
}

// getIdentifier returns the last pushed Identifier
func (x *XMLEncoder) getFromStack(stack []string) (string, error) {
	if len(stack) == 0 {
		return "", nil
	}
	return stack[len(stack)-1], nil
}

// popIdentifier returns the last pushed Identifier and removes it
func (x *XMLEncoder) popFromStack(stack []string) (string, error) {
	if len(stack) == 0 {
		return "", nil
	}
	identifier := stack[len(stack)-1]
	stack = stack[:len(stack)-2]
	return identifier, nil
}

// pushIdentifier adds an Identifier to the stack
func (x *XMLEncoder) pushToStack(stack []string, i string) {
	stack = append(stack, i)
}