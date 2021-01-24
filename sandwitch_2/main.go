package main

import (
	"fmt"
	_ "time"
)

type SpoonOfJam interface {
	String() string
}

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type Bread struct {
	val string
}

type StrawberryJam struct {
}

type OrangeJam struct {
}

type AppleJam struct {
}

func (a *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}

}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}

}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry "

}

type SpoonOfOrangeJam struct {
}

type SpoonOfAppleJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ orange"
}

func (a *SpoonOfAppleJam) String() string {
	return "+ Apple"

}
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	jam := &AppleJam{}
	bread.PutJam(jam)

	fmt.Println(bread)

}
