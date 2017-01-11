// This is Header Comment
package src

import "fmt"

// This is a type Comment
//+comment for machine
type TestType struct {
	// This is field Comment
	Field string
}

func (t *TestType) TestFunc() {
	fmt.Println("Do Nothing")
}

func TestFunc() {
	fmt.Println("To Level Nothing")
}
