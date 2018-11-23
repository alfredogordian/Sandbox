package routes

import (
	"fmt"
	"os"
	"strconv"
)

func Calcular() (int){
	input, _ := strconv.Atoi(os.Args[1]);

	for i:=1; i<input ; i++{

		if i%3 ==0 {
			fmt.Println("%v es divisible entre 3" , i)
		}
		if i%5 ==0 {
			fmt.Println(i)
		}
	}

	return 0
}
