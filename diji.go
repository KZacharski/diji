package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
	var indexcontent string = `<!DOCTYPE html>
<html>
<head>
<title>` + projname + `</title>
</head>
<html>`
	_, err2 := index.WriteString(indexcontent)
	check(err2)
	defer index.Close()
	fmt.Println(indname + " created.")
	configbytes, err := os.ReadFile(".diji-config/config.txt")
	if err != nil {
		fmt.Print(err)
	}
	configtext := string(configbytes)
	var quickmode bool = strings.Contains(configtext, "quick-mode = true")
	fmt.Println(quickmode)
}
