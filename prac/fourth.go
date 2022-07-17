package main

import (
	"fmt"
	"strings"
)

func main() {
	//words := strings.Fields("This , that, and the other.")
	//for i, word := range words {
	//	fmt.Println(i, "=>", word)
	//}

	s := "This, that, and the other."
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	s = replacer.Replace(s)
	Nwords := strings.Fields(s)
	for _, word := range Nwords {
		fmt.Println(word)
	}
}
