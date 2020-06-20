package main

import (
	"fmt"
)

type ObserverB struct {
	name string
}

func (obsb *ObserverB) SetName(name string) {
	obsb.name = name
}

func (obsb *ObserverB) Read(message string) {
	fmt.Println(obsb.name + " read: " + message)
}

func (obsb *ObserverB) Sub(subject *SubjectImpl) {
	subject.Register(obsb)
}

func (obsb *ObserverB) Unsub(subject *SubjectImpl) {
	subject.DelObserver(obsb)
}
