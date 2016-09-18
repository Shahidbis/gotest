package main

import (
	"./greeting"
	"fmt"
	"time"
)

func RenameToFrog(r greeting.Renamable){
	r.Rename("Frog")
}

func main(){
	//var s = greeting.Salutation{"Amy", "hello"}
	salutations := greeting.Saluations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "Hey"},
		{"Mary", "What is up?"},
	}

	//salutations[0].Rename("Sam")
	//RenameToFrog(&salutations[0])
	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	//slice = append(slice, greeting.Salutation{"Frank","Hi "})
	//greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 8)
	go salutations.Greet(greeting.CreatePrintFunction("<C>"), true, 6)
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 6)
	time.Sleep(100 * time.Microsecond)

}

