package main

import "fmt"

func main() {
	a := parse("SEND + MORE = MONEY")
	fmt.Println(a.addends, a.result, string(a.letters), a.lead)
}
