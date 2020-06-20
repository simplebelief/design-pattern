// ObserverPattern project main.go
package main

func main() {
	var newsOffic SubjectImpl //主题
	var xiaoA ObserverA
	xiaoA.SetName("xiaoA")
	xiaoA.Sub(&newsOffic)
	var xiaoB ObserverB
	xiaoB.SetName("xiaoB")
	xiaoB.Sub(&newsOffic)

	message := "shop is on saling"
	newsOffic.Update(message)
	xiaoB.Unsub(&newsOffic)

	message = "get together to see sea"
	newsOffic.Update(message)

}
