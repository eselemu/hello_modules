# Hello Modules

This is a repo to test Golang Modules

## Installation
Execute the following command:
bash
go get -u github.com/eselemu/hello_modules


## Usage
go
package main

import (

	"github.com/eselemu/hello_modules"
)

func main() {
  result := sum(5, 3)
  fmt.Println(result)
}
