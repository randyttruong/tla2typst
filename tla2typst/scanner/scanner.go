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
	for _, t := range DELIMITERS {
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

func (s *ScannerState) ScanContent() error {

	buf, err := s.GetBuffer()

	if err != nil {
		return errors.Wrapf(err, "Something went wrong with scanning, got %v", err)
	}

	// TODO: pos is inaccurate, create a global one

	for pos, curr := range buf {
		fmt.Printf("This is the current token: %v\n", curr)
		continue

		tok := &Token{}

		if _, exists := KEYWORDS[curr]; exists {
			tok.tokenType = KEYWORD
			tok.value = curr
		} else if _, exists := OPERATORS[curr]; exists {
			tok.tokenType = OPERATOR
			tok.value = curr
		} else if fnOp, exists := containsFnOp(curr); exists {
			op_tok := &Token{}

			op_tok.tokenType = OPERATOR
			op_tok.value = fnOp
			op_tok.pos = s.pos

			s.pos++

			target := stripFnOp(curr, fnOp)

			// test if the functional operator is the entire string, which
			// it shouldn't
			if target == "" {
				continue
			}

			delims, exists := containsDelimPair(curr)

			if exists {
				found, delim1, delim2, substr := stripDelimiters(curr, delims)

				// How does this deal with potentially recursive calls?
				// Seq(Something[Something[]])
				if found {
					s_delim_token := &Token{}
					s_delim_token.tokenType = DELIMITER
					s_delim_token.value = delim1
					s_delim_token.pos = s.pos

					s.pos++
					fmt.Printf("This is just so I can build: %v, %v", delim2, substr)
				}

			}

		} else if _, exists := containsDelimPair(curr); exists {
		} else {
			// check for sneaky functional operators
			// TODO: Check for potential problems, such as if
			// the operator is a substring of a var name, although
			// I'm sure that TLA+ probably wouldn't let you naming something
			// Sequence...

			// 1. Check to see if there is a functional operator
			is_fnop := false

			for _, op := range FUNCTIONAL_OPERATORS {
				if strings.Contains(curr, op) {
					tok.tokenType = OPERATOR
					tok.value = curr
					is_fnop = true
					break
				}
			}

			is_seq := false
			start, end := curr[0], curr[len(curr)-1]

			if start == '[' || end == ']' || start == '{' || end == '}' {
				var symbol string

				op_tok := &Token{}
				op_tok.tokenType = OPERATOR

				// Check for singleton operator (representing some comprehension)
				if len(curr) == 1 {
					op_tok.value = curr
					op_tok.pos = pos
					pos++
					s.stream = append(s.stream, op_tok)
					continue
				}

				// If not singleton, then check to see
				if containsStartingSymbol(curr) {
					symbol = string(curr[0])
					curr = curr[1:]

					op_tok.value = symbol
					op_tok.pos = pos
					pos++
					s.stream = append(s.stream, op_tok)
				} else if containsEndingSymbol(curr) {
					symbol = string(curr[len(curr)-1])
					curr = curr[0 : len(curr)-1]

				}

				pos++

				// check if we're subscripting something
				if containsBothBrackets(curr) {
					start, end := findStartAndEndBrackets(curr)
					fmt.Println("This is just so that I can build %v, %v", start, end)
				}

				// now check if it's a string literal, a numeric literal, or
				// an identifier

				// TODO: Check what happens when we try to index something, ie arr[1]
				// TODO(future): oooo we also have set notation to deal with too
				// TODO(future): oooo we also have set indexing to deal with too
			}

			if !is_fnop && !is_seq {

				// check if numeric
				if _, err := strconv.Atoi(curr); err == nil {
					tok.tokenType = NUM_LITERAL
					tok.value = curr
				} else if start == '"' && end == '"' {
					tok.tokenType = STRING_LITERAL
					tok.value = curr
				}
			}
		}

		tok.pos = s.pos
		s.stream = append(s.stream, tok)
		s.pos++
	}

	return nil
}
