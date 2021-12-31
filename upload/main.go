package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	UpLoadUrl   string `json:"upLoadUrl"`
	DelUrl      string `json:"delUrl"`
	Directory   string `json:"directory"`
	Repo        string `json:"repo"`
	Sfcsrftoken string `json:"sfcsrftoken"`
	Sessionid   string `json:"sessionid"`
	File        string `json:"file"`
}

func main() {
	bytes, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	config := Config{}
	json.Unmarshal(bytes, &config)
	sfcsrftoken := config.Sfcsrftoken
	sessionid := config.Sessionid
	cookie := "sfcsrftoken=" + sfcsrftoken + "; sessionid=" + sessionid + "; group_expanded=true;"
	delUrl := config.DelUrl
	file := config.File
	fileName := file[strings.LastIndex(file, "/")+1:]
	delUrl = strings.ReplaceAll(delUrl, "${repo}", config.Repo)
	delUrl = strings.ReplaceAll(delUrl, "${directory}", config.Directory)
	deleteOld(delUrl, cookie, sfcsrftoken, fileName)

	uploadUrl := config.UpLoadUrl
	uploadUrl = strings.ReplaceAll(uploadUrl, "${repo}", config.Repo)
	uploadUrl = strings.ReplaceAll(uploadUrl, "${directory}", config.Directory)
	uploadUrl = strings.ReplaceAll(uploadUrl, "${time}", strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	url, err := getUploadUrl(uploadUrl, cookie)
	if err != nil {
		fmt.Println(err)
		return
	}
	url = url[1 : len(url)-1]
	err = upload(url, cookie, file, config.Directory)
	if err != nil {
		fmt.Println(err)
	}

}

func getUploadUrl(from, cookie string) (string, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, from, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func upload(url, cookie, fileName, directory string) error {
	method := "POST"
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	file.Name()
	defer file.Close()
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("parent_dir", directory)
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func deleteOld(url, cookie, sfcsrftoken, fileName string) error {
	method := "DELETE"
	url = url + "/" + fileName
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Add("X-CSRFToken", sfcsrftoken)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
