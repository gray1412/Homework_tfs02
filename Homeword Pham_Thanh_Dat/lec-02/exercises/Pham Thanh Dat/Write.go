package main

import (  
    "fmt"
    "os"
)
func BubbleSort(arr [6]int)([6] int) {
    for x:= len(arr); x > 1; x-- {
        for i := 1; i < x; i++ {
            if(arr[i] < arr[i-1]) {
                arr[i], arr[i-1] = arr[i-1], arr[i]
            }

        }
    }
    return arr
}
func main() {  
    arr := [6]int{1,9,5,3,7,6}
    f, err := os.Create("BubbleSort.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(fmt.Sprint(BubbleSort(arr)))
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}