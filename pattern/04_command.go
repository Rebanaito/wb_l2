package main

import (
	"fmt"
)

type app interface {
	userLogIn()
	userLogOut()
	isOn() bool
}

type wildberries struct {
	loggedIn bool
}

func (w *wildberries) userLogIn() {
	w.loggedIn = true
	fmt.Println("Logging into Wildberries")
}

func (w *wildberries) userLogOut() {
	w.loggedIn = false
	fmt.Println("Logging out of Wildberries")
}

func (w *wildberries) isOn() bool {
	return w.loggedIn
}

type command interface {
	execute()
}

type logInCommand struct {
	app app
}

func (logIn *logInCommand) execute() {
	if !logIn.app.isOn() {
		logIn.app.userLogIn()
	}
}

type logOutCommand struct {
	app app
}

func (logOut *logOutCommand) execute() {
	if logOut.app.isOn() {
		logOut.app.userLogOut()
	}
}

type button struct {
	command command
}

func (b button) press() {
	b.command.execute()
}

// func main() {
// 	application := &wildberries{}
// 	login := &logInCommand{application}
// 	logout := &logOutCommand{application}
// 	inButton := &button{login}
// 	outButton := &button{logout}
// 	inButton.press()
// 	fmt.Println("Ordering some goodies...")
// 	time.Sleep(3 * time.Second)
// 	outButton.press()
// }
