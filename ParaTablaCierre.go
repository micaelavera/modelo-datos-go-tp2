package main

import ("fmt"
		"os"
		"log"
)

func main() {

	file, err := os.Create("cierre.txt")
	
	if err != nil {
		log.Fatal("Cannot create file", err)
    }
	defer file.Close()

	for i :=1; i < 13; i++{
		for j := 0; j < 10; j++ {
			fmt.Fprintf(file,"insert into cierre values(2018,%d,%d)\n", i,j)
		}
	}
}
