package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var (
	delay *string = flag.String("d", "500ms", "Line processing delay in following format <decimal><time-unit>. Valid time-units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.")
	usage         = func() {
		fmt.Fprintln(os.Stderr, "Usage: lt [-d <delay>ns|us|µs|ms|s|m|h")
		flag.PrintDefaults()
		os.Exit(2)
	}
)

func main() {
	flag.Usage = usage
	flag.Parse()

	dur, err := time.ParseDuration(*delay)
	if err != nil {
		log.Println(err)
		usage()
	}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	for {
		data, err := in.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else if strings.TrimSpace(string(data)) == "" {
			continue
		} else {
			out.Write(data)
			out.Flush()
		}
		// Throttle output by delaying next processing step
		time.Sleep(dur)
	}
}
