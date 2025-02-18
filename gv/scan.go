package main

import (
  "fmt"
  "runtime"
  "os"
  "path/filepath"
  "io/ioutil"
  "log"
  "strings"
)

// Function to get the current working directory
func getCurrentDir() string {
	dir, err := os.Getwd() // Get the current working directory
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return dir + "/"
}

// Function to get the directory of the current executable
func getExecutableDir() string {
	execPath, err := os.Executable() // Get the path of the current executable
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return filepath.Dir(execPath) // Extract the directory from the executable path
}

func getBranches() {
  var path = "../" + GIT_DIR + "/refs/heads"

  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {
    fmt.Println(file)
  }
}

func getCurrentBranch(path string) string {
  path = "../" + GIT_DIR + path
  content, err := ioutil.ReadFile(path)

  if err != nil {
    log.Fatal(err)
  }

  result := strings.Split(string(content), "/")
  return result[(len(result) - 1)]
}

func getUseableHooks(path string) int {
  var useablehooks = 0
  path = "../" + GIT_DIR + path 
  files, err := ioutil.ReadDir(path)

  if err != nil {
    log.Fatal(err)
  }

  for _, file := range files {
    
    if !strings.Contains(string(file.Name()), ".sample") {
      useablehooks += 1
      fmt.Println(file.Name())
    }
  }

  return useablehooks
}

// func getStaged() []string {}

func readFile(path string) string {
  path = "../" + GIT_DIR + path
  
  data, err := ioutil.ReadFile(path)

  if err != nil {
    log.Fatal(err)
  }
  return string(data)
}

func getOS() string {
  return runtime .GOOS
}

func getArch() string {
  return runtime.GOARCH
}
