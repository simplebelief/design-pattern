package main

type ObjInterface interface {
	Read(message string)
	Sub(subject *SubjectImpl)
	Unsub(subject *SubjectImpl)
}
