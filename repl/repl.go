package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/nelusnegur/monkeylang/token"
	"github.com/nelusnegur/monkeylang/lexer"
)

const PROPMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROPMT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}
		
		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
