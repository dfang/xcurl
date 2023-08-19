package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// https://go.dev/play/p/BV2GnfSiH1R
func execShell(cmd string, args []string) string {
	log.Println(cmd + " " + strings.Join(args, " "))
	var command = exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	var err = command.Start()
	if err != nil {
		return err.Error()
	}
	err = command.Wait()
	if err != nil {
		return err.Error()
	}
	return ""
}

func main() {
	log.Printf("xcurl\n")
	args := os.Args
	rawURL, prefixedURL := "", ""
	mirror := os.Getenv("GITHUB_MIRROR")
	if mirror == "" {
		mirror = "https://ghproxy.com"
	}

	for i := 0; i < len(args); i++ {
		if strings.Contains(args[i], "https://github.com") {
			rawURL = args[i]
			// args[i] = strings.Replace(args[i], "https://github.com", "https://ghproxy.com/https://github.com", -1)

			// mirrors.goproxyauth.com 只支持github.com这种格式的URL，即用来下载github release的URL
			prefixedURL = fmt.Sprintf("%s/%s", randomMirror(), rawURL)
		}

		if strings.Contains(args[i], "https://raw.githubusercontent.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		if strings.Contains(args[i], "https://gist.github.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		if strings.Contains(args[i], "https://gist.githubusercontent.com") {
			rawURL = args[i]
			prefixedURL = fmt.Sprintf("%s/%s", mirror, rawURL)
		}

		args[i] = prefixedURL
	}

	flag.Parse()

	// extract the file name from the URL path
	parsedURL, err := url.Parse(prefixedURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fileName := filepath.Base(parsedURL.Path)
	fmt.Println("File name:", fileName)

	// TOOD: extract filename from content-disposition header if filename is empty

	if flag.Parsed() && flag.Lookup("o") == nil {
		args = append(args, "-o")
		args = append(args, fileName)
	}

	execShell("curl", args[1:])
}

func randomMirror() string {
	// rand.Seed(time.Now().UnixNano())

	// Generate a secure random seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Use the secure random seed to initialize the random number generator
	r.Seed(r.Int63())

	slice := []string{"https://ghproxy.com", "https://mirrors.goproxyauth.com"}
	// Generate a random index
	randomIndex := rand.Intn(len(slice))
	// Access the random element from the slice
	randomElement := slice[randomIndex]
	return randomElement
}
