package main

import "./greeting"

func main(){
	//var s = greeting.Salutation{"Amy", "hello"}
	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Amy", "Hey"},
		{"Mary", "What is up?"},
	}
	greeting.Greet(slice, greeting.CreatePrintFunction("?"), true, 8)
}

