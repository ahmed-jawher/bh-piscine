# Piscine Go

My solutions to the Reboot01 **Go piscine**, the intensive selection bootcamp that starts the 01 Edu curriculum. It covers shell basics and Git, then core programming in Go: loops, recursion, strings, slices, pointers, maps and higher-order functions. Most functions are written from scratch instead of relying on the standard library.

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![Shell](https://img.shields.io/badge/Shell-4EAA25?logo=gnubash&logoColor=white)

## Topics Covered

| Topic | Example exercises |
|---|---|
| **Shell and Git** | `hello.sh`, `who-are-you.sh`, `myfamily.sh`, `to-git-or-not-to-git.sh`, `look`, `mastertheLS` |
| **Printing and loops** | `printalphabet`, `printreversealphabet`, `printdigits`, `printcomb`, `printcomb2`, `printnbr` |
| **Pointers** | `pointone`, `ultimatepointone`, `divmod`, `ultimatedivmod`, `swap`, `point` |
| **Math and recursion** | `iterativefactorial`, `recursivefactorial`, `iterativepower`, `recursivepower`, `fibonacci`, `sqrt` |
| **Strings and runes** | `strlen`, `strrev`, `firstrune`, `lastrune`, `nrune`, `index`, `compare`, `capitalize`, `toupper`, `tolower`, `alphacount`, `trimatoi` |
| **Character checks** | `isalpha`, `isnumeric`, `islower`, `isupper`, `isprintable`, `isnegative` |
| **Slices** | `appendrange`, `makerange`, `concat`, `concatparams`, `basicjoin`, `join`, `split`, `splitwhitespaces`, `printwordstables` |
| **Command-line programs** | `printprogramname`, `printparams`, `revparams`, `sortparams`, `displayfile`, `nbrconvertalpha`, `boolean` |
| **Higher-order functions** | `foreach`, `map`, `any`, `countif`, `issorted` |

## Getting Started

**Requirements:** Go 1.22+

Library exercises live at the root in the `piscine` package. Programs live in their own folders and print output with [`z01.PrintRune`](https://github.com/01-edu/z01).

```bash
git clone https://github.com/ahmed-jawher/bh-piscine.git
cd bh-piscine
go mod download

# run a program exercise
go run ./printalphabet          # abcdefghijklmnopqrstuvwxyz
go run ./revparams one two three
```

To try a library function, call it from [`test/main.go`](test/main.go):

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Capitalize("hello wORLD from go"))
}
```

```bash
go run ./test
```

## Author

**Ahmed Jawher** — [@ahmed-jawher](https://github.com/ahmed-jawher)
