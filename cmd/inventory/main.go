package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go-tangra-inventory/internal/collector"
)

func main() {
	outputDir := flag.String("o", "", "directory path to save inventory JSON (filename: HOSTNAME-DATE-TIME.json)")
	flag.Parse()

	inv, err := collector.Collect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: %v\n", err)
	}

	var w *os.File
	var outputPath string
	if *outputDir != "" {
		if err := os.MkdirAll(*outputDir, 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot create output directory: %v\n", err)
			os.Exit(1)
		}

		hostname := inv.Hostname
		if hostname == "" {
			hostname = "unknown"
		}
		hostname = strings.ReplaceAll(hostname, string(os.PathSeparator), "_")
		timestamp := time.Now().Format("20060102-150405")
		filename := fmt.Sprintf("%s-%s.json", hostname, timestamp)
		outputPath = filepath.Join(*outputDir, filename)

		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: cannot create output file: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(inv); err != nil {
		fmt.Fprintf(os.Stderr, "error: encoding inventory: %v\n", err)
		os.Exit(1)
	}

	if outputPath != "" {
		fmt.Fprintf(os.Stderr, "inventory written to %s\n", outputPath)
	}
}
