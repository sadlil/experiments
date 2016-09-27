package main

import (
	"test.com/tests"
	"test.org/pkg"
)

func main() {
	tests.PrintFromVendor()
	pkg.SecondTestPackageThatHaveVendor()
}
