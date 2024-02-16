/*markdown
# Lexical Analysis of CogSL (Cognitive Systems Language)

Consists of:

- Grammar: Declaring Syntax Rules
- Tokenization: Read in as stream of text, identify lexems & produce a stream of tokens
- Parsing: Read in the stream of tokens, assembling an AST

The design shall focus on streaming text, tokens & nodes.

The goal of CogSL is to express ideas & intent in a language whose form & function sits between natural language & programming languages.
CogSL should have a form closely related to natural language while providing functionality to rigidly structure language semantics.
To accomplish this, CogSL is self-defining:

- The lanuage defines a minimal set of grammar & syntax that it can use to describe itself.
- Users or Maintainers can extend or create supersets of the language by defining custom grammar & syntax.
- The defined grammar & syntax is extracted from the source for further semantic analysis.

## Grammar & Syntax

The reference implementation of CogSL uses a subset of English (herein lang) with emphasis on simple & straightforward prose that
follows well defined patterns. All grammar in lang must be defined for semantic meaning to be applied. Any grammar in the source is effectively ignored during analysis.

To be self describing, we first define a core set of language grammar & syntax:

- Sentence: A complete thought.
- Kinds of Sentences:
	- Declaratives: States a thought.
	- Imperatives: Authorative statement, commands or actions.
	- Conditionals: Expressing a particular set of circumstances
	- Interrogatives: Asks questions to clarify, cofirm or otherwise gather more information.
	- Obligations: Asserts a requirement or necessity.
	- Promises: statements of commitment to be eventually fullfilled
- Grammatical Structure of a Declarative Sentence:
	```
	SVO Sentence
	│
	├── Noun Phrase
	│		├── Noun
	├── Predicate Phrase
			├── Verb
			├── Object (Optional)
	```
  - Active Voice is the grammatical construction of choice.
		- Active Voice implements Subject-Verb-Object (SVO). The subject of the sentence performs the action expressed by the verb. The object recieves the action.
			- ex. She kicks the ball: `She` is the Subject, `kicks` is the verb, `the ball` is the object
	-
- Phrases, Clauses & Sentences



*/

package cogsl

import (
	"fmt"
	"os"
	"io"
	"unicode"
)

type TokenKind uint

const (
	TK_UNDEFINED TokenKind = iota // An undefined token
	TK_START_OF_DOC 						  // The start of the document
	TK_END_OF_DOC							    // The end of the document
	TK_WHITESPACE 							  // A collection of whitespace characters excluding newline unless it is escaped
	TK_LINE_BREAK 							  // An unescaped newline character
	TK_WORD 											// A collection of non-whitespace characters
)

type Token struct {
	Start int // The Starting index of the token in the buffer; ie. buffer[Start:Stop]
	Stop int // The Stopping index of the token in the buffer; ie. buffer[Start:Stop]
	Kind  TokenKind
}

type Tokenizer struct {
	TextSource io.Reader  // The Source of the text to analyze
	Buffer      []byte     // The buffer to write the source code into
	TokenSink  chan Token // The Channel on which to emit Lexical Tokens
}

func (t *Tokenizer) Tokenize(chunk_size int, growth_factor int, max_size int) error {
	/*
		Read the TextSource into the buffer. Assumes TextSource is a stream of UTF-8 characters from a single document.
		As characters are read, emit tokens accross the channel

		- chunk_size: The size of the chunk to read from the TextSource
		- growth_factor: How many chunk_sizes to grow the buffer when necessary

		# TODO

		- Use channels to control the tokenization process:
			- The Buffer becomes full
		- Dispatch the exact tokenization rules based on the state (FSM)
	*/
	growth_factor = max(growth_factor, 1)
	growth_size := chunk_size * growth_factor
	if t.Buffer == nil { return fmt.Errorf("the tokenizer's buffer was not initialized") }
	kind_log := make([]TokenKind, chunk_size)
	kind_log[0] = TK_START_OF_DOC
	t.TokenSink <- Token{0, 0, kind_log[0]}
	buf_iter_idx := 0 // where in the buffer are we currently
	escape_seq := false // Are we in an escape sequence?
	for {
		// Read a chunk of characters into the buffer
		chunk := make([]byte, chunk_size)
		n, err := t.TextSource.Read(chunk)
		fmt.Fprintf(os.Stderr, "Read %d bytes\n", n)
		if n == 0 && err == nil { continue } // TODO: Better handle this case; usually encountered in Non-Blocking IO
		if n > 0 && cap(t.Buffer) < len(t.Buffer) + n { // Grow the buffer if necessary
			extension_size := min(max_size - len(t.Buffer), growth_size)
			if extension_size == 0 { return fmt.Errorf("the tokenizer's buffer has reached the maximum size of %d", max_size) }
			t.Buffer = append(t.Buffer, make([]byte, 0, extension_size)...)
		}
		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "Buffer Iter Idx: %d\nBuffer Length: %d\n", buf_iter_idx, len(t.Buffer))
				// If there is text left in the buffer emit it as a token of the last kind & then emit the end of document token
				if buf_iter_idx < len(t.Buffer) - 1 { t.TokenSink <- Token{buf_iter_idx, len(t.Buffer) - buf_iter_idx, kind_log[len(kind_log)-1]} }
				t.TokenSink <- Token{buf_iter_idx, buf_iter_idx, TK_END_OF_DOC}
				break
			} else {
				close(t.TokenSink); return err
			}
		}
		chunk_start_idx := len(t.Buffer) // Where the Chunk starts in the buffer; just the end of the buffer before we append the chunk
		t.Buffer = append(t.Buffer, chunk[:n]...)
		fmt.Fprintf(os.Stderr, "Chunk Start Idx: %d\n", len(t.Buffer))

		// Parse the chunk into tokens
		var rune_kind TokenKind
		chunk_iter_idx := 0 // where in the chunk are we currently
		for i := 0; i < n; i++ {
			cur_rune := rune(t.Buffer[i])
			switch {
			// Regular Cases
			case !escape_seq && cur_rune == '\\': escape_seq = true // A New Escape Sequence
			case !escape_seq && unicode.IsSpace(cur_rune): if cur_rune == '\n' { rune_kind = TK_LINE_BREAK } else { rune_kind = TK_WHITESPACE } // Whitespace
			case !escape_seq: rune_kind = TK_WORD // If it's not an escape sequence or whitespace, then it's a word
			// Escape Sequence Cases
			case escape_seq && cur_rune == '\n': escape_seq = false; rune_kind = TK_WORD // A newline escape sequence
			case escape_seq && cur_rune == '\\': escape_seq = false; rune_kind = TK_WORD // Escaping the escape character
			// There are no other valid escape sequences so emit an undefined token
			default: escape_seq = false; rune_kind = TK_UNDEFINED
			}
			if escape_seq { continue; } // If we are in an escape sequence then we short circuit & jump to the next rune
			// If the kind of rune is different from the previous, then evaluate if we need to emit a token
			last_kind := kind_log[len(kind_log) - 1]
			if rune_kind != last_kind {
				// Edge Triggered: Emit the token, append the new kind to the log & update our indices
				chunk_iter_idx = i
				token_start_idx := buf_iter_idx // The Token starts at the current buffer index
				token_stop_idx := chunk_start_idx + chunk_iter_idx // The Token stops at the position in the buffer relative to the start of the chunk
				t.TokenSink <- Token{token_start_idx, token_stop_idx, last_kind}
				kind_log = append(kind_log, rune_kind)
				buf_iter_idx = token_stop_idx // The New Buffer index is where the token stopped
			}
			// otherwise we just continue
		}
	}

	// Cleanup our end of the channel
	close(t.TokenSink)
	return nil
}
