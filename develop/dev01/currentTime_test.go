package main

import (
	"bufio"
	"bytes"
	"testing"
	"time"
)

func TestValid(t *testing.T) {
	address := "0.beevik-ntp.pool.ntp.org"
	_, err := currentTime(address)
	if err != nil {
		t.Fatal("Error where it was not expected:", err)
	}
}

func TestInvalid(t *testing.T) {
	address := "invalid.address.dcdefg"
	_, err := currentTime(address)
	if err == nil {
		t.Fatal("No error where one is expected")
	}
}

func TestOutput(t *testing.T) {
	date := time.Date(2023, time.October, 28, 23, 59, 59, 0, time.UTC)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	printTime(date, writer)
	writer.Flush()
	if buffer.String() != "10-28-2023 23:59:59\n" {
		t.Fatal("Wrong output format")
	}
}
