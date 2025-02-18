package main

import (
  "fmt"
)

func collect(){}

func generateBasicReport() {
  getBranches() 
  fmt.Println("Cur branch -     ", getCurrentBranch("HEAD"))
  fmt.Println("Useable hooks -    ", getUseableHooks("hooks"))
  fmt.Println("Repo Des -     ", readFile(DES_NAM))
  fmt.Println("OS   -   ", getOS())
  fmt.Println("Sys Arch  -   ", getArch())
  fmt.Println("Cur Dir -    ", getCurrentDir())
  fmt.Println("Exe Dir -    ", getExecutableDir())
}

