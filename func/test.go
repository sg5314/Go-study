package main

import "fmt"

func Eat(name string) (bool ,error) {
	if name == ""{
		return false, fmt.Errorf("name is empty")
	}
	fmt.Println(name)
	return true, nil //nil = 空の状態
}

// main関数がないため、エラーになる