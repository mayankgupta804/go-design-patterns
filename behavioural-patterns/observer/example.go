package main

func main() {
	listener1 := DataListener{Name: "Listener 1"}
	listener2 := DataListener{Name: "Listener 2"}

	subject := DataSubject{}

	subject.RegisterObserver(listener1)
	subject.RegisterObserver(listener2)

	subject.ChangeItem("monday")
	subject.ChangeItem("wednesday")
	subject.ChangeItem("friday")

	subject.UnRegisterObserver(listener2)
	subject.ChangeItem("sunday")
	subject.ChangeItem("saturday")
}
