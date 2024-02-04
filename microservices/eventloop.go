package main
import "fmt"
func countdown(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := n; i > 0; i-- {
            ch <- i
        }
    }()
    return ch
}

func main() {
    for num := range countdown(5) {
        fmt.Println("Countdown:", num)
    }
}

