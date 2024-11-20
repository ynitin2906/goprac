package solidprinciples

import "fmt"

type Brid interface {
	MakeSound()
}
type FlyingBird interface {
	Brid
	Fly()
}
type Sparrow struct{}

func (s *Sparrow) MakeSound() {
	fmt.Println("Sparoww sound")
}
func (s *Sparrow) Fly() {

	fmt.Println("Sparoww flw")
}

type Kiwi struct{}

func (k *Kiwi) MakeSound() {
	fmt.Println("Kiwi sound")
}

func RunLiskovSubstituion() {
	sparrow := Sparrow{}
	sparrow.MakeSound()
	sparrow.Fly()
	kiwi := Kiwi{}
	kiwi.MakeSound()
}
