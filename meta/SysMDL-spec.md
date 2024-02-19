# Systems Modeling Domain Language (SysMDL) Specification

SysMDL is a language that enables domain engineers to comprehend & articulate a system's semantics & functionality. SysMDL is a subset of the english language that implements a deterministic lexicon, grammar & syntax which allows for precise & expressive description of domain concepts.

First we describe the core design principles of SysMDL:

- SysMDL is based on the English language. It is intended to be read & written like Natural Language (NL) while being fully deterministic to allow for lossless parsing.
  - Like English there are different classifications of words. In English, such classifications are nouns & verbs. Sequences of these word classifications constitute a phrase which is further classified; noun phrases & verb phrases. These phrase consitute clauses which constitute sentences. Furthermore, sentence structure is key to conveying semantic information of the sentence's content; Subject-Verb-Object (SVO) is one common structure.
  
  - However, to remain deterministic, SysMDL does not rely on the implicit syntax of English. Instead it defines itself in terms of similar syntactic structures:
    - SysMDL Introduces `lexical typing` to assign classification to `word types`, `phrase types`, `clause types` & `sentence types`. Each lexical type assigns semantics to grammar & syntax structure. These types aim to be explicit in their application as opposed to generic; this helps in reducing ambiguity in the language spec. Below is the hierarchical structure between types. ie. Sentences are composed of clauses, etc...

      ```SysMDL
      Sentence
      └── Clause
          └── Phrase
              └── Word
      ```

    - Each `sentence type` encodes semantics for a certain capability or feature of the language. For example, an `existence type` would encode the semantics & structure for describing the existence of an object:

      ```SysMDL
      # Defined as
      `subject-kind-verb`: `seq[ref-attr-val]`
      # Example
      Peter, a Person, exists:
        - his age is 29.
        - his domain is IT.
        - his goal is to seek the Holy Grail.          
      ```

    - Each `clause type` encodes syntactic structure as a sequence of phrases. In the `sentence type` example above, the First line would be a `subject-kind-verb type` that is structured into the following `phrase types`:

      ```SysMDL
      [SUBJECT], [KIND], [VERB]
      ```

      Where `Peter` maps to `SUBJECT`, `a Person` maps to `KIND` & `exists` maps to `VERB`. (Note; `kind` is a synonym for `type` in this context; it's used to prevent type names like `subject-type-verb type`).

    - Each `phrase type` encodes the syntactic structure as a sequence of words. The `clause type` example above breaks down accordingly:

      ```SysMDL
      # Subject Phrase
      ENTITY_NAME+
      # Kind Phrase
      ...? KIND{1}
      # Verb Phrase
      STATE{1}
      ```

      Where `Peter` maps to `ENTITY_NAME+`, `Person` maps to `KIND{1}` & `exists` maps to `STATE{1}`. Each of the words that map to these keys must be a word that meet's that kind's expectations (more about this below). The modifiers on the end of the keys denote quantity or conditionals. Note the ellipsis (`...?`) just indicates optional "filler" words such as articles; with an emphasis on explicitness, such usage is minimized.

    - Each `word type` encodes a set of valid grammatical symbols. This set can be represented as a simple set, a regex or a combination of the two. The `phrase type` example above breaks down thus:

      ```SysMDL
      # ENTITY_NAME
      REGEX:
        - [A-Z]{1}[A-Za-z0-9_\-]*
      # KIND
      REGEX:
        - [A-Z]{1}[A-Za-z0-9_\-]*
      # STATE
      EQUALS:
        - exists
      ```
  
- SysMDL introduces the concept of `model types` & `model protocols` which are lexical vessels for capturing functionality & context of concepts & models.

Second we describe the formal definition of the SysMDL Grammar & Syntax: