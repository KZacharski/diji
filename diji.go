package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	var projname string

	fmt.Println("地基")
	fmt.Println("diji 0.1-BETA")
	fmt.Println("by KZacharski")
	fmt.Print("Project name: ")
	fmt.Scanln(&projname)
	var createcss bool = true
	var createjs bool = true
	configbytes, err := os.ReadFile(".diji-config/config.txt")
	if err != nil {
		fmt.Print(err)
	}
	configtext := string(configbytes)
	var quickmode bool = strings.Contains(configtext, "quick-mode = true")
	fmt.Println(quickmode)

	if quickmode == false {
		var cssstr string
		var jsstr string
		fmt.Print("Create a css file (y/n, default y): ")
		fmt.Scanln(&cssstr)
		fmt.Print("Create a js file (y/n, default y): ")
		fmt.Scanln(&jsstr)
		if cssstr == "n" {
			createcss = false
		}
		if jsstr == "n" {
			createjs = false
		}
	}

	if err := os.Mkdir(projname, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	var indname string = projname + "/index.html"
	index, err := os.Create(indname)
	if err != nil {
		log.Fatal(err)
	}
	var assetspath string = projname + "/assets"
	if err := os.Mkdir(assetspath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	var favpath string = assetspath + "/favicon.png"
	copy(".diji-config/defaultfav.png", favpath)
	fmt.Println(favpath + " created.")
	var indexcontent1 string = `<!DOCTYPE html>
<html>
<head>
<title>` + projname + `</title>
<link rel="icon" type="image" href="./assets/favicon.png">
`
	var csstag string = `<link rel="stylesheet" type="text/css" href="style.css">
`
	var jstag string = `<script src="script.js"></script>
`
	var indexcontent2 string = `</head>
</html>`
	if createcss == true {
		indexcontent1 = indexcontent1 + csstag
	}
	if createjs == true {
		indexcontent1 = indexcontent1 + jstag
	}
	var indexcontent string = indexcontent1 + indexcontent2
	_, err2 := index.WriteString(indexcontent)
	check(err2)
	defer index.Close()
	fmt.Println(indname + " created.")

	if createcss == true {
		var cssname string = projname + "/style.css"
		var csscontent string = `body {
font-family: sans-serif;
}`
		style, err := os.Create(cssname)
		if err != nil {
			log.Fatal(err)
		}
		_, err2 := style.WriteString(csscontent)
		check(err2)
		defer style.Close()
		fmt.Println(cssname + " created.")
	}
	if createjs == true {
		var jsname string = projname + "/script.js"
		script, err := os.Create(jsname)
		if err != nil {
			log.Fatal(err)
		}
		defer script.Close()
		fmt.Println(jsname + " created.")
	}

}
