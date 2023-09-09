package main

import (
	"fmt"

	"github.com/aquasecurity/gobard"
	// "github.com/writeas/go-strip-markdown"
)

func main() {
	PSID := "ZggJZXBCnWM7Z7LDt9Ns_ZPDiClbxPeWr0xxzuO1HZI2kf7NABb9BlbtibU1sSWo0EQ0Lg."
	PSIDTS := "sidts-CjEBSAxbGZVJ7op7v7MGq8bXR7oHM5ZNVyIgb4ZbfMnJMHuHYV3HR4iJAWgMf0twrJrQEAA"

	bard01 := bard.New(PSID, PSIDTS)

	context := ""
	information := ""
	intent := "What is PSID and PSIDT in Google Bard cookie"
	format := "Only give the answer, no more explain"
	prompt := context + ". " + information + ". " + intent + ". " + format

	err := bard01.Ask(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	answer := bard01.GetAnswer()
	fmt.Printf("%s\n", answer)

	// for i := 0; i < bard01.GetNumOfAnswers(); i++ {
	// 	answer := bard01.GetAnswer()
	// 	fmt.Printf("%s\n", answer)
	// 	bard01.Next()
	// }

	// bard01.Next() // will continue the conversation using the first answer as a base
	// err = bard01.Ask("What is the area of that place?")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	// 	os.Exit(1)
	// }
	// for i := 0; i < bard01.GetNumOfAnswers(); i++ {
	// 	answer := bard01.GetAnswer()
	// 	fmt.Printf("%s\n", answer)
	// 	bard01.Next()
	// }
}