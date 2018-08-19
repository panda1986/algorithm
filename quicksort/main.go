package main

import "fmt"

func quickSort(a []int) {
    if len(a) <= 1 {
        return
    }

    mid := a[0]
    head, tail := 0, len(a) - 1

    for i := 1; i <= tail; {
        if a[i] < mid {
            // 与head值交换，i++
            a[i], a[head] = a[head], a[i]
            head++
            i ++
        } else if a[i] > mid {
            // 与tail值交换，tail--
            a[i], a[tail] = a[tail], a[i]
            tail --
        }
    }

    fmt.Println("....head is", head, a[:head], a[head+1:], tail)
    quickSort(a[:head])
    quickSort(a[head+1:])
    return
}

func main()  {
    a := []int{49, 38, 65, 97, 76, 13, 27}
    quickSort(a)
    fmt.Println(a)
}