package main

import (
	"fmt"

	"example.com/myproject/calc"
	"example.com/myproject/mathutils"
	"example.com/myproject/stringutils"
	"example.com/myproject/visibility"
)

func main() {
	fmt.Println("Сложение:", mathutils.Add(4, 9))
	fmt.Println("Обратная строка:", stringutils.Reverse("Orange"))
	fmt.Println("Квадрат 7:", calc.Square(7))

	visibility.PublicFunc()
	//visibility.PrivateFunc()
}
