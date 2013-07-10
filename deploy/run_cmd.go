package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
    "strings"
)

type AppInstance struct {
    Cmd string
    Args []string
    PipeIn chan string
    PipeOut chan string
}

func NewAppInstance() *AppInstance {
    return &AppInstance{PipeIn: make(chan string), PipeOut: make(chan string), Args: new([]string)}
}


func main() {
    // node command to run
    cmd := exec.Command("node", "test_io.js")
    //send input to it via stdin
    cmd.Stdin = strings.NewReader("some input")
    // receive the output via stdou
    var out bytes.Buffer
    cmd.Stdout = &out

    //run command
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    //print the command output to the console
    fmt.Printf("Node test output: %q\n", out.String())
}
