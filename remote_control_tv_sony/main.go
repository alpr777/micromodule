package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Print(
			"====================================================", "\n",
			"./[app name] [sony tv ip] [pass] [command]", "\n",
			"----------------------------------------------------", "\n",
			"EXAMPLE:", "\n",
			"./[app name] 192.168.1.55 1111 poweroff", "\n",
			"====================================================", "\n",
		)
		return
	}
	sony := &remote{}
	sony, err := sony.NewRemote(strings.TrimSpace(os.Args[1]), strings.TrimSpace(os.Args[2]))
	if err != nil {
		log.Println("fail connect for tv", os.Args[1], "[", os.Args[2], "]", "error:", err)
		return
	}
	if err = sony.Do(strings.TrimSpace(os.Args[3])); err != nil {
		log.Println("fail command for tv", os.Args[1], "[", os.Args[2], "]", "(", os.Args[3], ")", "error:", err)
		return
	}
	fmt.Println("command", os.Args[3], "successfully")
}
