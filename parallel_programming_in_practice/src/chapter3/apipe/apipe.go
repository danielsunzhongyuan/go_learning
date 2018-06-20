package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	runCmd()
	fmt.Println()
	runCmdWithPipe()
	demo2()
}

func demo2() {
	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "au")
	//cmd1 := exec.Command("ps", "x") // 不能同时 aux，要么 au， 要么 x，否则会在 cmd1 那里卡住
	// 而且直接在 IDE 里运行也是会出错的。 go run apipe.go 或者 go build apipe.go && ./apipe 都是可以的
	cmd2 := exec.Command("grep", "apipe")
	stdout1, err := cmd1.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe for command: %s", err)
		return
	}
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The command can not running: %s\n", err)
		return
	}
	outputBuf1 := bufio.NewReader(stdout1)
	stdin2, err := cmd2.StdinPipe()
	if err != nil {
		fmt.Printf("Error: Can not obtain the stdin pipe for command: %s\n", err)
		return
	}
	outputBuf1.WriteTo(stdin2)
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}
	err = stdin2.Close()
	if err != nil {
		fmt.Printf("Error: Can not close the stdio pipe: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Can not wait for the command: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}
func runCmdWithPipe() {
	fmt.Println("Run command `ps aux | grep apipe`: ")
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("Error: The first command can not be startup: %s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the first command: %s\n", err)
		return
	}
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("Error: The second command can not be startup: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("Error: Couldn't wait for the second command: %s\n", err)
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}

func runCmd() {
	useBufferedIO := false
	fmt.Println("Run command `echo -n \"My first command comes from golang.\"`: ")
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
	stdout0, err := cmd0.StdoutPipe()
	if err != nil {
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0: %s\n", err)
		return
	}
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
	}
	if !useBufferedIO {
		var outputBuf0 bytes.Buffer
		for {
			tempOutput := make([]byte, 5)
			n, err := stdout0.Read(tempOutput)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempOutput[:n])
			}
		}
		fmt.Printf("%s\n", outputBuf0.String())
	} else {
		outputBuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputBuf0.ReadLine()
		if err != nil {
			fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
			return
		}
		fmt.Printf("%s\n", string(output0))
	}
}
