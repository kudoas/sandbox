// package main

// import "fmt"

// type Student struct {
// 	Name   string
// 	Number int
// 	Grade  int
// }

// type Teacher struct {
// 	Name string
// }

// type Person interface {
// 	getEmail() string
// }

// func main() {
// 	var s, t Person

// 	s = Student{
// 		Name:   "Yamada",
// 		Number: 999,
// 		Grade:  5,
// 	}
// 	t = Teacher{
// 		Name: "Tsubomi",
// 	}

// 	cxtStu := sendEmail(s)
// 	fmt.Println(cxtStu)

// 	cxtTea := sendEmail(t)
// 	fmt.Println(cxtTea)
// }

// func (s Student) getEmail() string {
// 	return s.Name + "@student.ed.jp"
// }
// func (t Teacher) getEmail() string {
// 	return t.Name + "@teacher.ed.jp"
// }

// func sendEmail(p Person) (context string) {
// 	from := p.getEmail()
// 	context = `
//   送信元 : ` + from + `
//   これはテスト用のメールです。
//   よろしくお願いします。
//   `
// 	return context
// }
