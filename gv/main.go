package main

import (
  "flag"
  "fmt"
)

const DES_NAM = "description"
const GIT_DIR = ".git/"

func main() {
  all := flag.Bool("all", false, "Run with everything")
  flag.Parse()

  fmt.Println("")
  if *all {
    compileBasicReport()
  }
} 
