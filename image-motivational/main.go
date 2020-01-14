package main

import "fmt"

func main() {

	n := []string{"xyz"}
	//var n []string
	f(n)
	fmt.Println(n)
}

func f(a []string) {

	a[0] = "xyz2"
	a = nil
	a = append(a, "abc")

}
