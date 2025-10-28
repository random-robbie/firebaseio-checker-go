/*
 * Firebase Database Security Checker
 * Copyright @random_robbie (c) 2018
 *
 * A security audit tool to check Firebase Realtime Database instances
 * for publicly accessible data and misconfigured security rules.
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
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/remeh/sizedwaitgroup"
)

var (
	inputFile   = flag.String("input", "list.txt", "Input file containing Firebase URLs")
	outputDir   = flag.String("output", "./cfg/", "Output directory for JSON files")
	concurrency = flag.Int("concurrency", 10, "Number of concurrent requests")
	timeout     = flag.Int("timeout", 30, "HTTP request timeout in seconds")
	colors      = flag.Bool("colors", true, "Enable or disable colored output")
	au          aurora.Aurora
)

func init() {
	flag.Parse()
	au = aurora.NewAurora(*colors)
}

// checkDatabase attempts to retrieve the Firebase database contents
func checkDatabase(url string, swg *sizedwaitgroup.SizedWaitGroup) {
	defer swg.Done()

	// Prepare HTTP client with custom transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(*timeout) * time.Second,
	}

	// Construct the full URL
	fullURL := fmt.Sprintf("https://%s/.json", strings.TrimSpace(url))

	// Create request with proper headers
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("%s Failed to create request for %s: %v\n", au.Red("[!]"), url, err)
		return
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:59.0) Gecko/20100101 Firefox/59.0")

	fmt.Printf("%s Testing %s\n", au.Cyan("[*]"), fullURL)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s Request failed for %s: %v\n", au.Red("[!]"), url, err)
		return
	}
	defer resp.Body.Close()

	// Handle response based on status code
	switch resp.StatusCode {
	case 200:
		fmt.Printf("%s Database accessible! Saving to file...\n", au.Green("[+]"))
		saveDatabase(url, resp.Body)
	case 401, 403:
		fmt.Printf("%s Access denied (permission rules configured)\n", au.Yellow("[-]"))
	case 404:
		fmt.Printf("%s Database not found\n", au.Red("[-]"))
	default:
		fmt.Printf("%s Unexpected status code: %d\n", au.Yellow("[!]"), resp.StatusCode)
	}
}

// saveDatabase writes the database contents to a JSON file
func saveDatabase(url string, body io.Reader) {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("%s Failed to create output directory: %v\n", au.Red("[!]"), err)
		return
	}

	// Generate filename from URL
	dbName := strings.Replace(url, ".firebaseio.com", "", 1)
	dbName = strings.TrimSpace(dbName)
	outputPath := fmt.Sprintf("%s%s.json", *outputDir, dbName)

	// Create output file
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("%s Failed to create output file %s: %v\n", au.Red("[!]"), outputPath, err)
		return
	}
	defer outFile.Close()

	// Copy response body to file
	bytesWritten, err := io.Copy(outFile, body)
	if err != nil {
		fmt.Printf("%s Failed to write data to file: %v\n", au.Red("[!]"), err)
		return
	}

	fmt.Printf("%s Successfully saved %d bytes to %s\n", au.Green("[+]"), bytesWritten, outputPath)
}

func main() {
	fmt.Println(au.Blue("===================================="))
	fmt.Println(au.Blue("  Firebase Database Security Checker"))
	fmt.Println(au.Blue("  By @random_robbie"))
	fmt.Println(au.Blue("===================================="))
	fmt.Println()

	// Open input file
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open input file %s: %v", *inputFile, err)
	}
	defer file.Close()

	// Create sized wait group for concurrency control
	swg := sizedwaitgroup.New(*concurrency)

	// Read URLs from file and process them
	scanner := bufio.NewScanner(file)
	urlCount := 0
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if url == "" || strings.HasPrefix(url, "#") {
			continue
		}

		urlCount++
		swg.Add()
		go checkDatabase(url, &swg)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Wait for all goroutines to complete
	swg.Wait()

	fmt.Println()
	fmt.Printf("%s Scan complete! Processed %d URLs\n", au.Green("[✓]"), urlCount)
}
