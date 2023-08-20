package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

func main() {
    if len(os.Args) < 3 {
 	   fmt.Println("Usage: loganalyzer <log-file-path> <search-pattern>")
 	   return
    }

    filepath := os.Args[1]
    pattern := os.Args[2]

    file, err := os.Open(filepath)
    if err != nil {
 	   fmt.Println("Error opening file:", err)
 	   return
    }
    defer file.Close()

    re, err := regexp.Compile(pattern)
    if err != nil {
 	   fmt.Println("Error compiling regex:", err)
 	   return
    }

    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
 	   line := scanner.Text()
 	   if re.MatchString(line) {
 		   count++
 	   }
    }

    if scanner.Err() != nil {
 	   fmt.Println("Error reading file:", scanner.Err())
 	   return
    }

    fmt.Printf("Pattern '%s' matched %d times in file %s\n", pattern, count, filepath)
}
