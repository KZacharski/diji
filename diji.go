package main

import (
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var projname string

	fmt.Println("地基")
	fmt.Println("diji 0.1-BETA")
	fmt.Println("by KZacharski")
	fmt.Print("Project name: ")
	fmt.Scanln(&projname)
	if err := os.Mkdir(projname, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	var indname string = projname + "/index.html"
	index, err := os.Create(indname)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := index.WriteString("test")
	check(err2)
	defer index.Close()
	fmt.Println(indname + " created.")
}
