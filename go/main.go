package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func run(command string, args ...string){
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main(){
 asd := `package main;import "fmt" ;func main() {fmt.Print("hello")}`
	err := ioutil.WriteFile("test.go", []byte(asd), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	run("go", "run", "test.go")
	run("pwd")
	run("python3.6", "../python/main.py")

}

