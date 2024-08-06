package helloDeps5

import "fmt"

func PrintPhrase(phrase string) string {
	fmt.Println(phrase, "called helloDeps5.PrintPhrase")
	return phrase
}
