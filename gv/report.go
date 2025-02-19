package main

import (
  "fmt"

  "github.com/fatih/color"
)

func collect(){}

func compileBasicReport() {
  fmt.Println("")
  userInfo()
  branchInfo()
  repoMetaData() 
}

func branchInfo() {
  color.Green("Branches -     %s", getFiles("packed-refs")) 
  color.Green("Cur branch -    %s", getCurrentBranch("HEAD"))
}

func repoMetaData() {
  color.Blue("Useable hooks -   %s", getUseableHooks("hooks"))
  color.Blue("Repo Des -     %s", readFile(DES_NAM))
}

func userInfo() {
  color.Red("OS   -   %s", getOS())
  color.Red("Sys Arch  -   %s", getArch())
  color.Red("Cur Dir -    %s", getCurrentDir())
  color.Red("Exe Dir -    %s", getExecutableDir())
}
