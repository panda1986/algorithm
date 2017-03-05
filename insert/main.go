package main

import "fmt"

func InsertSort(src []int) (dst []int) {
    // the outter loop from 2 to n, index is j
    for j := 1; j < len(src); j ++ {
        // inner loop starts at j-1
        k := src[j]
        i := j - 1
        // the goal of each time through the loop is to increase
        for {
            if i < 0 || src[i] < k {
                break
            }
            // keep copying up until we find the place where the key goes
            src[i + 1] = src[i]
            i = i - 1
        }
        src[i + 1] = k
        // is to add one to the length of the things that are sorted
    }
    return src
}

func main()  {
    a := []int{8, 9, 3, 12, 6}
    fmt.Println(fmt.Sprintf("dst:%+v", InsertSort(a)))
}


