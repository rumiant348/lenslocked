package main

import "fmt"

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("woof")
}

type Husky struct {
	Speaker
}

type Speaker interface {
	Speak()
}

type SpeakerPrefixer struct {
	Speaker
}

func (sp SpeakerPrefixer) Speak() {
	fmt.Print("Prefix: ")
	sp.Speaker.Speak()
}

func main() {
	h := Husky{SpeakerPrefixer{Dog{}}}
	h.Speak() // meant h.dog.Speak()
}
