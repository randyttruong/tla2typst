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
	loader *Loader
	stream []*Token
	val    string
	pos    int
}

var (
	Scanner = &ScannerState{pos: 0}
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

func (s *ScannerState) pushToken(tok *Token) *Token {
	tok.pos = s.pos
	tok.value = strings.TrimSpace(s.val)

	s.stream = append(s.stream, tok)
	s.val = ""
	s.pos++
	return &Token{tokenType: UNASSIGNED}
}

func (s *ScannerState) ScanContent() error {

	buf, bufLen, err := s.GetBuffer()

	fmt.Printf("This is the len: %v", bufLen)

	if err != nil {
		return errors.Wrapf(err, "Something went wrong with scanning, got %v", err)
	}

	tok := &Token{tokenType: UNASSIGNED}

	for pos, curr := range *buf {

		currStr := string(curr)

		s.val += currStr

		// fmt.Println("---")
		// fmt.Printf("[DEBUG] This is the current char: %v\n", currStr)
		// fmt.Printf("[DEBUG] This is the current token: %v\n", s.val)
		// fmt.Printf("[DEBUG] This is the pos: %v\n", pos)

		// TODO: Add comment support
		if tok.tokenType == UNASSIGNED {
			// TODO: Check out how newlines and whitespace are addressed
			// Should I keep track of the position?
			if _, exists := DELIMITERS[s.val]; exists {
				tok.tokenType = DELIMITER
				tok = s.pushToken(tok)
			} else if _, exists := OPERATORS[s.val]; exists {
				tok.tokenType = OPERATOR
				if curr == '=' {
					continue
				}
				tok = s.pushToken(tok)
			} else if _, exists := KEYWORDS[s.val]; exists {
				tok.tokenType = KEYWORD
			} else if curr == '"' {
				// NOTE: Because there aren't any while-loops in Go,
				// it is necessary to abstract the token logic
				tok.tokenType = STRING_LITERAL
			} else if unicode.IsNumber(curr) {
				tok.tokenType = NUM_LITERAL
			} else if unicode.IsLetter(curr) {
				tok.tokenType = IDENTIFIER
			} else if unicode.IsSpace(curr) {
				s.val = ""
			}
		} else if tok.tokenType == KEYWORD {
			if unicode.IsSpace(curr) {
				s.val = s.val[0 : len(s.val)-1]
				tok = s.pushToken(tok)
			} else if unicode.IsLetter(curr) || unicode.IsNumber(curr) {
				tok.tokenType = IDENTIFIER
			} else {
				err = fmt.Errorf("[FATAL ERROR]: Attempted adding a nonalpha character to a keyword/identifier")
				return err
			}
		} else if tok.tokenType == IDENTIFIER {
			// cases:
			// + cannot have delimiters or punctuation
			if curr == ' ' || curr == '\n' {
				s.val = s.val[0 : len(s.val)-1]
				tok = s.pushToken(tok)
			} else if curr == '(' || curr == ')' || curr == '{' || curr == '}' {
				s.val = s.val[0 : len(s.val)-1]
				tok = s.pushToken(tok)

				tok.tokenType = DELIMITER
				s.val += currStr
				tok = s.pushToken(tok)
			}
		} else if tok.tokenType == OPERATOR {
			if curr == ' ' {
				tok = s.pushToken(tok)
			}
		} else if tok.tokenType == STRING_LITERAL {
			if curr == '"' || curr == '\n' {
				tok = s.pushToken(tok)
			} else if pos == *bufLen-1 {
				err = fmt.Errorf("[FATAL ERROR]: Unclosed literal")
				return err
			}
		} else if tok.tokenType == NUM_LITERAL {
			// error handling here
			if !unicode.IsNumber(curr) {
				err = fmt.Errorf("[FATAL ERROR]: Numeric literal found at (LOL I NEED TO KEEP TRACK OF POS) inserts non-numeric char: %v", s.val)
				return err
			}

			if curr == ' ' {
				s.val = s.val[0 : len(s.val)-1]
				tok = s.pushToken(tok)
			}
		}
	}

	return nil
}
