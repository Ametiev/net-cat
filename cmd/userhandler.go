package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func init() {
	Users = make(map[string]*Client)
}

func (User *Client) Handle() {
	User.Greet()
	User.Register()
	defer User.Disconnect()
	scanner := bufio.NewScanner(User.Conn)
	for scanner.Scan() {
		fmt.Fprint(User.Conn, "\033[2K\r")
		fmt.Fprint(User.Conn, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"["+User.Name+"]:")
		text := scanner.Text()
		if !CheckMsg(text) {
			msg := fmt.Sprintf("[%s][%s]:%s", time.Now().Format("2006-01-02 15:04:05"), User.Name, text)
			allMessages += msg + "\n"
			User.Message <- msg
			User.Messages = append(User.Messages, msg)
		} else {
			fmt.Fprint(User.Conn, "\033[2K\r")
			io.WriteString(User.Conn, "Message contains invalid characters, please choose another\n")
			fmt.Fprint(User.Conn, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"["+User.Name+"]:")
			continue
		}
	}
}

func (User *Client) Greet() {
	txt, errTxt := ioutil.ReadFile("greetings.txt")
	if errTxt != nil {
		os.Exit(0)
	}
	io.WriteString(User.Conn, string(txt)+"\n")
}
func (user *Client) Register() {
	for {
		io.WriteString(user.Conn, "[ENTER YOUR NAME]: ")
		scanner := bufio.NewScanner(user.Conn)
		scanner.Scan()
		name := strings.TrimSpace(scanner.Text())
		if name == "" {
			continue
		}
		if !isPrintable(name) {
			io.WriteString(user.Conn, "Name contains invalid characters, please choose another\n")
			continue
		}
		mutex.Lock()
		if _, ok := Users[name]; ok {
			mutex.Unlock()
			io.WriteString(user.Conn, "Name already taken, please choose another\n")
			continue
		}
		msg := fmt.Sprintf("%s has joined the chat...", name)
		allMessages += msg + "\n"
		user.Name = name
		user.Message <- msg
		Users[user.Name] = user
		mutex.Unlock()
		wlm := "Welcome to the chat, " + user.Name + "!\n"
		io.WriteString(user.Conn, wlm)
		if len(allMessages) > 0 {
			io.WriteString(user.Conn, allMessages)
		}
		for _, u := range Users {
			if u != user {
				fmt.Fprint(u.Conn, "\033[2K\r")
				fmt.Fprint(u.Conn, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"["+u.Name+"]:")
			}
		}
		fmt.Fprint(user.Conn, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"["+user.Name+"]:")
		break
	}
}
func (user *Client) Disconnect() {
	mutex.Lock()
	delete(Users, user.Name)
	mutex.Unlock()
	user.Conn.Close()
	msg := fmt.Sprintf("%s has left the chat...", user.Name)
	allMessages += msg + "\n"
	user.Message <- msg
}

func (user *Client) ReadMessages() {
	for {
		select {
		case msg := <-user.Message:
			mutex.Lock()
			for _, u := range Users {
				if u != user {
					fmt.Fprintln(u.Conn, "\n\033[1A"+"\033[K"+msg)
					fmt.Fprint(u.Conn, "["+time.Now().Format("2006-01-02 15:04:05")+"]"+"["+u.Name+"]:")
				}
			}
			mutex.Unlock()
		}
	}
}
