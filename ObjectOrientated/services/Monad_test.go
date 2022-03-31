package services

import (
	"fmt"
	models "goplayground/objects/models"
	"testing"
)

func TestFunctionalPredicate(t *testing.T) {
	var container = Container[models.YoungPerson]{
		People: []models.YoungPerson{
			models.YoungPerson{
				Person:            models.Person{Message: "hi", Age: 30},
				AdditionalMessage: "I am young Person",
			},
		},
	}

	var peopleOver1 = container.Filter(func(person models.YoungPerson) bool {
		return person.Person.Age > 1
	})

	if len(peopleOver1.People) != 1 {
		t.Fail()
	}

	var peopleOver45 = peopleOver1.Filter(func(person models.YoungPerson) bool {
		return person.Person.Age > 45
	})
	if len(peopleOver45.People) > 1 {
		t.Fail()
	}
}

func TestPolyMorphism(t *testing.T) {

	var person = models.Person{
		Message: "base person",
	}
	var youngPerson models.IPerson = models.YoungPerson{
		Person:            person,
		AdditionalMessage: "I am young",
	}
	var middleAgedPerson models.IPerson = models.MiddleAgedPerson{
		Person:            person,
		AdditionalMessage: "I am middle aged",
	}

	var oldPerson models.IPerson = models.OldPerson{
		Person:            person,
		AdditionalMessage: "I am old",
	}

	fmt.Println(youngPerson)
	fmt.Println(middleAgedPerson)
	fmt.Println(oldPerson)

}

func TestPanicHandling(t *testing.T) {
	var container = Container[models.MiddleAgedPerson]{
		People: []models.MiddleAgedPerson{{
			Person:            models.Person{Message: "hi", Age: 30},
			AdditionalMessage: "I am young Person"},
		},
	}
	container.DoSomePanicing(true)
}

func TestPanicFailing(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	var container = Container[models.OldPerson]{
		People: []models.OldPerson{{
			Person: models.Person{
				Message: "hi",
				Age:     30,
			},
			AdditionalMessage: "I am old",
		}},
	}
	container.DoSomePanicing(false)
}

func TestErrorHandling(t *testing.T) {
	var container = Container[models.YoungPerson]{
		People: []models.YoungPerson{
			models.YoungPerson{
				Person:            models.Person{Message: "hi", Age: 30},
				AdditionalMessage: "I am Young"}},
	}
	result, error := container.DoSomeCheckedExecptions()
	if error == nil || result != "" {
		t.Fail()
	}

}

func BenchmarkStuff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmt.Println(n)
	}
}

func FuzzStuff(f *testing.F) {
	f.Add(1, 2, "life")

	f.Fuzz(func(t *testing.T, a int, b int, c string) {
		if c == "" {
			t.Fail()
		}
	})

}
