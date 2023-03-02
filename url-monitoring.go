package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoringLoop = 2
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()
		input := readInput()

		switch input {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid input")
			os.Exit(-1)
		}
	}

}

func showIntroduction() {
	version := 1.2
	fmt.Println("Hello!")
	fmt.Println("This software is on version", version)
}

func showMenu() {
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Quit")
	fmt.Println()
}

func readInput() int {
	var input int
	fmt.Scan(&input)
	fmt.Println("Value of the input:", input)
	fmt.Println()

	return input
}

func startMonitoring() {

	urls := readUrlsFromFile()

	for i := 0; i < monitoringLoop; i++ {
		for i, url := range urls {
			fmt.Println("Monitoring site", i, ":", url)
			verifyUrl(url)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

	fmt.Println()
}

func verifyUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("URL:", url, "had an Error when verifying ->", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("URL:", url, "had a successful return")
		createLog(url, true)
	} else {
		fmt.Println("URL:", url, "returned Status Code:", resp.StatusCode)
		createLog(url, false)
	}
}

func readUrlsFromFile() []string {
	var urls []string
	file, err := os.Open("urls.txt")

	if err != nil {
		fmt.Println("Error when opening file:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		urls = append(urls, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return urls
}

func createLog(url string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
