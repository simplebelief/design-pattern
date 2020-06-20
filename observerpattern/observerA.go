package main

import (
	"fmt"
)

type ObserverA struct {
	name string
}

func (obsa *ObserverA) SetName(name string) {
	obsa.name = name
}

func (obsa *ObserverA) Read(message string) {
	fmt.Println(obsa.name + " read: " + message)
}

func (obsa *ObserverA) Sub(subject *SubjectImpl) {
	subject.Register(obsa)
}

func (obsa *ObserverA) Unsub(subject *SubjectImpl) {
	subject.DelObserver(obsa)
}
