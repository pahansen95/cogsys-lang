# Programming & Systems Theory

## What & Why

A Type is a grouping of rules, functionality & boundaries that can be applied towards data.

In essence, a type assigns semantics to raw data. For example, an Integer is a numerical whole number while a string is a collection of characters in a language; in both cases the raw data of each type is represented as a set of bytes in memory.

Computer Programming is the act of instructing a computer to operate in a desired way; it is a tool. Generally speaking, we don't desire the computer to shift bit values or move data between memory registries. Instead we desire the application of some logical or cognitive concept. We as humans, as sources of intelligence, assign semantics to the state of the computer. Through these semantics, we model our logical or cognitive systems through the capabilities of the computer. If we conceptualize an infinitely complex problem, then the computer will execute infinitely. If we conceptualize a finite problem then the computer will execute in polynomial time.

This implies that we as a source of intelligence must:

- properly model our logical or cognitive concepts to achieve our desired outcomes.
- properly understand how a computer, as a tool, operates & functions.

Therefore, we must properly instruct the computer how to adhere to the semantics of our cognitive models. These instructions are a middleware between our cognitive systems & the computing systems.

In otherwords, a program or script is really a translation of a cognitive system into a computing system.

Types therefore are ways we conceptualize the capabilities & limits of specific systems. As such, types belong to discrete systems; the type sets of cognitive systems & computing systems are disjoint as they operate inside different models.

Accurately & Precisely modeling each system, cognitive & computing, is essential to successfullying achieving our desired results. I assert that a common problem encountered by programmers is the mixing of the two systems. Capabilities & limitations from one are incorrectly applied to the other which cause a class of errors in our programs; we are able to define illegal system states & infer capabilities & limitations which enact on these illegal system states to achieving our desired results.

I postulate that to properly develop programs that translate cognitive systems to computing systems, we must:

- Have a well defined language that models the capabilities & limitations of the computing system we execute on.
- Have a well defined language that models the capabilities & limitations of the cognitive system we are applying.
- Have a well defined process or framework that translates a cognitive system into a computing system.
  - This process or framework must enforce strict delinations between different systems as to avoid overlapping

If we can enforce these principles then we can better produce our desired outcomes:

- We force ourselves to articulate & understand the systems we model.
- We can scope our problem solving process to a single system, thereby simplifying our solution & reducing time spent on "illegal" solutions.
- We can grow or shrink either system to in a tradeoff between system complexity & solution time.

Re; About `well defined process or framework that translates a cognitive system into a computing system`. This middleware must explicitly define how the cognitive system is represented by the computing system.

In the cognitive system we define conceptual ideas as types. Types encode semantics & capabilities of the system's substates; meaning, functionality & limitations:

- Meaning encodes representation of the data that constitues the substate.
- functionality is work applied to or applied by the system to induce a state change (ie. A side effect).
- Limitations are implicit: if a meaning or functionality is not defined, then the limitations of the type is the boundary encapsulating these definitions.

Let's unpack this middleware. In our desire is to represent our cognitive systems through computing systems then fundamentally, conceptual types must be encoded through computing types. Cognitive types can be efficiently represented as a hierarchy of other cognitive types but after expansion a cognitive type must map to a set of computing types. This can be a 1 to 1 mapping or a 1 to many mapping. Cognitive Meaning is therefore a composite of computing meaning. Likewise cognitive functionality implies the work done on or by a cognitive system is a composite of the functionality of the consituent computing types. Since Cognitive limitations are implied, the computing limitations is a byproduct of the boundaries of the union the set of computing types; in application the cognitive limitation is a subset of the computing limitation.

## How

How do we implement such a language that enforces the deliniation of cognitive systems & computing systems & then provides a framework to represent cognitive models in terms of computing models?

First, we must model the computing system & the cognitive system. Each model must define Types in order to establish the semantics & capabilities of the model. Next we must declare how the cognitive system is represented by the computing system.

> I suspect that in implementation, the cognitive system can be represented by the computing system in multiple ways.

I propose, instead of a single language, we instead implement a family of languages, each tailored to a specific intent. Holistically they will implement the design laid out:

- Computing Systems Language: Describes the semantics & functionality of the computer
- Cognitive Systems Language: Describes the semantics & functionality of the conceptual domain
- Systems Shim Language: Maps the Cognitive System's implementation to the Computing System's implementation.

Each language shall be independent & a "programming language" a developer uses to build part of the overall solution. No language is designed to be "calculated" by a software system or automation. Each language is designed for an intelligent entity to encode, as source code, their knowledge & application thereof.

Holistically, the computing system & cognitive system are modeled in their respective languges. Then the `shim` is written such that it maps the cognitive system to the computing system. This process is iterated until the maintainer identifies the holistic program meets their success criteria. In practice, application of this framework should generate modular & decoupled systems designs. There can exist many Computational Systems Designs that implement instructions specific to hardware or technologies (ie. Single Computer vs Distributed Systems). There can also exist many Cognitive Systems Designs that model common concepts & frameworks delineated by domain (ie. Manufacturing Automation vs Financial Systems). However, (I postulate) Shims should be unique in their implementation as they map specific cognitive model implementations to specific computational model implementations.

In practice, computer engineers would maintain the computing systems model, "domain" engineers maintain the Cognitive Systems model & the `shim` would be maintained through a collaboartion between software engineers, "domain" engineers & computer engineers.

> Note: `computer`, `domain` & `software` engineers are roles that can be fulfilled by individuals, or by teams; there is no distinction.

### Computing Systems Language

> TODO: Need to think through this PoV more.
>
> First we focus on a language for modeling computing systems. At the most fundamental level, this is assembly. We can simplify our approach however & use a higher level general purpose computer language that is turing complete. For fast prototyping & providing a "reference" implementation we can use a dynamic language like Python or even Lua. However, for "production" implementations we should use C or even go so far as to write optimized implementations using the target system's assembly language.
>
> The goal of this language is to provide a set of tooling & frameworks that can be used to instruct a computer to operate. It should make no declarations beyond the capabilities & functionality of the computer itself. This language must declare a set of types that directly map to the functionality of the computer. Composite types (ex. A table) must describe semantics & functionality for operating the computer or of the computer itself.

### Cognitive Systems Language

Next we focus on a language for modeling cognitive systems whose goal is to allow the domain engineer to model conceptual systems of ideas. It defines a minimalistic syntax that allows the system designer to craft a language superset tailored to their cogntive model. This lexical framework then allows the developer to properly articulate a system's semantics & functionality (ie cognitive types). It should make no declarations or assumptions how the underlying computer functions or operates, purely focusing on articualtion of ideas, concepts & processes.

> This doesn't preclude a developer from cognitively modeling a computer in such a format that it can be compiled to machine code, but the cognitive systems language itself is agnostic to the application thereof.

### Systems Shim Language

Lastly we focus on a language that declaratively represents how cognitive systems are represented by computing systems. This language does not implement new types but purely maps cognitive semantics & functionality to a composite of computing semantics & functionality.

The language should not provide a way to declare new types. Instead it must provide a framework for describing how cognitive types are implemented in terms of computing typyes: how is cognitive data represented & how is cognitive functionality enacted.