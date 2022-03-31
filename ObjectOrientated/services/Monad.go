package services

import (
	"errors"
	"fmt"
	models "goplayground/objects/models"
)

type Container[T models.IPerson] struct {
	People []T
}

func (container Container[T]) Filter(pred func(T) bool) Container[T] {
	var positivePeople []T = make([]T, 0)
	for _, containerPerson := range container.People {
		if pred(containerPerson) {
			positivePeople = append(positivePeople, containerPerson)
		}
	}
	return Container[T]{People: positivePeople}
}

func (container Container[T]) DoSomePanicing(shouldRecover bool) {

	defer func() {
		if shouldRecover {
			if r := recover(); r != nil && shouldRecover {
				fmt.Println("Recovered in f", r)
			}
		}
	}()

	panic("Somthing went Wrong")
}

func (container Container[T]) DoSomeCheckedExecptions() (string, error) {
	return "", errors.New("Something went wrong")
}
