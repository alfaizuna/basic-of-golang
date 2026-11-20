package main

import "fmt"

// DescribePackage mendeskripsikan tipe package berdasarkan status isMain.
func DescribePackage(name string, isMain bool) string {
	// TODO: Kembalikan "<name> (executable)" jika isMain true, selain itu "<name> (library)"
	return ""
}

func main() {
	fmt.Println(DescribePackage("main", true))
	fmt.Println(DescribePackage("mathutil", false))
}
