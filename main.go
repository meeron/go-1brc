package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	Min   float32
	Sum   float32
	Count int
	Max   float32
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-1brc <input_file>")
		return
	}

	filePath := os.Args[1]
	fmt.Println(filePath)

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	result := make(map[string]*Data)
	keys := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ";")
		city := parts[0]
		temp, _ := strconv.ParseFloat(parts[1], 32)
		temp32 := float32(temp)

		data, ok := result[city]
		if !ok {
			result[city] = &Data{
				Min:   temp32,
				Sum:   temp32,
				Count: 1,
				Max:   temp32,
			}
			keys = append(keys, city)
		} else {
			data.Count++
			data.Sum += temp32

			if data.Min > temp32 {
				data.Min = temp32
			}

			if data.Max < temp32 {
				data.Max = temp32
			}
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		data := result[key]
		mean := data.Sum / float32(data.Count)
		fmt.Printf("%s;%.1f;%.1f;%.1f\n", key, data.Min, mean, data.Max)
	}
}
