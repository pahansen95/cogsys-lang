import re

LEXICON_SYMBOLS = {
  'Alphabet': { 'kind': 'regex', 'regex': r'A-Za-z', },
  'Numeric': { 'kind': 'regex', 'regex': r'0-9', },
  'Whitespace': { 'kind': 'set', 'set': ( ' ', '\t', '\v', '\r', '\n') },
  'Punctuation': { 'kind': 'set', 'set': ( ':', ';', ',', '.', '?', '!' ) },
  'Seperator': { 'kind': 'set', 'set': ( "'", '"', '`', '/', '\\', '|', '~', '-', '_', '{', '}', '<', '>', '(', ')', '[', ']' ) },
  'Symbolic': { 'kind': 'set', 'set': ( '@', '#', '$', '%', '^', '&', '*', '+', '=' ) }
}
"""The legal set of symbols in the Language"""

LEXICON_RULESETS = {
  'WhiteSpace': { 'kind': 'function', 'fn': lambda token: ... }, # An combination of Whitespace Symbols
  'DecimalNumber': { 'kind': 'function', 'fn': lambda token: ... }, # A (Base 10) Number: ex. 1 | 100 | 1,000 | 0.25 | 1,234.56789 | 1/3 | etc...
  'Word': { 'kind': 'default' }, # The Default Fallthrough Rule
}
"""Set of Rules used to group together symbols into words

TODO

I'm not sure if Symbol rulesets should be exist here
or if they should be a FSM that evaluates IR Tokens.
I guess it depends on how we parse & evaluate the
source text.
"""

LEXICON_DICTIONARY = {
  'TODO': 'A task to be completed in the future',
  # TODO: Can we generate a Dictionary from some established source (ie. The Oxford English Dictionary)?
}
"""Known/Reserved words of the language & their implied semantics"""

SYNTAX_FORMAT = {
  'Document': { ... },
  'Section': { ... },
  'Paragraph': { ... },
  'Sentence': { ... },
  'Clause': {
    'Independent': { ... },
    'Dependent': { ... }
  },
}
"""Rules detailing grouping of Lexical Tokens based on some structured pattern"""

SYNTAX_PROSE = {
  'SVO': { ... }, # Subject-Verb-Object
}
"""Rules detailing accepted sequences of Lexical tokens"""