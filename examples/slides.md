# Welcome to Slides

A Terminal based preview tool for markdown

```go
package main

import "fmt"

func main() {
  fmt.Println("Written in Go!")
}
```

---

# Pager example

## Let’s talk about artichokes

The *artichoke* is mentioned as a garden plant in the 8th century BC by Homer **and** Hesiod. The naturally occurring variant of the artichoke, the cardoon, which is native the the Mediterranean area, also has records of use as a food among the ancient Greeks and Romans. Pliny the Elder mentioned growing of *carduus* in Carthage and Cordoba.

> He holds him with a skinny hand, ‘There was a ship,’ quoth he. ‘Hold off! unhand me, grey-beard loon!’ An artichoke, dropt he.

--Samuel Taylor Coleridge, [The Rime of the Ancient Mariner](https://poetryfoundation.org/poems/43997/)

## Other foods worth mentioning

1. Carrots
2. Celery
3. Tacos
   - Soft
   - Hard
4. Cucumber

## Things to eat today

- Carrots
- Ramen
- Currywurst

### Power levels of the aforementioned foods

| Name       | Power | Comment          |
| ---------- | ----- | ---------------- |
| Carrots    | 9001  | It’s over 9000?! |
| Ramen      | 9002  | Also over 9000?! |
| Currywurst | 10000 | What?!           |

## Currying Artichokes

Here’s a bit of code in [Haskell](https://haskell.org/), because we are fancy. Remember that to compile Haskell you’ll need `ghc`.

```
module Main where

import Data.Function ( (&) )
import Data.List ( intercalculate )

hello :: String -> String
hello s =
    "Hello, " ++ s ++ "."

main :: IO ()
main =
    map hello [ "artichoke", "alcachofa" ] & intercalculate "\n" & putStrLn
```

---

# h1
## h2
### h3
#### h4
##### h5
###### h6

---

# Markdown components

You can use everything in markdown!
* Like bulleted list
* You know the deal

1. Numbered lists too

---

# Tables

| Tables | Too    |
| ------ | ------ |
| Even   | Tables |

---

# Graphs

```
digraph {
    rankdir = LR;
    a -> b;
    b -> c;
}
```

```
┌───┐     ┌───┐     ┌───┐
│ a │ ──▶ │ b │ ──▶ │ c │
└───┘     └───┘     └───┘
```
