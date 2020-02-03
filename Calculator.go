package main

import (
	"fmt"
	"strconv"
)

const Max_Size = 100

type Stack struct {
	data [Max_Size]string
	top  int
}

func (s *Stack) init() {
	s.top = -1
}

func (s *Stack) push(elem string) {
	if s.top == Max_Size-1 {
		//	fmt.Println("Error: the stack is full")
	} else {
		s.top++
		s.data[s.top] = elem
	}
}

func (s *Stack) pop() string {
	var elem string
	if s.top < 0 {
		//	fmt.Println("Error: the stack is empty")
	} else {
		elem = s.data[s.top]
		s.top--
	}
	return elem
}

func main() {
	//read the arithmetic expression
	fmt.Println("Input an expression like (10/2+3)*4-5")
	var expr string
	fmt.Scanln(&expr)
	//transform the expression into Reverse Polish notation
	RPN := Get_RPN(expr)
	//calculate by RPN
	Calculate(RPN)
}

func Get_RPN(src string) [Max_Size]string {
	var dst [Max_Size]string
	var op Stack
	op.init()
	var j int
	var t string
	for i := 0; i < len(src); {
		if i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9") { //It's operand
			for t = ""; i < len(src) && (string(src[i]) >= "0" && string(src[i]) <= "9"); i++ {
				t = t + string(src[i])
			}
			dst[j] = t
			j++
		}

		if i < len(src) && string(src[i]) == "(" {
			op.push(string(src[i]))
			i++
		} else if i < len(src) && string(src[i]) == ")" {
			for op.top != -1 && op.data[op.top] != "(" {
				dst[j] = op.pop()
				j++
			}
			op.pop() //pop "("
			i++
		} else if i < len(src) && op.top == -1 {
			op.push(string(src[i]))
			i++
		} else if i < len(src) && op.data[op.top] == "(" {
			op.push(string(src[i]))
			i++
		} else if i < len(src) && (string(src[i]) == "*" || string(src[i]) == "/") && (op.data[op.top] == "+" || op.data[op.top] == "-") {
			op.push(string(src[i]))
			i++
		} else if i < len(src) {
			dst[j] = op.pop()
			j++
		}
	}
	for op.top != -1 {
		dst[j] = op.pop()
		j++
	}
	return dst
}

func Calculate(RPN [Max_Size]string) {
	var num Stack
	num.init()
	for i := 0; i < len(RPN); i++ {
		if RPN[i] != "" {
			if string(RPN[i][0]) >= "0" && string(RPN[i][0]) <= "9" {
				num.push(RPN[i])
			} else if RPN[i] == "+" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := x + y
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "-" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y - x
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "*" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y * x
				num.push(strconv.Itoa(n))
			} else if RPN[i] == "/" {
				x, _ := strconv.Atoi(num.pop())
				y, _ := strconv.Atoi(num.pop())
				n := y / x
				num.push(strconv.Itoa(n))
			}
		}
	}
	fmt.Println(num.data[0])
}
