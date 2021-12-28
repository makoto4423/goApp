package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Cmd struct {
	Version string   `json:"version"`
	Command []string `json:"command"`
	File    string   `json:"file"`
	Project string   `json:"project"`
}

func main() {
	bytes, err := os.ReadFile("./config.json")
	if err != nil {
		log.Default().Fatalln(err)
	}
	m := make(map[string]Cmd)
	json.Unmarshal(bytes, &m)
	args := os.Args[1:]
	command := m[args[0]]
	for _, c := range command.Command {
		c = strings.ReplaceAll(c, "${version}", command.Version)
		c = strings.ReplaceAll(c, "${project}", command.Project)
		log.Default().Println(c)
		arr := strings.Split(c, " ")
		cmd := exec.Command(arr[0], arr[1:]...)
		bytes, err := cmd.Output()
		if err != nil {
			log.Default().Println(string(bytes))
			log.Default().Fatalln("cmd is ", c, "; error is ", err)
		}
	}
}

// args... 数组和可变参数通过...相互转换
func Test(args ...string) {
	arr := make([]string, 0)
	arr = append(arr, args...)
	fmt.Println(arr)
}
