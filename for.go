package main

import ( "fmt")
func main(){
        /* for loop

        /*for i:=0;i<5;i++ {
            fmt.Printf("Karan\n")
        }*/

        /* Switch case

        /*var i int
        fmt.Printf("enter the value of i between (1-5)\n")
        fmt.Scan(&i)
        switch (i) {
        case 1:
                fmt.Printf("entered 1\n")
        case 2:
                fmt.Printf("entered 2\n")
        case 3:
                fmt.Printf("entered 3\n")
        case 4:
                fmt.Printf("entered 4\n")
        case 5:
                fmt.Printf("entered 5\n")
        default:
                fmt.Printf("out of range\n")
        }*/

        /* Iterate through array */
        var x [5] int = [5] int {10,20,30,40,50}
        for i, v:=range x {
                fmt.Printf("Index %d, value %d\n",i,v)
        }
}

