package Lv1_And_Lv2

import "fmt"

// Dog 一只狗
type Dog struct {
	Name  string
	Age   int8
	Color string
}

type DogBehavior interface {
	Bark()
	Eat()
	Jump()
}

func (b *Dog) Bark() {
	fmt.Println("汪汪汪")
}
func (b *Dog) Eat() {
	fmt.Println(b.Name, "想吃东西了")
}
func (b *Dog) Jump() {
	fmt.Println(b.Name, "跳了起来")
}

/*func main() {
	var dog1 Lv1_And_Lv2.DogBehavior = &Lv1_And_Lv2.Dog{
		Name:  "小白",
		Age:   2,
		Color: "黑色",
	}
	dog1.Bark()
	dog1.Jump()
}*/
