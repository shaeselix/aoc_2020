package utils

import (
	"net/http"
	"os"
	"io"
	"fmt"
)

func DownloadFile(url string, filename string, token string) {
	if fileExists(filename) {
		fmt.Printf("%v already exists!", filename)
		return
	}
	cookie := http.Cookie{Name: "session", Value: token}
	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	Check(err)
	req.AddCookie(&cookie)
	resp, err := client.Do(req)
	Check(err)
	defer resp.Body.Close()
	writeFile(resp.Body, filename)
	fmt.Printf("%v downloaded", filename)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func writeFile(body io.ReadCloser, filename string) {
	out, err := os.Create(filename)
	Check(err)
	defer out.Close()
	_, err = io.Copy(out, body)
	Check(err)
}
