package visibility

import "fmt"

func PublicFunc() {
	fmt.Println("Это публичная функция")
	PrivateFunc()
}

func PrivateFunc() {
	fmt.Println("Это приватная функция")
}
