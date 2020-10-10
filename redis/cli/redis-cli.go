package main

import (
	"fmt"
	"os"
)

func main()  {
	var name string
	var age int
	fmt.Scanf("%s%d", &name, &age)
	fmt.Println(name, age)


	fmt.Println(os.Args)
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	for i, args := range os.Args {
		fmt.Printf("args[%d]=%s\n",i,args)
	}
}
