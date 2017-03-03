package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	d := make([][]string, 100)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		s := scanner.Text()
		if i < n/2 {
			s = "-"
		}
		d[x] = append(d[x], s)
	}
	io := bufio.NewWriter(os.Stdout)
	for _, r := range d {
		for _, s := range r {
			io.WriteString(s)
			io.WriteString(" ")
		}
	}
	io.Flush()
}
