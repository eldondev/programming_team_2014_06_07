// A _line filter_ is a common type of program that reads
// input on stdin, processes it, and then prints some
// derived result to stdout. `grep` and `sed` are common
// line filters.

// Here's an example line filter in Go that writes a
// capitalized version of all input text. You can use this
// pattern to write your own Go line filters.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    array := strings.Fields(scanner.Text())
    mult2 := func(f e) e { 
        switch f.(type) {
        case int:
            return f.(int) * 2 
        case string:
            return f.(string) + f.(string)
        }   
        return f
    }
    Map(&array, mult2);
    fmt.Println(array) 


    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())

        // Write out the uppercased line.
        fmt.Println(ucl)
    }

    // Check for errors during `Scan`. End of file is
    // expected and not reported by `Scan` as an error.
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
