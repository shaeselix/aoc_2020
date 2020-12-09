package utils

import (
	"net/http"
	"os"
	"io"
	"fmt"
)

const (
	aocTokenEnvVar = "AOC_SESSION"
	chromeSessionPath = "chrome://settings/cookies/detail?site=adventofcode.com"
	dataFileFormat = "%d.txt"
	urlFormat = "https://adventofcode.com/2020/day/%d/input"
)

func DownloadFile(day int, dataDir string, verbose bool) {
	filename := GetFileName(day, dataDir)
	if fileExists(filename) {
		if verbose { fmt.Printf("%v already exists!\n", filename) }
		return
	}
	token := getToken()
	cookie := http.Cookie{Name: "session", Value: token}
	client := new(http.Client)
	url := getURL(day)
	req, err := http.NewRequest("GET", url, nil)
	Check(err)
	req.AddCookie(&cookie)
	resp, err := client.Do(req)
	Check(err)
	defer resp.Body.Close()
	writeFile(resp.Body, filename)
	if verbose { fmt.Printf("%v downloaded\n", filename) }
}

func GetFileName(day int, dataDir string) string {
	file := fmt.Sprintf(dataFileFormat, day)
    return fmt.Sprintf("%s%s", dataDir, file)
}

func getURL(day int) string {
	return fmt.Sprintf(urlFormat, day)
}

func getToken() string {
	token, ok := os.LookupEnv(aocTokenEnvVar)
	if !ok {
		panic(
			fmt.Sprintf("%s is unset. Visit %s and set env var.",
			aocTokenEnvVar,
			chromeSessionPath,
		))
	}
	return token
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
