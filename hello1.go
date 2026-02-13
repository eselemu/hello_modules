package hello_modules

import (
	"fmt"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello %v from hello1", name)
	return message
}
