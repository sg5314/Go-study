package main

import "fmt"

func Eat(name string) bool {
	if (name != ""){
		fmt.Println(name)
		return true
	}else{
		return false
	}
}

func main(){
	var name1 string = "GYUDON"
	if ok := Eat(name1); !ok {
		fmt.Println("cannot eat:", name1)
	}

	var name2 string = ""
	if !Eat(name2){
		fmt.Println("cannot eat:", name2)
	} 
}