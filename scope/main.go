package main

import "myapp/packageone"

var myVar = "Package level var."

func main() {

	var blockVar = "block var."

	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)

}
