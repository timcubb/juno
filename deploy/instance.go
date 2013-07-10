package deploy

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
    PipeIn chan AppInput
    PipeOut chan AppOutput
}

type AppInput struct {
    
}

type AppOutput struct {

}

func NewAppInstance() *AppInstance {
    return &AppInstance{PipeIn: make(chan AppInput), PipeOut: make(chan AppOutput), Args: new([]string)}
}

func (i *AppInstance) Install() {
    //install app on remote node
}

func (i *AppInstance) Run() {
    //install app on remote node
}

func (i *AppInstance) SendData(data *AppInput) {
    //install app on remote node
}

func (i *AppInstance) RecvData(data *AppOutput) {
    //install app on remote node
}

func (i *AppInstance) Completed() {
    //install app on remote node
}

func (i *AppInstance) Crashed() {
    //install app on remote node
}

func main() {

    ai := deploy.NewAppInstance()

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
