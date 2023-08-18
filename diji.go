package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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
	var createsample bool = false
	var initgit bool = true
	var creategitignore bool = true
	configbytes, err := os.ReadFile(".diji-config/config.txt")
	if err != nil {
		fmt.Print(err)
	}
	configtext := string(configbytes)
	var quickmode bool = strings.Contains(configtext, "quick-mode = true")
	var debug bool = strings.Contains(configtext, "debug = true")

	var cssstr string
	var jsstr string
	var samplestr string
	var langstr string = "en"
	var gitstr string
	var gitignorestr string
	var gifiles string

	if quickmode == false {

		fmt.Print("Create a css file (y/n, default y): ")
		fmt.Scanln(&cssstr)
		fmt.Print("Create a js file (y/n, default y): ")
		fmt.Scanln(&jsstr)
		fmt.Print("Website language (default en): ")
		fmt.Scanln(&langstr)
		fmt.Print("Insert sample content (y/n, default n): ")
		fmt.Scanln(&samplestr)
		fmt.Print("Initialize a git repo (y/n, default y): ")
		fmt.Scanln(&gitstr)
		fmt.Print("Create .gitignore (y/n, default y): ")
		fmt.Scanln(&gitignorestr)
		fmt.Print("Add files/file types to .gitignore: ")
		fmt.Scanln(&gifiles)
		if cssstr == "n" {
			createcss = false
		}
		if jsstr == "n" {
			createjs = false
		}
		if samplestr == "y" {
			createsample = true
		}
		if gitstr == "n" {
			initgit = false
		}
		if gitignorestr == "n" {
			creategitignore = false
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
<html lang="` + langstr + `">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="ie=edge">
<title>` + projname + `</title>
<link rel="icon" type="image" href="./assets/favicon.png">
`
	var csstag string = `<link rel="stylesheet" type="text/css" href="style.css">
`
	var jstag string = `<script src="script.js"></script>
`
	var indexcontent2 string = `</head>
<body>
</body>
</html>`
	var samplecontent string = `<h1>` + projname + `</h1>
<h3>Generated with diji</h3>
<p>Website content</p>`
	if createsample == true {
		indexcontent2 = `</head>
<body>
` + samplecontent + `
</body>`
	}
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

	var cssname string
	var csscontent string

	if createcss == true {
		cssname = projname + "/style.css"
		csscontent = `body {
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

	var jsname string
	if createjs == true {
		jsname = projname + "/script.js"
		script, err := os.Create(jsname)
		if err != nil {
			log.Fatal(err)
		}
		defer script.Close()
		fmt.Println(jsname + " created.")
	}

	var giname string
	var gicontent string

	if creategitignore == true {
		giname = projname + "/.gitignore"
		gicontent = ".DS_Store " + gifiles
		gignore, err := os.Create(giname)
		if err != nil {
			log.Fatal(err)
		}
		_, err2 := gignore.WriteString(gicontent)
		check(err2)
		defer gignore.Close()

		fmt.Println(giname + " created.")
	}

	if initgit == true {
		cmd := exec.Command("git", "init", projname)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		cmd1 := exec.Command("git", "add", ".")
		cmd1.Dir = "./" + projname
		if err := cmd1.Run(); err != nil {
			log.Fatal(err)
		}

		cmd2 := exec.Command("git", "commit", "-a", "-m", `"Initial commit"`)
		cmd2.Dir = "./" + projname
		if err := cmd2.Run(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Git repo initialized.")
	}

	if debug == true {
		fmt.Println("projname(string): " + projname)
		fmt.Print("createcss(bool): ")
		fmt.Print(createcss)
		fmt.Print("createjs(bool): ")
		fmt.Print(createjs)
		fmt.Print("configbytes([]byte]): ")
		fmt.Print(configbytes)
		fmt.Println("configtext(string): " + configtext)
		fmt.Print("quickmode(bool): ")
		fmt.Print(quickmode)
		fmt.Print("debug(bool): ")
		fmt.Print(debug)
		fmt.Println("cssstr(string): " + cssstr)
		fmt.Println("jsstr(string): " + jsstr)
		fmt.Println("indname(string): " + indname)
		fmt.Println("assetspath(string): " + assetspath)
		fmt.Println("favpath(string): " + favpath)
		fmt.Println("indexcontent1(string): " + indexcontent1)
		fmt.Println("csstag(string): " + csstag)
		fmt.Println("jstag(string): " + jstag)
		fmt.Println("indexcontent2(string): " + indexcontent2)
		fmt.Println("indexcontent(string): " + indexcontent)
		fmt.Println("cssname(string): " + cssname)
		fmt.Println("csscontent(string): " + csscontent)
		fmt.Println("jsname(string): " + jsname)
	}

	fmt.Println("Project created in ./" + projname + ".")
	fmt.Println("Thanks for using diji.")

}
