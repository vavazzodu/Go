package main

import (
        "fmt"
)
func main(){
        arr:= [5]int {10,20,30,40,50}
        slic1 := arr[1:3]  //slic1 contains elements at 1st and 2nd index, not 3rd
        slic2 := arr[2:5]  //slic2 contains elements from 2nd index to 4th index not 5th

        fmt.Printf("Printing slice 1:\n")
        for i:=0; i<len(slic1); i++ {
                fmt.Printf("slic1[%d] = %d\n",i,slic1[i])
        }

        fmt.Printf("Printing slice 2:\n")
        for i:=0; i<len(slic2); i++ {
                fmt.Printf("slic2[%d] = %d\n",i,slic2[i])
        }
}

