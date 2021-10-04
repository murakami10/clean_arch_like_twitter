package main

import (
	"clean_arch/infra"
	"fmt"
)

func main(){
	infra.NewSqlHandler()
	fmt.Println("aaa")
}
