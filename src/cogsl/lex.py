"""

Lexical Analyzer for the Cognitive Systems Language

"""
from __future__ import annotations
from typing import TextIO, Any, TypedDict, NotRequired, Literal, Generic, TypeVar
from collections.abc import Iterator, Iterable

import io
from . import LEXICON_SYMBOLS

token_kind_t = Literal['Word']

class DocumentToken(TypedDict):
  buffer: str
  """The Token itself"""
  pos: int
  """Where in the document the Token starts"""

class LexicalToken(TypedDict):
  """The Minimal set of attributes necessary for a lexical Token"""
  buffer: str
  """The Token itself"""
  kind: token_kind_t
  """The Classification of Token"""

  class Word(TypedDict):
    """The attributes expected of a Word"""
    kind: Literal['Word']
    """The Classification of Token"""
    reserved: bool
    """Denotes if this word is defined in the Lexicon"""

ID = TypeVar('ID')
NODE = TypeVar('NODE')
class Tree(Generic[ID, NODE], TypedDict):
  root_node: ID
  nodes: dict[ID, NODE]
  edges: dict[ID, set[ID]]

token_seq_t = Iterator[DocumentToken]
lexical_seq_t = Iterator[LexicalToken]
ast_t = Tree[int, Any]

CHUNK_SIZE = 16384

class Scanner:

  @staticmethod
  def _is_legal(symbol: str) -> bool:
    """Evaluates if the Symbol is legal"""
    raise NotImplementedError

  @staticmethod
  def scan(source: TextIO, chunk_size: int = CHUNK_SIZE) -> TextIO:
    """Reads a Text Source checking for illegal symbols"""
    buffer = io.StringIO()
    buf_pos = 0
    eof = False
    while not eof:
      chunk = source.read(chunk_size)
      eof = chunk < chunk_size
      for symbol in chunk:
        illegal = not Scanner._is_legal(symbol)
        if illegal: raise ValueError(f"Illegal Symbol `{symbol}` at Position `{buf_pos}` in the Stream")
        buffer.write(symbol)
        buf_pos += 1
    return buffer

class Tokenizer:

  class Document:

    @staticmethod
    def _split_on_symbols(buffer: str, symbols: set[str], max_split: int = None) -> list[tuple[str, str | None]]:
      """Splits the Buffer into items after each occurance of any one of the symbols.

      If `max_split` times is set then stop splitting after N matches

      returns the list of items as a tuple of (buffer, symbol)
      """
      raise NotImplementedError

    @staticmethod
    def tokenize(document: TextIO, chunk_size: int = CHUNK_SIZE) -> token_seq_t:
      """Convert a document into a Token Sequence IR"""
      stash = io.StringIO()
      tokens: list[DocumentToken] = []
      eof = False
      while not eof:
        chunk = document.read(chunk_size)
        eof = chunk < chunk_size
        buffer = stash.read() + chunk # Prepend the previous stash

        split_order = (
          (set(LEXICON_SYMBOLS['Whitespace']), None),
          (set(LEXICON_SYMBOLS['Punctuation']), None),
        )
        items = [ { 'buffer': buffer } ] # Start with a single item
        for _split_eval_idx, (symbols, max_split) in enumerate(split_order):
          """
          
          Here we take the buffer & split it on the symbols declared.
            The last most item is stashed for later evaluation if we haven't reached EoF since we don't know if the token extends into the next chunk
          We then further split each item on the next set of symbols repeating until all splitting symbols are evaluated.

          TODO: Add Document Positions & Other MetaData
          
          """
          _items = []
          for item in items:
            # For Each Item; split the item's buffer on the set of symbols
            if max_split is not None: subitems = Tokenizer._split_on_symbols(item['buffer'], symbols, max_split=max_split)
            else: subitems = Tokenizer._split_on_symbols(item['buffer'], symbols)
            
            # Stash the last Item if we haven't encountered EoF
            if not eof and _split_eval_idx == 0:
              _buf, _sep = subitems[-1]
              stash.write(_buf)
              if _sep is not None: stash.write(_sep)
              subitems = subitems[:-1]

            # Flatten the List
            for _buf, _sep in subitems:
              _items.append({ 'buffer': _buf })
              if _sep is not None: _items.append( { 'buffer': _sep } )

          # Swap out the current set of items to evaluate
          items = _items
      
      assert len(stash.getvalue()) == 0, "Stash has leftovers when it shouldn't"

      # TODO: Fill-in remaining Token Metadata

      return iter(tokens)

  class Lexicon:

    @staticmethod
    def _eval_ruleset(token_seq: Iterable[DocumentToken]) -> tuple[bool, LexicalToken | None]:
      """Attempts to evaluate the sequence of tokens into a single Lexical token.
      
      If evaluation was succesful then returns (True, LexicalToken)
      otherwise returns (False, None)
      """
      raise NotImplementedError

    @staticmethod
    def _lookup_word(word: LexicalToken) -> tuple[bool, LexicalToken.Word | None]:
      """Looks up the Word in the Lexicon's Dictionary"""
      assert word['kind'] == 'Word'
      raise NotImplementedError

    ### TODO: Probably need to revisit this once AST Assembly is implemented
    #
    #
    #
    @staticmethod
    def tokenize(token_seq: token_seq_t) -> lexical_seq_t:
      """Tokenize the DocumentToken Sequence IR against the Lexicon to generate a LexicalToken Sequence"""

      tokens: list[LexicalToken] = []
      tkn_stash: list[DocumentToken] = []
      for doc_tkn in token_seq:
        tkn_stash.append(doc_tkn)
        ok, lex_tkn = Tokenizer.Lexicon._eval_ruleset(tkn_stash)
        if not ok: continue
        tkn_stash.clear()
        tokens.append(lex_tkn)
        if lex_tkn['kind'] == 'Word':
          ok, word = Tokenizer.Lexicon._lookup_word(lex_tkn)
          if ok: tokens[-1] |= word
      
      return iter(tokens)
    #
    #
    #
    ###

class Parser:

  @staticmethod
  def _eval_format(seq: Iterable[LexicalToken]) -> tuple[bool, Any | None]:
    """Evaluates the Sequence of Lexical Tokens against the Formatting Rules & returns the largest possible formatting group"""
    raise NotImplementedError
  
  @staticmethod
  def _eval_prose(seq: Iterable[LexicalToken]) -> tuple[bool, Any | None]:
    """Evaluates the sequence of Lexical Tokens against the Prose Rules & returns the largest possible prose group"""
    raise NotImplementedError
  
  @staticmethod
  def parse(lex_seq: lexical_seq_t) -> ast_t:
    """Produce an Abstract Syntax Tree from the provided Lexical Sequence."""
    
    """NOTE: Proposed Implementation
    
    1. First, we evaluate the sequence against formatting rules grouping the tokens matching a specific rule.
      - Recursively Evaluate each grouping's sub-sequence against the formatting rules until no rules match.
    2. Next, for each formatting group, evaluate the lexical (sub)sequence against the prose rules.
      - Recursively Evaluate each grouping's sub-sequence against the prose rules un until no rules match.
      - At this point each Prose Group is considered "terminal syntax" & no further evaluation can occur.
    3. Assemble the AST walking the recursive evaluation hierarchy.

    """

    raise NotImplementedError