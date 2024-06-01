package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	url := flag.String("url", "", "URL of the file to download")
	output := flag.String("output", "", "Output file path")
	flag.Parse()

	if *url == "" || *output == "" {
		fmt.Println("Usage: go run main.go -url <file_url> -output <output_path>")
		return
	}

	err := downloadFile(*url, *output)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadFile(url, output string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to initiate request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	outFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	progress := &Progress{Total: int64(size)}

	_, err = io.Copy(outFile, io.TeeReader(resp.Body, progress))
	if err != nil {
		return fmt.Errorf("failed to download file: %v", err)
	}

	fmt.Println("\nDownload completed!")
	return nil
}

type Progress struct {
	Total      int64
	Downloaded int64
	LastUpdate time.Time
}

func (p *Progress) Write(data []byte) (int, error) {
	n := len(data)
	p.Downloaded += int64(n)

	now := time.Now()
	if now.Sub(p.LastUpdate) >= 100*time.Millisecond {
		p.LastUpdate = now
		p.printProgress()
	}

	return n, nil
}

func (p *Progress) printProgress() {
	percentage := float64(p.Downloaded) / float64(p.Total) * 100
	progressBar := strings.Repeat("=", int(percentage/2)) + strings.Repeat(" ", 50-int(percentage/2))
	fmt.Printf("\r[%s] %.2f%%", progressBar, percentage)
}
