package pkg

import (
	"fmt"
	"test.com/tests"
)

func SecondTestPackageThatHaveVendor() {
	fmt.Println("Test.org")
	tests.PrintFromVendor()
}
