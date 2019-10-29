package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main(){

 	asd := `package main;import "fmt" ;func main() {fmt.Print("hello")}`
	err := ioutil.WriteFile("test.go", []byte(asd), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	run("go", "run", "test.go")
	run("pwd")
	defer timeTrack(time.Now(), "run")
	run("python3.6", "./python/test.py")

	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

