package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

type user struct{
	Name string
	//city string
	Phone string
	DOB  string

}


func main(){
	 v :=user{faker.Name(),faker.Phonenumber(),faker.Date()}
	 fmt.Println(v)


}