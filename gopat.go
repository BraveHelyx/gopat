package main

import (
    "fmt"
    "flag"
    "strings"
)

const Upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Lower = "abcdefghijklmnopqrstuvwxyz"
const Num = "0123456789"

func main() {
    var size int
    var find string
    flag.IntVar(&size, "size", 16, "The required size of the pattern string.")
    flag.StringVar(&find, "find", "", "The string to locate within the created pattern in order to determine the offset.")
    flag.Parse()

    var pattern string
    uindex := 0
    lindex := 0
    nindex := 0
    for i := 0; i < size; i++ {
        selector := (i % 3)
        switch selector {
        case 0:
            pattern = pattern + string(Upper[uindex])
        case 1:
            pattern = pattern + string(Lower[lindex])
        case 2:
            pattern = pattern + string(Num[nindex])
        }

        if selector == 2 {
            nindex = nindex + 1
            if nindex == 10 {
                nindex = 0
                lindex = lindex + 1
                if lindex == 24 {
                    lindex = 0
                    uindex = uindex + 1
                    if uindex == 24 {
                        fmt.Println("[WARNING] Pattern uniqueness limit has been hit.")
                        uindex = 0
                    }
                }
            }
        }
    }

    if len(find) != 0 {
        offset := strings.Index(pattern, find)
        if offset < 0 {
            fmt.Printf("Substring sequence was not found within created pattern of size %d\n", size)
        } else {
            fmt.Printf("Found substring at offset: %d\n", offset)
        }
    } else {
        fmt.Println(pattern)
    }
}
