package main

import "fmt"

func main()  {

	fmt.Println("Hello world")
	
}

func GetHostName() string {
	name, _ := os.Hostname()
	return name
}
