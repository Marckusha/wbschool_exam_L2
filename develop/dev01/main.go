package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR!", err)
		return
	}

	fmt.Println(time)
}
