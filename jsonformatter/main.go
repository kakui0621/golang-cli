package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var input string
    for scanner.Scan() {
 	   input += scanner.Text()
    }

    if scanner.Err() != nil {
 	   fmt.Println("Error reading input:", scanner.Err())
 	   return
    }

    var jsonData interface{}
    if err := json.Unmarshal([]byte(input), &jsonData); err != nil {
 	   fmt.Println("Error parsing JSON:", err)
 	   return
    }

    formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
    if err != nil {
 	   fmt.Println("Error formatting JSON:", err)
 	   return
    }

    fmt.Println(string(formattedJSON))
}
