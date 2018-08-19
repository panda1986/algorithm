package main

import "fmt"

func kmp_table(w []string, t []int)  {
    pos := 1 // the current position we are computing in t
    cnd := 0 // the zero-based index in w of the next char of the current candidate substring

    t[0] = -1

    for {
        if pos >= len(w) {
            break
        }
        if w[pos] == w[cnd] {
            fmt.Println(fmt.Sprintf("..w[%v](%v)=w[%v](%v)", pos, w[pos], cnd, w[cnd]))
            t[pos] = t[cnd]
            fmt.Println(fmt.Sprintf("..when w match, set t[%v]=t[%v] %v", pos, cnd, t[pos]))
            pos = pos + 1
            cnd = cnd + 1
        } else {
            fmt.Println(fmt.Sprintf("..w[%v](%v) mismatch w[%v](%v)", pos, w[pos], cnd, w[cnd]))
            t[pos] = cnd
            cnd = t[cnd] // to increase performance
            fmt.Println(fmt.Sprintf("..set t[%v]=%v, cnd=%v", pos, t[pos], cnd))
            for {
                if cnd >= 0 && w[pos] != w[cnd] {
                    cnd = t[cnd]
                    fmt.Println(fmt.Sprintf("..second for loop set cnd=%v", cnd))
                } else {
                    break
                }
            }
            pos = pos + 1
            cnd = cnd + 1
            fmt.Println(fmt.Sprintf("..after mismatch, pos=%v, cnd=%v", pos, cnd))
        }
    }

    //t[pos] = cnd
}

func main()  {
    w := []string{"A", "B", "A", "B", "A", "A", "B"}
    //w := []string{"A", "B", "A", "C", "A", "B", "A", "B", "C"}
    //w := []string{"P", "A", "R", "T", "I", "C", "I", "P", "A", "T", "E", "", "I", "N", "", "P", "A", "P"}
    t := []int{}
    for i := 0; i < len(w); i ++ {
        t = append(t, 0)
    }
    kmp_table(w, t)

    for i := 0; i < len(t); i ++ {
        fmt.Println(i, t[i])
    }

}