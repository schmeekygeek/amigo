package main

import (
  "amigo/repl"
  "os"
)

func main() {
  repl.Start(os.Stdin, os.Stdout)
}
