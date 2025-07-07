package main

import (
	"bufio"
	"fmt"
	"lru-cache/lru"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("go run main.go <filename> <capacity>")
		return
	}

	filename := os.Args[1]
	capacity, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("invalid capacity:", os.Args[2])
		return
	}

	cache := lru.New(capacity)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tokens := strings.Fields(line)
		cmd := strings.ToUpper(tokens[0])

		switch cmd {
		case "GET":
			if len(tokens) != 2 {
				fmt.Println("Invalid PUT line:", line)
				continue
			}

			key, _ := strconv.Atoi(tokens[1])
			if val, ok := cache.Get(key); ok {
				fmt.Printf("GET %d -> %d\n", key, val)
			} else {
				fmt.Printf("GET %d -> -1\n", key)
			}
		case "PUT":
			if len(tokens) != 3 {
				fmt.Println("Invalid GET line:", line)
				continue
			}
			key, _ := strconv.Atoi(tokens[1])
			value, _ := strconv.Atoi(tokens[2])
			cache.Put(key, value)

		default:
			fmt.Println("Error command", cmd)

		}

	}
}
