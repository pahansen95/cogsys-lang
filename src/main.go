package main

import (
	"github.com/pahansen95/cogsyslang/cogsl"
	"fmt"
	"os"
)

func main() {
	tokenizer := cogsl.Tokenizer{
		TextSource: os.Stdin, // TODO: Read directly from the file
		Buffer: make([]byte, 0, 1024 * 1024 * 10), // 10MB File Buffer
		TokenSink: make(chan cogsl.Token), // TODO: Make this a buffered channel
	}

	// Spawn the Tokenizer; wrap it in a function to handle errors
	go func () {
		err := tokenizer.Tokenize(
			1024 * 4, // 4K Chunks
			4, // Grow by 16K (4 Chunks)
			1024 * 1024 * 100, // 100MB Max Size
		)
		if err != nil { fmt.Fprintf(os.Stderr, "Fatal: %s\n", err) }
		if _, ok := <-tokenizer.TokenSink; ok { close(tokenizer.TokenSink) }
	}()

	// Read the Tokens & log them to stderr
	for token := range tokenizer.TokenSink {
		var _kind string
		switch token.Kind {
		case cogsl.TK_START_OF_DOC: _kind = "START_OF_DOC"
		case cogsl.TK_END_OF_DOC: _kind = "END_OF_DOC"
		case cogsl.TK_WORD: _kind = "WORD"
		case cogsl.TK_WHITESPACE: _kind = "WHITESPACE"
		case cogsl.TK_LINE_BREAK: _kind = "LINE_BREAK"
		case cogsl.TK_UNDEFINED: _kind = "UNDEFINED"
		default: _kind = fmt.Sprintf("UNKNOWN<%d>", token.Kind)
		}
		msg := fmt.Sprintf("%s<%d> @ [%d:%d]\n", _kind, token.Kind, token.Start, token.Stop)
		if token.Kind == cogsl.TK_WORD { msg = fmt.Sprintf("%s%s\n", msg, tokenizer.Buffer[token.Start:token.Stop]) } 
		_, err := fmt.Fprintf(os.Stderr, msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal: %s\n", err)
			close(tokenizer.TokenSink)
			break
		}
	}

	// Close the TokenSink if it's still open
	if _, ok := <-tokenizer.TokenSink; ok { close(tokenizer.TokenSink) }
}