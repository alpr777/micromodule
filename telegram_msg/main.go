package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Print(
			"====================================================", "\n",
			"./[app name] [bot token] [chat id] [msg]", "\n",
			"----------------------------------------------------", "\n",
			"EXAMPLE:", "\n",
			"./[app name] 123456789:zxfsdfsadsdfad-dd -987654321 test message ", "\n",
			"----------------------------------------------------", "\n",
			"SPECIAL SYMBOLS:", "\n",
			"next line: /n", "\n",
			"====================================================", "\n",
		)
		return
	}
	msg := strings.ReplaceAll(fmt.Sprintln(os.Args[3:]), "/n", "%0A")
	resp, err := http.Get("https://api.telegram.org/bot" + strings.TrimSpace(os.Args[1]) + "/sendMessage?chat_id=" + strings.TrimSpace(os.Args[2]) + "&text=" + msg)
	fmt.Print("resp error:", err, "close error:", resp.Body.Close())
}
