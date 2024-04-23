package main
import "fmt"
import "math"

func main() {
        var x float64
        fmt.Printf("Enter the value of x=")
        fmt.Scanf("%f",&x)
        TruncatedInt:= int(math.Trunc(x))
        fmt.Printf("x=%d\n",TruncatedInt)
}

