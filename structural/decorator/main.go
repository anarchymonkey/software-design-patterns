package main

import "fmt"

type Notifier interface {
	Send() string
}

type ConcreteNotifier struct{}

func (concreteNotifier *ConcreteNotifier) Send() string {
	return "concrete notifier"
}

type NotificationDecorator interface {
	Send() string
}

type FacebookDecorator struct {
	notifier Notifier
}

func (fb *FacebookDecorator) Send() string {
	return "Facebook " + fb.notifier.Send()
}

type EmailDecorator struct {
	notifier Notifier
}

func (email *EmailDecorator) Send() string {
	return "Email " + email.notifier.Send()
}

func main() {
	concreteNotifier := ConcreteNotifier{}

	facebook := &FacebookDecorator{
		notifier: &concreteNotifier,
	}

	email := &EmailDecorator{
		notifier: facebook,
	}

	fmt.Println(email.Send())

}
