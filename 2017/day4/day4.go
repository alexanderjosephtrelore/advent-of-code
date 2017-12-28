package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"sort"
)

type fn func([]string) bool

func countValidPassphrases(fileName string, f fn) int {
	counter := 0
	file, _ := os.Open(fileName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		var list []string
		for _, p := range strings.Fields(scanner.Text()) {
			list = append(list, p)
		}
		if f(list) {
			counter++
		}
	}
	return counter
}

func isValidPassphrase(input []string) bool {
	set := StringSet{}
	for _, value := range input {
		if set.add(value) == false {
			return false
		}
	}
	return true
}

func isOrderedValidPassphrase(input []string) bool {
	set := StringSet{}
	for _, value := range input {
		if set.addOrdered(value) == false {
			return false
		}
	}
	return true
}

type StringSet struct {
	set map[string]bool
}

func (set *StringSet) add(i string) bool {
	if set.set == nil {
		set.set = make(map[string]bool)
	}
	_, found := set.set[i]
	set.set[i] = true
	return !found   //False if it existed already
}

func (set *StringSet) addOrdered(i string) bool {
	tmp := strings.Split(i, "")
	sort.Strings(tmp)
	i = strings.Join(tmp, "")
	if set.set == nil {
		set.set = make(map[string]bool)
	}
	_, found := set.set[i]
	set.set[i] = true
	return !found   //False if it existed already
}

func main()  {
	fmt.Println("Part 1:", countValidPassphrases("input.txt", isValidPassphrase))
	fmt.Println("Part 2:", countValidPassphrases("input.txt", isOrderedValidPassphrase))
}