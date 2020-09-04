package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tkiraly/gonames/packagename"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please supply a package name as an argument")
		os.Exit(1)
	}
	name := os.Args[1]
	ok, err := packagename.Ok(name)
	if err != nil {
		log.Fatalln(err)
	}
	if ok {
		fmt.Printf("%s is a good package name.\n", name)
	} else {
		fmt.Printf("%s is a BAD package name.\n", name)
	}
}
