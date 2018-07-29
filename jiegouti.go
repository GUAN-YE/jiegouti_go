package main

import "fmt"


type book struct{
	name string
	address string
	pice int
}

func main(){
	fmt.Println(book{name:"ll",address:"bb",pice:200})
}