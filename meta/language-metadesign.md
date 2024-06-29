# Language MetaDesign

Designing & Implementation of a language is broken down into the following hierarchy of ideas/concepts

```plaintext
Language
├── Coding
│   └── Tokenization
└── Linguistics
    ├── Lexicon
    │   ├── Symbols
    │   │   ├── Alphabet
    │   │   ├── Numerics
    │   │   ├── Punctuation
    │   │   └── Whitespace
    │   └── Semantics
    │       ├── Words
    │       └── Phrases
    ├── Syntax
    ├── Prose
    └── Formatting
        └── Document
            └── Section
                ├── Section
                │   └── ... (recursive)
                └── Paragraph
                    └── Sentence
                        └── Clause
                            ├── Independent
                            └── Dependent
```

- **Coding**: Dealing with the symbolic representation, transmission & storage of the language
- **Linguistics**: Dealing with the structure & intrinsics of the language
  - **Lexicon**: Legal words, phrases & symbols of the Language & the intrinsic information defining them
    - **Symbols**: Dealing with the characters representing (traditionally phonetic sounds)
      - TODO: How is this better defined.
    - **Semantics**: The well defined set of ordered sequences of symbols each encapsulating information (usually one of a lower dimension)
      - TODO: How is this better defined
  - **Syntax**: Defining legal arrangements of the Lexicon; (ex. SVO Sentence Structure)
  - **Prose**: Dealing with the application of Lexicon & Syntax to convey higher dimensional information (ie. How the language is written to convey big ideas)
  - **Formatting**: Dealing with the holistic structure of the language
    - **Document**: An ordered sequence of Sections.
    - **Section**: An ordered sequence of Paragraphs or (nested) sections
    - **Paragraph**: An ordered sequence of Sentences.
    - **Sentence**: An ordered sequence of clauses.
    - **Clause**: The smallest grouping of Semantics that encapsulate "knowledge"
      - **Dependent**: A Clause that references another Clause when expressing "knowledge".
      - **Independent**: A Clause that wholly encapsulates "knowledge", without referencing any other clauses.
