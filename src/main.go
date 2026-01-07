package main

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"syscall"
	"golang.org/x/sys/windows"
)

func lemmeknowsilent() bool {
	for _, arg := range os.Args {
		if arg == "--silent" {
			return true
		}
	}
	return false
}


func isadmin() bool {
	sid, _ := windows.CreateWellKnownSid(windows.WinBuiltinAdministratorsSid) // idk well what it creates
	token := windows.Token(0)
	ok, _ := token.IsMember(sid)
	return ok
}

func run(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.SysProcAttr = &syscall.SysProcAttr{ 
		HideWindow: true,
	}
	_ = c.Run()
}

// why did i wrote this

func main() {
	silent := lemmeknowsilent()

	if !silent {
		banner()
		fmt.Println("IP changing utility for Windows, written in Go.")
		fmt.Println("Made by https://github.com/sowmething")
	}

	if !isadmin() && !silent {
		fmt.Println("\x1b[33m[!] Admin recommended.\x1b[0m")
	}

	if !silent {
		fmt.Println("\x1b[32m[+] Releasing IP. Disconnect!\x1b[0m")
	}
	run("ipconfig", "/release")

	time.Sleep(1500 * time.Millisecond)

	if !silent {
		fmt.Println("\x1b[32m[+] Renewing IP. Reconnect!\x1b[0m")
	}
	run("ipconfig", "/renew")

	if !silent {
		fmt.Println("\x1b[32m[+] Complete.\x1b[0m")
	}

	if silent {
		fmt.Println("Operation completed successfully.")
	}

	os.Exit(0)
}