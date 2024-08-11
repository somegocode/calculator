package main

import (
    "fmt"
    "strconv"
)


func rome_to_arab(num string) int {
    arab := 0
    symbols := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100}
    for i := range num {
        n := string(num[i])
        if symbols[n] == 0 {
            panic("Value error.")
        } else if i < len(num) - 1 {
            if symbols[n] >= symbols[string(num[i+1])] {
                arab += symbols[n]
            } else {
                arab -= symbols[n]
            }
        } else {
            arab += symbols[n]
        }
    }
    return arab
}


func arab_to_rome(num string) string {
    x := 100
    num_int, _ := strconv.Atoi(num)
    if num_int < 1 {
        panic("Value error.")
    }
    sym := 0
    rome := ""
    symbols := map[int]string{100: "C", 90: "XC", 80: "LXXX", 70: "LXX", 60: "LX", 50: "L", 40: "XL", 30: "XXX", 20: "XX", 10: "X", 9: "IX", 8: "VIII", 7: "VII", 6: "VI", 5: "V", 4: "IV", 3: "III", 2: "II", 1: "I"}
    for num_int > 0 {
        var multiply = num_int / x
        if multiply > 0 {
            sym = x * multiply
            rome += symbols[sym]
            num_int -= sym
        }
        x = x / 10
    }
    return rome
}


func calculate(x int, y int, arif string) string {
    var calc int
    if x < 1 || x > 10 || y < 1 || y > 10 {
        panic("Value error.")
    }
    if arif == "+" {
        calc = x + y
    } else if arif == "-" {
        calc = x - y
    } else if arif == "*" {
        calc = x * y
    } else if arif == "/" {
        calc = x / y
    }
    return strconv.Itoa(calc)
}


func main() {
    var a, b, arif, extra string
    fmt.Scanln(&a, &arif, &b, &extra)
    if extra != "" {
        panic("формат математической операции не удовлетворяет заданию")
    }
    var answer string
    x, err := strconv.Atoi(a)
    if err != nil {
        x := rome_to_arab(a)
        y := rome_to_arab(b)
        c := calculate(x, y, arif)
        answer = arab_to_rome(c)
    } else {
        y, err := strconv.Atoi(b)
        if err != nil {
            panic(err)
        }
        answer = calculate(x, y, arif)
    }
    fmt.Println(answer)
}
