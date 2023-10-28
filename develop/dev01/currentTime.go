package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	address := "0.beevik-ntp.pool.ntp.org"
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	currentTime, err := currentTime(address)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	printTime(currentTime, writer)
}

func currentTime(address string) (currentTime time.Time, err error) {
	currentTime, err = ntp.Time(address)
	return
}

func printTime(currentTime time.Time, writer *bufio.Writer) {
	fmt.Fprintf(writer, "%d-%02d-%02d %02d:%02d:%02d\n",
		currentTime.Month(), currentTime.Day(),
		currentTime.Year(), currentTime.Hour(),
		currentTime.Minute(), currentTime.Second())
}
