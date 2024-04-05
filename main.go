package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"unsafe"
)

type Data struct {
	Min   float64
	Sum   float64
	Count int
	Max   float64
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		buf := scanner.Bytes()
		ix := bytes.IndexByte(buf, ';')
		cityUnsafe := unsafe.String(&buf[0], len(buf[:ix]))
		temp := parseFloatFast(buf[ix:])

		data, ok := result[cityUnsafe]
		if !ok {
			city := string(buf[:ix])
			result[city] = &Data{
				Min:   temp,
				Sum:   temp,
				Count: 1,
				Max:   temp,
			}
		} else {
			data.Count++
			data.Sum += temp

			if data.Min > temp {
				data.Min = temp
			}

			if data.Max < temp {
				data.Max = temp
			}
		}
	}

	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	wr := bufio.NewWriter(os.Stdout)
	for _, key := range keys {
		data := result[key]
		mean := data.Sum / float64(data.Count)
		wr.WriteString(fmt.Sprintf("%s;%.1f;%.1f;%.1f\n", key, data.Min, mean, data.Max))
	}
	wr.Flush()
}

func parseFloatFast(bs []byte) float64 {
	var intStartIdx int // is negative?
	if bs[0] == '-' {
		intStartIdx = 1
	}

	v := float64(bs[len(bs)-1]-'0') / 10 // single decimal digit
	place := 1.0
	for i := len(bs) - 3; i >= intStartIdx; i-- { // integer part
		v += float64(bs[i]-'0') * place
		place *= 10
	}

	if intStartIdx == 1 {
		v *= -1
	}
	return v
}
