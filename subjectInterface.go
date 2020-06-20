package main

type SubjectInterface interface {
	Register(obj ObjInterface)
	DelObserver(obj ObjInterface)
	NotifyOberservers()
}
