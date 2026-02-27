// Copyright (c) Sidero Labs, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"
	"syscall"
)

func utsFieldToString(f [65]int8) string {
	var buf []byte
	for _, b := range f {
		if b == 0 {
			break
		}
		buf = append(buf, byte(b))
	}
	return string(buf)
}

func main() {
	var utsname syscall.Utsname
	if err := syscall.Uname(&utsname); err != nil {
		fmt.Fprintf(os.Stderr, "uname: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s %s %s %s\n",
		utsFieldToString(utsname.Sysname),
		utsFieldToString(utsname.Release),
		utsFieldToString(utsname.Version),
		utsFieldToString(utsname.Machine),
	)
}
