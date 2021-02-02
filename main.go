package main

import (
	"fmt"
	"os"
	"strings"
)

func mergeStrings(strs []string) string {
	sum := ""
	for _, arg := range strs {
		sum = sum + " " + arg
	}
	return strings.Trim(sum, " ")
}

func main() {
	srcPath := mergeStrings(os.Args[1:])
	if srcPath == "--help" {
		fmt.Println(`WSLPATH https://github.com/uanid/wslpath v1.0
Usage: wslpath C:\\Users\\musong
return: /mnt/c/Users/musong`)
		return
	}

	driveIndex := strings.Index(srcPath, ":")
	driveLetter := strings.ToLower(srcPath[0:driveIndex])

	destPath := "/mnt/" + driveLetter + "/" + strings.Replace(srcPath[driveIndex+2:], "\\", "/", -1)
	fmt.Print(destPath)
}
