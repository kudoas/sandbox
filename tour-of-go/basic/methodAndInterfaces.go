package basic

import (
	"fmt"
	"math"
)

// classはない
// あくまで型に紐づく
type Matrix struct {
	X, Y float64
}

// Abs method
// レシーバはfunc キーワードとメソッド名の間に自身の引数リストで表現
func (m Matrix) Abs() float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

func AbsFunc(m Matrix) float64 {
	return math.Sqrt(m.X*m.X + m.Y*m.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Methods and pointer indirection

// !!!!!
// pointer receiversを持つmethod
// ポインタレシーバを持つメソッドはレシーバが指す変数を変更できる！

// ポイントレシーバを使う理由
// 1. メソッドがレシーバー先の変数を変更するため
// 2. メソッドの呼び出し毎に変数のコピーを避けるため
func (m *Matrix) ScaleP(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

// 一般的には、変数レシーバ、または、ポインタレシーバのどちらかですべてのメソッドを与え、混在させるべきではありません。
// receiverを直接変更しない！
func (m Matrix) Scale(f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

// function
// *Matrixだからmはpointer型！
// ポインタを引数にとる！
func ScaleFunc(m *Matrix, f float64) {
	m.X = m.X * f
	m.Y = m.Y * f
}

// & point型, *
func CheckScale() {
	m := Matrix{3, 4}

	// ScalePはポイントレシーバを持っているため、(&m).ScaleP(10)で解釈される仕様
	m.ScaleP(10) // (&m).ScaleP(10)
	fmt.Println(m)

	ScaleFunc(&m, 20.0)
	fmt.Println(m)
}

func CheckAbs() {
	m := Matrix{10, 20}
	fmt.Println(AbsFunc(m))

	//　ポイントレシーバを持っていない場合はポイントを渡すと(*(&m)).Abs()として解釈される
	fmt.Println(m.Abs())
	fmt.Println((&m).Abs())
}

// interface
// interface(インタフェース)型は、メソッドのシグニチャの集まり
type Decer interface {
	Dec() float64
}

func (m *Matrix) Dec() float64 {
	return m.X - m.Y
}

func CheckDec() {
	var a Decer
	m := Matrix{2, 3}
	a = &m
	fmt.Println(a.Dec())
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func CheckM() {
	var i I = T{"Hello"}
	i.M()
}
