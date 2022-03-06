package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type listItem = struct {
	SN     int
	Item   string
	Status bool
}

var list []listItem

func CsvCreate() {
	csvfile, err := os.Create("data.csv")
	if err != nil {
		log.Fatalf("csv create failed")
	}
	csvwriter := csv.NewWriter(csvfile)
	for _, Input := range list {
		_ = csvwriter.Write(Input)
	}
	csvwriter.Flush()
	csvwriter.Close()
}
func (t *listItem) Add(Item string, index int) bool {
	ItemCheck := strings.Trim(Item, " ")
	if ItemCheck != "" {
		list = append(list, listItem{
			SN:     index,
			Item:   Item,
			Status: false,
		})
		return true
		fmt.Println("Item Successfully Added")
	}
	fmt.Println("please enter a valid input string")
	return false
}
func (t *listItem) Done(serialNo int) bool {
	if serialNo != 0 {
		for i, _ := range list {
			if i == (serialNo - 1) {
				t.Status = true
			}
		}
		fmt.Println("item status updated")
		return true
	}
	fmt.Println("please input a valid serial no above 0")
	return false
}
func (t *listItem) UnDone(serialNo int) bool {
	if serialNo != 0 {
		for i, _ := range list {
			if i == (serialNo-1) && t.Status == true {
				t.Status = false
			}
			fmt.Println("item status updated")
			return true
		}
		fmt.Println("item status is already Undone please enter the correct serial no")
		return false
	}
	fmt.Println("please input a valid serial no above 0")
	return false
}

func (t *listItem) CleanUp() bool {
	for i, value := range list {
		if value.Status == true {
			list = append(list[:i], list[1+1:]...)
			return true
			fmt.Println("Item Successfully Added")
		}
	}
	fmt.Println("please enter a valid input string")
	return false
}
func (t *listItem) PrintList() {
	for i, value := range list {
		if value.Status == false {
			fmt.Printf("%v, %v/n", i+1, value.Item)
			continue
		}
	}
}
