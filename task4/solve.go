package main

import (
    "strings"
    "unicode"
)

func RemoveEven(a []int) []int {
    var res []int
    for _, v := range a {
        if v % 2 != 0 {
            res = append(res, v)
        }
    }
    return res
}

func PowerGenerator(n int) func() int {
    currNum := 1
    return func() int {
        currNum *= n
        return currNum
    }
}

func DifferentWordsCount(s string) (res int) {
    s = strings.ToLower(s)
    var c string
    q := make(map[string]int)

    for i := 0; i < len(s); i++ {
        if !unicode.IsLetter(int32(s[i])) {
            if len(c) > 0 {
                q[c] = 1
            }
            c = ""
        } else {
            c += string(s[i])
        }
    }
    if len(c) > 0 {
        q[c] = 1
    }
    return len(q)
}

/*
func main() {
    input := []int{0, 3, 2, 5}
    result := RemoveEven(input)
    fmt.Println(result)
    gen := PowerGenerator(2)
    fmt.Println(gen()) // Должно напечатать 2
    fmt.Println(gen()) // Должно напечатать 4
    fmt.Println(gen())
    fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
