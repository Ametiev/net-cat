package cmd

import (
	"fmt"
	"time"
)

func Broadcast() {
	for {
		time.Sleep(5 * time.Minute)
		msg := fmt.Sprintf("[%s] - %d users online\n", GetFormatTime(), len(Users))
		mutex.Lock()
		for _, u := range Users {
			fmt.Fprintln(u.Conn, "\n\033[1A"+"\033[K"+msg)
		}
		mutex.Unlock()
	}
}
