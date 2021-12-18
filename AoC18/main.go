package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

type value struct {
	v     int
	depth int
}

type tree []value

type stack tree

func main() {

	start := time.Now()
	task()
	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)

}

func task() {

	file, err := os.Open("./input18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {

		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	Task01(lines)
	Task02(lines)

}

func Task01(lines []string) {

	exp := tree{}
	for _, line := range lines {
		exp = add(exp, parse(line))
		exp = reduce(exp)
	}
	fmt.Println(magnitude(exp))
}

func Task02(lines []string) {
	values := []tree{}

	for _, line := range lines {
		values = append(values, reduce(parse(line)))
	}

	max := 0
	for i, a := range values {
		for j, b := range values {
			if i != j {
				exp := add(a, b)
				exp = reduce(exp)
				m := magnitude(exp)
				if m > max {
					max = m
				}
				exp = add(b, a)
				exp = reduce(exp)
				m = magnitude(exp)
				if m > max {
					max = m
				}
			}
		}
	}
	fmt.Println(max)
}

func parse(line string) tree {
	res := tree{}
	depth := 0
	i := 0
	for i < len(line) {
		switch {
		case line[i] == '[':
			depth++
			i++
		case line[i] == ']':
			depth--
			i++
		case line[i] == ',':
			i++
		case unicode.IsDigit(rune(line[i])):
			v := int(line[i] - '0')
			i++
			res = append(res, value{v, depth})
		}
	}
	return res
}

func add(t1 tree, t2 tree) tree {
	if len(t1) == 0 {
		return t2
	}
	if len(t2) == 0 {
		return t1
	}

	res := tree{}
	res = append(res, t1...)
	res = append(res, t2...)
	for i, _ := range res {
		res[i].depth += 1
	}
	return res
}

func reduce(l tree) tree {
	reduced := true
	for reduced {
		l, _ = explode(l)
		l, reduced = split(l)
	}
	return l
}

func explode(t tree) (tree, bool) {
	res := t
	reduced := false
	i := 0
	for i < len(res)-1 {
		if res[i].depth >= 5 && res[i].depth == res[i+1].depth {
			left := res[i].v
			right := res[i+1].v
			res = removeIndex(res, i+1)
			res[i].v = 0
			res[i].depth -= 1
			if i > 0 {
				res[i-1].v += left
			}
			if i < len(res)-1 {
				res[i+1].v += right
			}
			reduced = true
		} else {
			i++
		}
	}
	return res, reduced
}

func removeIndex(s tree, i int) tree {
	res := make(tree, i, len(s)-1)
	copy(res, s[:i])
	res = append(res, s[i+1:]...)
	return res
}

func split(l tree) (tree, bool) {
	for i := 0; i < len(l); i++ {
		if l[i].v >= 10 {
			a := l[i].v / 2
			b := l[i].v - a
			newDepth := l[i].depth + 1
			res := replaceIndex(l, i, value{a, newDepth}, value{b, newDepth})
			return res, true
		}
	}
	return l, false
}

func replaceIndex(s tree, i int, a value, b value) tree {
	res := make(tree, i, len(s)+1)
	copy(res, s[:i])
	res = append(res, a)
	res = append(res, b)
	return append(res, s[i+1:]...)
}

func magnitude(t tree) int {
	stack := newStack()
	for _, v := range t {
		stack.pushMagnitude(v)

	}
	top := stack.removeStack()
	return top.v
}

func (s *stack) pushMagnitude(v value) {
	if s.IsEmpty() {
		*s = append(*s, v)
		return
	}

	top := s.getLastEntry()
	if v.depth == top.depth {
		s.removeStack()
		s.pushMagnitude(value{3*top.v + 2*v.v, v.depth - 1})
	} else {
		*s = append(*s, v)
	}
}

func newStack() stack {
	return make([]value, 0)
}

func (s *stack) removeStack() value {
	l := len(*s)
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top
}

func (s *stack) getLastEntry() value {
	return (*s)[len(*s)-1]
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
