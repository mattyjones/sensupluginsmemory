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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yieldbot/sensuplugin/sensuutil"
)

var warnThreshold int
var critThreshold int
var checkKey string
var meminfo = "/proc/meminfo"

// checkMemoryInfoCmd represents the checkMemoryInfo command
var checkMemoryInfoCmd = &cobra.Command{
	Use:   "checkMemoryInfo",
	Short: "Check against any value in /proc/meminfo",
	Long:  `This load /proc/meminfo into a map and allows a user to pass in a key and a warn/crit value to compare against`,
	Run: func(cmd *cobra.Command, args []string) {

		data := createMap(meminfo)

		if debug {
			for k, v := range data {
				fmt.Println("Key: ", k, "    ", "Current value: ", v)
			}
			sensuutil.Exit("Debug")
		}

		if overThreshold(data[checkKey], int64(critThreshold)) {
			fmt.Printf("%v is over the critical threshold of %v", checkKey, critThreshold)
			sensuutil.Exit("critical")
		} else if overThreshold(data[checkKey], int64(warnThreshold)) {
			fmt.Printf("%v is over the warning threshold of %v", checkKey, warnThreshold)
			sensuutil.Exit("warning")
		} else {
			sensuutil.Exit("ok")
		}
	},
}

func init() {
	RootCmd.AddCommand(checkMemoryInfoCmd)

	checkMemoryInfoCmd.Flags().IntVarP(&warnThreshold, "warn", "", 100000, "the alert warning threshold")
	checkMemoryInfoCmd.Flags().IntVarP(&critThreshold, "crit", "", 200000, "the alert critical threshold")
	checkMemoryInfoCmd.Flags().StringVarP(&checkKey, "checkKey", "", "MemFree", "the key to check")
}
