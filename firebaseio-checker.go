/*
 * Copyright @random_robbie (c) 2018.
 */

package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/remeh/sizedwaitgroup"
)

var (
	fileofurls  = "list.txt"
	outputfile  = "./cfg/"
	filepathurl = "/.json"
	au          aurora.Aurora
	colors      = flag.Bool("colors", true, "enable or disable colors")
)

func init() {
	flag.Parse()
	au = aurora.NewAurora(*colors)
}

func grabURL(URL string, output string, filepathurl string, swg *sizedwaitgroup.SizedWaitGroup) {

	defer swg.Done()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	newurl := "https://" + URL + filepathurl
	resp, err := client.Get(newurl)
	if err != nil {
		// handle err
		log.Fatalf("could not send request: %v", err)
	}

	resp.Header.Set("Accept-Encoding", "gzip, deflate")
	resp.Header.Set("Accept", "*/*")
	resp.Header.Set("Accept-Language", "en")
	resp.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0")
	fmt.Println("[*] Testing ", newurl, "[*]")

	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println(au.Red("[*] Database Not Found [*]"))
	}
	if resp.StatusCode == 403 {
		fmt.Println(au.Red("[*] Access Denied [*]"))
	}
	if resp.StatusCode == 401 {
		fmt.Println(au.Red("[*] Access Denied [*]"))
	}

	if resp.StatusCode == 200 {
		if resp.Body != nil {

			removefilebase := ""
			removefilebase = strings.Replace(URL, ".firebaseio.com", "", 1)
			capturedfile := outputfile + removefilebase + ".json"

			htmlfile, err := os.Create(capturedfile)

			if err != nil {
				log.Fatalf("could not create file: %v", err)
				os.Exit(1)
			}

			defer htmlfile.Close()
			fmt.Println(au.Green("[*] Saving Database to file [*]"))
			io.Copy(htmlfile, resp.Body)

		}

	}

}

func main() {
	swg := sizedwaitgroup.New(10)
	fmt.Println(au.Blue("[*] Firebase IO Checker - By @random_robbie [*]"))
	file, err := os.Open(fileofurls)
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		swg.Add()
		URL := strings.TrimSpace(scanner.Text())
		go grabURL(URL, outputfile, filepathurl, &swg)
		swg.Wait()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
