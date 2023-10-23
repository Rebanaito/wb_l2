package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

var recursive bool
var MAX_DEPTH int = 5

func main() {
	targets, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !recursive {
		nonRecursiveWget(targets)
	} else {
		wg := &sync.WaitGroup{}
		for _, target := range targets {
			wg.Add(1)
			go recursiveWget(target, wg)
		}
		wg.Wait()
	}
}

func parseArgs(argv []string) (targets []string, err error) {
	if len(argv) == 0 {
		return nil, fmt.Errorf("wget: missing URL")
	}
	for i, arg := range argv {
		switch arg {
		case "-r", "--recursive":
			recursive = true
		case "-l":
			if i == len(argv)-1 {
				return nil, fmt.Errorf("wget: option requires an argument -- 'l'")
			}
			num, err := strconv.Atoi(argv[i+1])
			if err != nil {
				return nil, err
			}
			MAX_DEPTH = num
			if num == 0 {
				MAX_DEPTH = 2147483647
			}
		default:
			if arg[0] == '-' {
				return nil, fmt.Errorf("wget: unrecognized option '%s'", arg)
			}
			if strings.Contains(arg, "http://") {
				arg = arg[7:]
			} else if strings.Contains(arg, "https://") {
				arg = arg[8:]
			}
			targets = append(targets, arg)
		}
	}
	return targets, nil
}

func nonRecursiveWget(targets []string) {
	for i, target := range targets {
		req, err := http.NewRequest(http.MethodGet, "https://"+target, nil)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		filename, err := parseFilename(target)
		if err != nil {
			return
		}
		if i != 0 {
			filename = filename + "." + strconv.Itoa(i)
		}
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func recursiveWget(target string, wg *sync.WaitGroup) {
	defer wg.Done()
	localWg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	linksDict := make(map[string]struct{})
	err := os.Mkdir(target, 0777)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	os.Chdir(target)
	localWg.Add(1)
	go recursiveRequest(target, 0, linksDict, localWg, mu)
	localWg.Wait()
	os.Chdir("../")
}

func recursiveRequest(target string, depth int, linksDict map[string]struct{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	req, err := http.NewRequest(http.MethodGet, "https://"+target, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	filename, err := parseFilename(target)
	if err != nil {
		return
	}
	err = os.WriteFile(filename, body, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if depth < MAX_DEPTH {
		subtargets := parseLinks(target, body, linksDict, mu)
		for _, subtarget := range subtargets {
			wg.Add(1)
			go recursiveRequest(subtarget, depth+1, linksDict, wg, mu)
		}
	}
}

func parseLinks(target string, body []byte, linksDict map[string]struct{}, mu *sync.Mutex) (subtargets []string) {
	wg := &sync.WaitGroup{}
	tokenizer := html.NewTokenizer(bytes.NewReader(body))
out:
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			break out
		case html.StartTagToken, html.EndTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" && linkNotDuplicate(attr.Val, linksDict, mu) {
						subtargets = append(subtargets, attr.Val)
						break
					}
				}
			} else if token.Data == "script" || token.Data == "link" {
				for _, attr := range token.Attr {
					if (attr.Key == "src" || attr.Key == "href") && linkNotDuplicate(attr.Val, linksDict, mu) {
						wg.Add(1)
						go saveFile(target, attr.Val, wg)
						break
					}
				}
			}
		}
	}
	wg.Wait()
	return subtargets
}

func linkNotDuplicate(link string, linksDict map[string]struct{}, mu *sync.Mutex) bool {
	mu.Lock()
	defer mu.Unlock()
	_, ok := linksDict[link]
	if ok {
		return false
	} else {
		linksDict[link] = struct{}{}
	}
	return true
}

func parseFilename(target string) (filename string, err error) {
	_, filename, found := strings.Cut(target, "/")
	if !found || filename == "" {
		return "index.html", nil
	}
	split := strings.Split(filename, "/")
	if len(split) > 1 {
		pwd, _ := os.Getwd()
		defer os.Chdir(pwd)
		for i := 0; i < len(split)-1; i++ {
			err := os.Mkdir(split[i], 0777)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return "", err
			}
			os.Chdir(split[i])
		}
	}
	return filename, nil
}

func saveFile(target, link string, wg *sync.WaitGroup) {
	defer wg.Done()
	if target[len(target)-1] == '/' {
		target += link
	} else {
		target = target + "/" + link
	}
	req, err := http.NewRequest(http.MethodGet, "https://"+target, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	parseFilename(target)
	err = os.WriteFile(link, body, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
