package main

import "fmt"

type Student struct {
	Name   string
	Number int
	Grade  int
}

type Teacher struct {
	Name string
}

func (s Student) getEmail() string {
	return s.Name + "@std.com"
}

func (t Teacher) getEmail() string {
	return t.Name + "@tcr.com"
}

type Person interface {
	getEmail() string
}

func sendEmail(p Person) (context string) {
	from := p.getEmail()
	context = `
  送信元 : ` + from + `
  これはテスト用のメールです。
  よろしくお願いします。
  `
	return context
}

func main() {
	var s, t Person
	s = Student{
		Name:   "hoge",
		Number: 1,
		Grade:  3,
	}
	t = Teacher{
		Name: "fuga",
	}
	fmt.Println(sendEmail(s))
	fmt.Println(sendEmail(t))
}
