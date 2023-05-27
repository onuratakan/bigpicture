package main

import (
	"bytes"
	"fmt"
	"github.com/ismailbayram/bigpicture/internal/browser"
	"github.com/ismailbayram/bigpicture/internal/graph"
	"os"
	"strings"
)

func main() {
	moduleName := getModuleName()
	rootNode := graph.NewNode(moduleName, ".", graph.Dir, nil)
	tree := graph.NewTree(rootNode)

	if err := browser.Browse(".", moduleName, rootNode, tree); err != nil {
		panic(err)
	}
	tree.GenerateLinks()
	fmt.Println(tree.ToJSON())

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web"))))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, "./web/index.html")
	//
	//})
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	panic(err)
	//}
}

func getModuleName() string {
	file, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	firstLine := string(file[:bytes.IndexByte(file, '\n')])
	return strings.Split(firstLine, " ")[1]
}
