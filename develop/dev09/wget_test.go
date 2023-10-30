package main

import (
	"math"
	"os"
	"os/exec"
	"sync"
	"testing"
)

func TestNonRecursive(t *testing.T) {
	target := "google.com"
	command := exec.Command("wget", target)
	command.Run()
	original, err := os.ReadFile("index.html")
	os.Remove("index.html")
	if err != nil {
		t.Fatal(`Unexpected wget error`)
	}

	targets := []string{target}
	nonRecursiveWget(targets)
	myWget, err := os.ReadFile("index.html")
	os.Remove("index.html")
	if err != nil {
		t.Fatal(`myWget did not produce an index.html file`)
	}
	if math.Abs(float64(len(original)-len(myWget))) > 100 {
		t.Fatal(`index.html doesn't match the one produced by the original wget`)
	}
}

func TestRecursive(t *testing.T) {
	target := "website.org"
	command := exec.Command("wget", target, "-r")
	command.Run()
	original, err := os.ReadDir(target)
	os.RemoveAll(target)
	if err != nil {
		t.Fatal(`Unexpected wget error`)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	recursiveWget(target, wg)
	myWget, err := os.ReadDir(target)
	os.RemoveAll(target)
	if err != nil {
		t.Fatal(`myWget did not produce an a directory with website contents`)
	}
	if math.Abs(float64(len(original)-len(myWget))) > 10 {
		t.Fatal(`Too much discrepancy between website contents`)
	}
}
