package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var s string
	var err error
	if len(os.Args) <= 1 {
		s, err = readFromStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	} else {
		s = strings.Join(os.Args[1:], " ")
	}

	s, err = convert(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	io.WriteString(os.Stdout, s)
	fmt.Fprintln(os.Stdout)
}

func readFromStdin() (string, error) {
	var lines []string
	for {
		var s string
		n, err := fmt.Scan(&s)
		if n > 0 {
			lines = append(lines, s)
		}
		if err != nil {
			if err == io.EOF {
				return strings.Join(lines, "\n"), nil
			}
			return "", err
		}
	}
	return strings.Join(lines, "\n"), nil
}

func splitWellFormedString(s string) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		if r == ',' || r == ';' {
			return true
		}
		if unicode.IsSpace(r) {
			return true
		}
		return false
	})
}

func splitString(s string) []string {
	// split the string into the parts
	bytes := splitWellFormedString(s)
	if len(bytes) <= 0 {
		return nil
	}
	// nothing happend (maybe this is a 0x120x34 string?)
	if bytes[0] != s {
		return bytes
	}
	// split on 0x
	bytes = strings.Split(s, "0x")
	if len(bytes) <= 0 {
		return nil
	}
	// maybe now?
	return splitWellFormedString(strings.Join(bytes, ",0x"))
}

func convert(s string) (string, error) {
	parts := splitString(s)
	if len(parts) <= 0 {
		return "", nil
	}
	var err error
	nums := make([]int64, len(parts))
	for i, b := range parts {
		nums[i], err = strconv.ParseInt(b, 0, 64)
		if err != nil {
			return "", fmt.Errorf("unable to convert %s to number", b)
		}
	}

	var sb strings.Builder
	for i := 0; i < len(nums); i++ {
		fmt.Fprintf(&sb, "%c", nums[i])
	}
	return sb.String(), nil
}
