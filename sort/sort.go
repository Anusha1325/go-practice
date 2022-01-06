package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"sort"
)

type fruits struct {
	Fruit string
	Quantity int
}

type ByQuantity []fruits

func (a ByQuantity) Len() int {return len(a)}
func (a ByQuantity) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByQuantity) Less(i, j int) bool {return a[i].Quantity < a[j].Quantity}

func main() {
	//data, _ := os.ReadFile("data.txt")
	//fmt.Println(data)
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	items := []fruits{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		var item fruits
		item.Fruit = parts[0]
		item.Quantity, _ = strconv.Atoi(parts[1])
		items = append(items, item)
	}
	//fmt.Println(items)
	
	sort.Sort(ByQuantity(items))
	for _, v := range items {
		fmt.Println(v.Fruit, v.Quantity)
	}
}