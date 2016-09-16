package main

import "./greeting"

func main(){
	//var s = greeting.Salutation{"Amy", "hello"}
	salutations := greeting.Saluations{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "Hey"},
		{"Mary", "What is up?"},
	}

	salutations[0].Rename("Sam")
	//slice = append(slice, greeting.Salutation{"Frank","Hi "})
	//greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 8)
	salutations.Greet(greeting.CreatePrintFunction("?"), true, 8)


}

