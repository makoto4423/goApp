package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	m := make(map[string]interface{})
	scan("E:/thunder2", m)
	var s string
	for k := range m {
		s = s + k + "\n"
	}
	fmt.Println(s)
}

func scan(path string, m map[string]interface{}) bool {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	last := false
	var b bool
	for _, d := range dir {
		if d.IsDir() {
			if d.Name() != "baidu" && d.Name() != "System Volume Information" ||
				d.Name() != "$RECYCLE.BIN" {
				if d.Name() != "thunder3" && d.Name() != "thunder4" && d.Name() != "thunder2" &&
					d.Name() != "thunder5" && d.Name() != "Thunder" {
					last = true
				}
				b = scan(path+"/"+d.Name(), m)
			}
		}
	}
	if b {
		for _, d := range dir {
			if !d.IsDir() {
				name := d.Name()
				arr := strings.Split(name, ".")
				name = arr[len(arr)-1]
				//if name == "mka" || name == "apk" || name == "xltd" ||
				//	name == "zip" || name == "ini" {
				//	fmt.Println(path + "/" + d.Name())
				//}
				if name != "torrent" {
					m[d.Name()] = nil
				}
			}
		}
	} else {
		if path != "E:/baidu" && path != "E:/System Volume Information" ||
			path != "E:/$RECYCLE.BIN" {
			m[path] = nil
		}
	}
	return last
}
