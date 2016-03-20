// Copyright © 2016 Yieldbot <devops@yieldbot.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()

}

func createMap(meminfo string) map[string]int64 {
	m := make(map[string]int64)
	var key string
	var val int64
	lines, err := readLines(meminfo)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	reSpace := regexp.MustCompile(`[\s]+`)
	reNum := regexp.MustCompile(`[0-9]+`)

	for _, line := range lines {
		l := strings.Split(line, ":")

		for i := range l {
			if i == 0 {
				key = l[i]
			} else {
				r := reSpace.Split(l[i], -1)
				for _, n := range r {
					if val, err = strconv.ParseInt(reNum.FindString(n), 10, 32); err == nil {
						m[key] = val
					}
				}
			}
		}
	}
	return m
}

func overThreshold(curVal int64, threshold int64) bool {
	if curVal > threshold {
		return true
	}
	return false
}
