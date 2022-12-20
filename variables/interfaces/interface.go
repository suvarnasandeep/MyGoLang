package interfaces

import "fmt"

type Animal interface {
	says() string
	howManyLegs() int
}

type Dog struct {
	name  string
	sound string
	legs  int
}

func (d *Dog) says() string {
	return d.sound
}

func (d *Dog) howManyLegs() int {
	return d.legs
}

type Cat struct {
	name    string
	sound   string
	legs    int
	hasTail bool
}

func (c *Cat) says() string {
	return c.sound
}

func (c *Cat) howManyLegs() int {
	return c.legs
}

func TestInterface() {
	dog := Dog{
		name:  "dog",
		sound: "woof",
		legs:  4,
	}
	riddle(&dog)

	var cat Cat
	cat.name = "cat"
	cat.sound = "meow"
	cat.legs = 4
	cat.hasTail = true

	riddle(&cat)
}

func riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has "%d" legs`, a.says(), a.howManyLegs())
	fmt.Println(riddle)
}
