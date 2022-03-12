package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type ListItem struct {
	Item   string `json:"item"`
	Status bool   `json:"status"`
}

var L = ListItem{}

var list = []ListItem{}

func (t *ListItem) unmarshalCsv() {
	CsvData, _ := ioutil.ReadFile("todo.csv")
	if string(CsvData) != "" {
		err := json.Unmarshal(CsvData, &list)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	L.unmarshalCsv()
}

func (t *ListItem) savetoCsv() {
	jsonData, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("todo.csv", jsonData, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (t ListItem) Add(Item string) string {
	itemCheck := strings.Trim(Item, " ")
	if itemCheck == "" {
		return "invalid input string, please enter a valid"
	}
	t.Item = Item
	list = append(list, t)
	t.savetoCsv()
	return Item + " successfully added"
}

func (t *ListItem) Done(serlNo string) string {

	serialNo, err := strconv.Atoi(serlNo)
	if err != nil {
		log.Fatal("error marking as done")
	}
	if serialNo == 0 {
		fmt.Println("please input a valid serial no above 0")
		return "please input a valid serial no above 0"
	}
	for i := range list {
		if i == (serialNo - 1) {
			list[i].Status = true
		}
	}
	t.savetoCsv()
	return "item status updated"
}

func (t *ListItem) UnDone(serNo string) string {
	serialNo, err := strconv.Atoi(serNo)
	if err != nil {
		log.Fatal("error converting string")
	}
	if serialNo == 0 {
		fmt.Println("please input a valid serial no above 0")
		return "please input a valid serial no above 0"
	}
	for i := range list {
		if i == (serialNo - 1) {
			list[i].Status = false
			t.savetoCsv()
			fmt.Println("item status updated")
			return "item status updated"
		}

	}
	return "invalid input"
}

func (t *ListItem) CleanUp() bool {
	for i, value := range list {
		if value.Status == true {
			list = append(list[:i], list[i+1:]...)
			t.savetoCsv()
			fmt.Println("list successfully cleaned")
			return true
		}
		continue
	}
	fmt.Println("list is clean")
	return false
}

func (t *ListItem) PrintList() string {
	for i, value := range list {
		if value.Status == false {
			fmt.Printf("%v %v\n", i+1, value.Item)
			continue
		}
	}
	return "The Todo List"
}
