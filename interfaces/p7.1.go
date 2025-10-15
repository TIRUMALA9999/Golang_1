//if we need to retrieve many values at once

package main
import "fmt"

type Guitar struct{}

func main() {
    vals := []interface{}{25, "Teja", true, 3.14, Guitar{}}

    for _, v := range vals {
        switch val := v.(type) {
        case int:
            fmt.Println("Integer:", val*2)
        case string:
            fmt.Println("String:", val+"!!")
        case bool:
            fmt.Println("Boolean:", val)
        case float64:
            fmt.Println("Float:", val+1.1)
        case Guitar:
            fmt.Println("Guitar struct:", val)
        default:
            fmt.Println("Unknown type")
        }
    }
}
