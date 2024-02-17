# Cognitive Systems Design Language (CogSDL) Specification

CogSDL is a language for articulating cognitive systems so a domain engineer can model conceptual systems of ideas. It defines a minimalistic syntax that is self-describing; engineers can craft a language superset tailored to their domains. This lexical framework focuses on a system's semantics & functionality (ie cognitive types) & should make no assertions, assumptions or obligations of the underlying computational system & framework.

CogSDL is a subset of the english language that implements a deterministic lexicon, grammar & syntax. As the language is self-describing, the only implicit languuage rules are those that enable self-description.

We now generally describe the fundamental (implicit) design of CogSDL:

> A Meta Note on the following spec:
>
> - `keywords` indicate some common definition or related concept.
> - An ellipsis `...` when used in place of a word or phrase, implies any undefined or arbitrary syntax. When appended to the end of a work implies one or more instances.
> - any syntax wrapped in `[]` square brackets, is optional.

- All `grammar` is explicitly defined: words, symbols & delimiters.
  - There is one set of symbols for the entire language. (ie. letters, number, punctuation, special characters, math symbols, etc...)
  - There is one set of delimiters for the entire language. (ie. whitespace, `\0` chars, etc...)
  - Symbols & delimiters are disjoint.
  - There is one set of words for the entire language.
    - Each element in the set is a sequence of any length of any valid symbol.
    - Multiple words may be grouped together into further sets which define congruency of syntax or semantics.
- Words & Phrases have a `lexical type` that captures their sytnax, semantics & functionality as read.
  - A Phrase is an ordered sequence of `lexical word types`. Groupings of `type` sequence map to a unique identifying key, as defined by the `type` of the phrase.
  - The `lexical types` for Words & Phrases share disjoint spaces.
  - There are fundamental `lexical word types` for words:
    - noun: Identifies a particular object or thing
    - pronouns: Reference to a preceeding noun or group of nouns within the same sentence. Strict usage rules apply to avoid ambiguity.
    - adjectives: Modifies a noun or pronoun.
    - prepositions: Describes relationship between nouns in a sentence.
    - verbs: Implies action or intent exercised by a noun.
    - conjuction: connects words together (such as `and`)
  - The are fundamental `lexical phrase types` for phrases:
    - Noun Phrase: `... Noun...`
    - Prepositional Phrase: `Preposition ... Noun`
    - Verb Phrase: `Verb ...`
- A `sentence` is the smallest lexical unit of the language that is considered complete & independent.
  - A `sentence` encodes one or more thoughts.
  - A `sentence` is a composite of `clauses` which are composites of `phrases` which is a composite of `words` and/or `symbols`.
  - Puncutation delineates internal & external structural boundaries of a sentence.
  - A `sentence` is terminated by a `.`, `!` or `?` followed by a whitespace.
    > For Example...
    > This an example Sentence. This is also a sentence.
  - Word lists follow the form `word[, word,...] CONJUCTION word`
  - Independent `clause` are split on `;` punctuation.
    > For Example...
    > The cat in the hat is at bat; the cat in the hat hit the ball.
  - Dependent `clause` are split on `,` punction excluding word lists
    > For Example...
    > The cat in the hat is at bat, the kids watched him.
  - A Complex `clause` can be broken up on `:` punctuation where each line preceded by a `-` symbol is a seperate clause. The sentence may continue as normal or otherwise be terminated with a double line break
    > For Example...
    > When building a tree house:
    > - The size must be <= 10 sq m
    > - It cannot cause visual obstruction
    > - You might require a permit from your local building department
    > but never forget to have fun!
    >
    > Another Example...
    > This List has 3 items:
    > - 1
    > - 2
    > - 3
    >
    > This is a new sentence even though it was not preceded by puncutation.
- A `sentence` has a `lexical sentence type` that captures the structure, semantics & functionality of the `sentence` (but not of its encoded ideas).
  - Each `lexical sentence type` must declare:
    - Constituent order (ex. Subject-Verb-Object) as a sequence of `lexical phrase types`. Groupings of constituents are assigned an identifying key (ie. Subject in SVO)
    - Clause order as a sequence of `lexical clause types`.
    - Any special formatting or considerations associated w/ punctuation
  - The following implicit `lexical sentence types` are...
    - `Grammar Rule`: A Sentence capturing grammatical definitions: words, symbols & delimiters.
      - ex. Define grammar such that `CHARSET` is|are symbols.
      - ex. Define grammar such that `WORD`[, ...] is|are `TYPE` word [belonging to `GROUP_NAME`[, `GROUP_NAME`...]].
    - `Syntax Rules`: A Sentence describing the lexical structure & it's corresponding lexical construct & identifier.
      - ex. Define syntax such that `TODO` is a `TYPE` phrase.
