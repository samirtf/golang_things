package main

import "fmt"
/* https://www.tutorialspoint.com/go/go_variables.htm */

func main() {
    var i, j, k int;
    i = 1;
    j = 2;
    k = i + j;
    var cd, ch byte;
    var someInt = 42;
    var f, salary float32;
    var d = 42;
    var x float64;
    x = 20.0;

    var y1 float64 = 20.0
    y2 := 42
    y3 := 42.0

    var a, b, c = 3, 4.0, "foo"

    fmt.Println(i)
    fmt.Println(j)
    fmt.Println(k)
    fmt.Println(cd)
    fmt.Println(ch)
    fmt.Println(someInt)
    fmt.Println(f)
    fmt.Println(salary)
    fmt.Println(d)
    fmt.Println(x)
    fmt.Printf("x is of type %T\n", x)
    fmt.Println(y1)
    fmt.Println(y2)
    fmt.Println(y3)
    fmt.Printf("y1 is of type %T\n", y1)
    fmt.Printf("y2 is of type %T\n", y2)
    fmt.Printf("y3 is of type %T\n", y3)
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Printf("a is of type %T\n", a)
    fmt.Printf("b is of type %T\n", b)
    fmt.Printf("c is of type %T\n", c)
    
}
