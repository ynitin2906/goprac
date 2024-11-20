package solidprinciples

import "fmt"

type Shape interface {
	Area()
}

type Sqauare struct {
	side int
}
type Rectangle struct {
	length  int
	breadth int
}

func (s *Sqauare) Area() {
	fmt.Println(s.side * s.side)
}
func (r *Rectangle) Area() {
	fmt.Println(r.length * r.breadth)
}

func RunOpenClosePrinciple() {
	sq := Sqauare{side: 10}
	sq.Area()
	rect := Rectangle{length: 12, breadth: 12}
	rect.Area()
}
