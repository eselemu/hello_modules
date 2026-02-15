# Hello Modules

This is a repo to test Golang Modules

## Installation
### Create a new directory for your project
mkdir myproject
cd myproject

### Initialize a new Go module
go mod init myproject

### Download and install the module
go get github.com/eselemu/hello_modules

### Run the main program
go run main.go

## Program example
package main

import (
    "fmt"
    
    "github.com/eselemu/hello_modules"
)

func main() {
    result := hello_modules.Sum(5, 3)
    fmt.Println(result) // Output: 8
}
