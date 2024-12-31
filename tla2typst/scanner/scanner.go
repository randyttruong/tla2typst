package scanner

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// ScannerState represents the scanner's current state.
type ScannerState struct {
	loader    *Loader
	stream    []*Token
	val       string
	line      int
	col       int
	tokenType TokenType
	pos       int // FIXME: DEPRECATE THIS IN FAVOR OF ScannerState.line and ScvannerState.col
}

var (
	Scanner = &ScannerState{pos: 0, tokenType: UNASSIGNED}
)

// InitScanner passes the loader into ScannerState
func InitScanner(loader *Loader) error {
	if loader == nil || GetLoader().buf == "" {
		return errors.New("Loader or buffer does not exist/is empty.")
	}

	Scanner.loader = loader

	return nil
}

func GetScanner() *ScannerState {
	return Scanner
}

func (s *ScannerState) GetBuffer() (*string, *int, error) {
	if Scanner.loader == nil || Scanner.loader.buf == "" {
		err := errors.New("Unable to get buffer, loader or bytes array does not exist. Exiting.")

		return nil, nil, err
	}

	return &(Scanner.loader.buf), &(Scanner.loader.bufLen), nil
}

func containsBothBrackets(s string) bool {
	return strings.Contains(s, "[") && strings.Contains(s, "]")
}

func findStartAndEndBrackets(s string) (int, int) {
	return strings.LastIndex(s, "["), strings.LastIndex(s, "]")
}

func containsStartingSymbol(s string) bool {
	return s[0] == '[' || s[0] == '{'
}

func containsEndingSymbol(s string) bool {
	end := len(s) - 1

	return s[end] == ']' || s[end] == '}'
}

func stripFnOp(s string, fnOp string) string {
	idx := strings.LastIndex(s, fnOp)

	t := s[idx+1:] // delimiter is dedicated to its own token

	return t
}

func stripDelimiters(s string, delims string) (bool, string, string, string) {

	s_idx, e_idx := -1, -1
	start, end := "", ""

	switch delims {
	case "()":
		s_idx, e_idx = strings.LastIndex(s, "("), strings.LastIndex(s, ")")
		start = "("
		end = ")"
	case "[]":
		s_idx, e_idx = strings.LastIndex(s, "["), strings.LastIndex(s, "]")
		start = "("
		end = ")"
	case "{}":
		s_idx, e_idx = strings.LastIndex(s, "{"), strings.LastIndex(s, "}")
		start = "("
		end = ")"
	}

	if s_idx == -1 || e_idx == -1 {
		return false, "", "", s
	}

	return true, start, end, s[s_idx+1 : e_idx]
}

func containsDelimPair(s string) (string, bool) {
	for t, _ := range DELIMITERS {
		first, second := string(t[0]), string(t[1])

		if strings.Contains(s, first) && strings.Contains(s, second) {
			return t, true
		}
	}
	return "", false
}

func isNumLiteral(s string) bool {
	_, err := strconv.Atoi(s)

	if err != nil {
		return false
	}

	return true
}

func (s *ScannerState) pushToken() *Token {
	s.val = strings.TrimSpace(s.val)
	new_tok := &Token{tokenType: s.tokenType, value: s.val}

	s.tokenType = UNASSIGNED
	s.stream = append(s.stream, new_tok)
	s.val = ""
	s.pos++
	return &Token{tokenType: UNASSIGNED}
}

func (s *ScannerState) ScanContent() error {

	buf, bufLen, err := s.GetBuffer()

	fmt.Printf("This is the len: %v\n", *bufLen)

	if err != nil {
		return errors.Wrapf(err, "Something went wrong with scanning, got %v", err)
	}

	for pos, curr := range *buf {

		currStr := string(curr)

		s.val += currStr

		// fmt.Println("---")
		// fmt.Printf("[DEBUG] This is the current char: %v\n", currStr)
		// fmt.Printf("[DEBUG] This is the current token: %v\n", s.val)
		// fmt.Printf("[DEBUG] This is the pos: %v\n", pos)

		// HACK: This is literally just for all of the generic ones,
		// definitely make this a lot more finegrained in the future
		// FIXME: Add support for all of the new tokens
		if s.tokenType == UNASSIGNED {
			if _, exists := DELIMITERS[s.val]; exists {
				s.tokenType = DELIMITER
				s.pushToken()
			} else if _, exists := SPECIALS[s.val]; exists {
				s.tokenType = SPECIAL
			} else if _, exists := OPERATORS[s.val]; exists {
				s.tokenType = OPERATOR
				if curr == '=' {
					continue
				}
				s.pushToken()
			} else if _, exists := KEYWORDS[s.val]; exists {
				s.tokenType = KEYWORD
			} else if curr == '"' {
				// NOTE: Because there aren't any while-loops in Go,
				// it is necessary to abstract the token logic
				s.tokenType = STRING_LITERAL
			} else if unicode.IsNumber(curr) {
				s.tokenType = NUM_LITERAL
			} else if unicode.IsLetter(curr) {
				s.tokenType = IDENTIFIER
			} else if unicode.IsSpace(curr) {
				s.val = ""
			}
		} else if s.tokenType == KEYWORD {
			if unicode.IsSpace(curr) {
				s.val = s.val[0 : len(s.val)-1]
				s.pushToken()
			} else if unicode.IsLetter(curr) || unicode.IsNumber(curr) {
				s.tokenType = IDENTIFIER
			} else {
				err = fmt.Errorf("[FATAL ERROR]: Attempted adding a nonalpha character to a keyword/identifier")
				return err
			}
		} else if s.tokenType == IDENTIFIER {
			// cases:
			// + cannot have delimiters or punctuation
			//  FIXME: What if an identifier turns out to be a keyword?
			if curr == ' ' || curr == '\n' {
				s.val = s.val[0 : len(s.val)-1]

				if KEYWORDS[s.val] {
					s.tokenType = KEYWORD
				} else if OPERATORS[s.val] {
					s.tokenType = OPERATOR
				}

				s.pushToken()
			} else if curr == '(' || curr == ')' || curr == '{' || curr == '}' {
				s.val = s.val[0 : len(s.val)-1]
				s.pushToken()

				s.tokenType = DELIMITER
				s.val += currStr
				s.pushToken()
			}
		} else if s.tokenType == OPERATOR {
			if curr == ' ' {
				s.pushToken()
			}
		} else if s.tokenType == STRING_LITERAL {
			if curr == '"' || curr == '\n' {
				s.pushToken()
			} else if pos == *bufLen-1 {
				err = fmt.Errorf("[FATAL ERROR]: Unclosed literal")
				return err
			}
		} else if s.tokenType == NUM_LITERAL {
			// error handling here
			if !unicode.IsNumber(curr) {
				err = fmt.Errorf("[FATAL ERROR]: Numeric literal found at (LOL I NEED TO KEEP TRACK OF POS) inserts non-numeric char: %v", s.val)
				return err
			}

			if curr == ' ' {
				s.val = s.val[0 : len(s.val)-1]
				s.pushToken()
			}
		} else if s.tokenType == SPECIAL {
			switch s.val[0] {
			case '(':
				if curr == '*' {
					s.tokenType = BLOCK_COMMENT
				} else {
					body := s.val[1:]
					s.val = string(s.val[0])

					s.tokenType = DELIMITER
					s.pushToken()

					s.val = body

					if unicode.IsLetter(curr) {
						s.tokenType = IDENTIFIER
					} else if curr == '"' {
						s.tokenType = STRING_LITERAL
					} else if unicode.IsNumber(curr) {
						s.tokenType = NUM_LITERAL
					} else if _, exists := DELIMITERS[string(curr)]; exists {
						s.tokenType = DELIMITER
						s.pushToken()
					}
				}
			case '\\':
				if curr == '*' {
					s.tokenType = INLINE_COMMENT
				} else if curr == '/' {
					s.tokenType = OPERATOR
				} else if unicode.IsLetter(curr) {
					s.tokenType = UNASSIGNED
				} else if unicode.IsNumber(curr) {
					err := fmt.Errorf("[FATAL ERROR]: Unknown token %v, breaking", s.val)
					return err
				}
			case '/': // HACK: This is probably not a very robust way of handling '/'-prefixed stuff-- but it works for now
				if curr == ' ' || curr == '\\' || curr == '=' {
					s.tokenType = OPERATOR
				}
			}
		} else if s.tokenType == INLINE_COMMENT {
			if curr == '\n' || pos == *bufLen-1 {
				s.pushToken()
			}
		} else if s.tokenType == BLOCK_COMMENT {
			if s.val[len(s.val)-2:len(s.val)] == "*)" {
				s.pushToken()
			}
		}
	}

	s.tokenType = EOF
	s.pushToken()

	return nil
}
