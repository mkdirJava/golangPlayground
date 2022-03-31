package models

import "fmt"

type Person struct {
	Message   string
	Age       uint
	HairColor string
}

func (person Person) DoSomthing() string {
	return fmt.Sprintf("")
}

type YoungPerson struct {
	Person            Person
	AdditionalMessage string
}

func (youngPerson YoungPerson) DoSomthing() string {
	return fmt.Sprintf("%s %s", youngPerson.Person.Message, youngPerson.AdditionalMessage)
}

type MiddleAgedPerson struct {
	Person            Person
	AdditionalMessage string
}

func (middleAgedPerson MiddleAgedPerson) DoSomthing() string {
	return middleAgedPerson.AdditionalMessage + middleAgedPerson.Person.Message
}

type OldPerson struct {
	Person            Person
	AdditionalMessage string
}

func (oldPerson OldPerson) DoSomthing() string {
	return oldPerson.AdditionalMessage + oldPerson.Person.Message
}
