package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	path = flag.String("p", "", "Path")
	page = flag.String("s", "", "Page")
)

func main() {
	flag.Parse()
	/*out, _ := exec.Command("wget", "--mirror", "--convert-links", "--page-requisites", "--no-parent", "-P", "home/", "http://svecha.69.tilda.ws/").Output()
	fmt.Println(string(out))*/
	if *path != "" && *page != "" {
		out, _ := exec.Command("wget", "--mirror", "--convert-links", "--page-requisites", "--no-parent", "-P", *path, *page).Output()

		fmt.Println("succses", string(out))
	}
}

/*
wget --mirror --convert-links --page-requisites --no-parent -P home/ http://svecha.69.tilda.ws/
*/
