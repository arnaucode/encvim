package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) <= 1 {
		color.Red("not enough arguments")
		os.Exit(0)
	}

	filename := os.Args[1]

	_, err := readFile(filename)
	var out []byte
	if err == nil {
		//file exists
		//open the current existing file, decrypts it, and save it to tempfile
		//decrypt file
		out, err = exec.Command("bash", "-c", "openssl aes-256-cbc -d -a -in "+filename+" -out tempfile").Output()
		if err != nil {
			cmd := exec.Command("bash", "-c", "rm tempfile")
			cmd.Run()
			log.Fatal(err)
		}
		fmt.Println(out)
	}
	//open or create a tempfile to edit
	cmd := exec.Command("bash", "-c", "vim tempfile")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()

	if !strings.Contains(filename, ".encvim") {
		filename = filename + ".encvim"
	}
	//encrypt the file
	out, err = exec.Command("bash", "-c", "openssl aes-256-cbc -a -salt -in tempfile -out "+filename).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	//delete tempfile
	cmd = exec.Command("bash", "-c", "rm tempfile")
	cmd.Run()

	fmt.Println(filename)

}
func readFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	return string(b), err
}
