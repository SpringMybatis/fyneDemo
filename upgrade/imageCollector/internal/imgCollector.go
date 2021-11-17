package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)


var listFilePrefix string = "  "

type ImgCollector struct {
	dir string
	outputDir string
	imgTypes []string
}

func (img *ImgCollector) SetDir(dir string) {
	img.dir = dir
}

func (img *ImgCollector) SetOutputDir(dir string) {
	img.outputDir = dir
}

func (img *ImgCollector) SetFilters(filters []string) {
	img.imgTypes = filters
}

func (img *ImgCollector) Run() bool {
	if len(img.imgTypes) == 0 {
		fmt.Println("no filters")
		//return false
	} else {
		fmt.Println(img.imgTypes[0] + " ==")
	}
	// print current directory

	if img.dir == "" {
		//var cwd string
		//re, _ := os.Executable()
		//cwd = filepath.Dir(re)
		//fmt.Println("pwd:" + cwd)
		//img.dir = cwd
		fmt.Println("no input dir")
		return false
	}
	fmt.Println("img.dir:" + img.dir)
	fmt.Println("parent dir:" + filepath.Dir(img.dir))

	// if no outpuDir ,create a new temporary one in cwd
	if img.outputDir == "" {
		pdir := filepath.Dir(img.dir)
		tmpdir, err := ioutil.TempDir(pdir, "imgcollector-")
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("temp dir : " + tmpdir)
		img.outputDir = tmpdir
	}

	// list all files
	pathSeparator := string(os.PathSeparator)
	level := 8
	img.listAllFileByName(level, pathSeparator, img.dir)


	return true
}

func (img *ImgCollector) listAllFileByName(level int, pathSeparator string, fileDir string)  {
	files, _ := ioutil.ReadDir(fileDir)

	tmpPrefix := ""
	for i := 1; i < level; i++ {
		tmpPrefix = tmpPrefix + listFilePrefix
	}

	for _, onefile := range files {
		if onefile.IsDir() {
			//fmt.Printf("\033[34m %s %s \033[0m \n" , tmpPrefix, onefile.Name());
			//fmt.Println(tmpPrefix, onefile.Name(), "目录:")
			fmt.Println(onefile.Name())
			fmt.Println(filepath.Dir(onefile.Name()))
			img.listAllFileByName(level + 1, pathSeparator, fileDir+ pathSeparator+ onefile.Name())
		}else {
			// judge postfix and move it
			postfix := path.Ext(onefile.Name())
			for _, v := range img.imgTypes {
				if postfix == v {
					typePath := img.outputDir + "/" + postfix[1 : len(postfix)]
					_, err := os.Stat(typePath)
					if err != nil {
						if os.IsNotExist(err) {
							// make dir
							os.Mkdir(typePath, os.ModePerm)
						}
					}
					os.Rename(fileDir + "/" + onefile.Name(), typePath + "/" + onefile.Name())
				}
			}

		}
	}
}
