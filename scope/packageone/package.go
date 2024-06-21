package packageone

import "fmt"

var PackageVar = "Package Var exported."

func PrintMe(s1, s2, s3 string) {
	fmt.Println(s1, s2, s3)
}
