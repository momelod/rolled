package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("open", "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	fmt.Println("You've just been rolled!")
	cmd.Start()
}
