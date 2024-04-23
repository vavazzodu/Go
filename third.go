package main

import (
        "fmt"
        "strings"
        "bufio"
        "os"
)

func main(){
        fmt.Printf("Enter a string:")
        reader := bufio.NewReader(os.Stdin)
        s1, err := reader.ReadString('\n')
        if err != nil{
                fmt.Printf("Error\n")
        }
        s2 := strings.TrimSpace(s1)
        if (strings.HasPrefix (s2, "i") || strings.HasPrefix (s2, "I")) &&
            (strings.HasSuffix(s2, "n") || strings.HasSuffix(s2, "N")) {
                if strings.ContainsAny(s2, "a") || strings.ContainsAny(s2, "A"){
                        fmt.Printf("Found!\n")
                }else{
                        fmt.Print("Not Found!\n")
                }
        }else{
                fmt.Print("Not Found!\n")
        }
}

