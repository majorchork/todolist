package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type ListItem struct {
	Item   string `json:"item"`
	Status bool   `json:"status"`
}

var L = ListItem{}

var list = []ListItem{}

func (t *ListItem) unmarshalCsv() {
	CsvData, _ := ioutil.ReadFile("todo.csv")
	err := json.Unmarshal(CsvData, &list)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	L.unmarshalCsv()
}

func (t *ListItem) savetoCsv() {
	jsonData, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}

	//CsvData = append(CsvData, jsonData...)
	err = ioutil.WriteFile("todo.csv", jsonData, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (t ListItem) Add(Item string) string {
	if Item != "" {
		t.Item = Item
		list = append(list, t)
		t.savetoCsv()
		return "item successfully added"
	}
	return "invalid input string, please enter a valid"
}

func (t *ListItem) Done(serlNo string) []ListItem {

	serialNo, err := strconv.Atoi(serlNo)
	if err != nil {
		panic(err)
	}
	if serialNo != 0 {
		for i := range list {
			if i == (serialNo - 1) {
				list[i].Status = true
				fmt.Println("item status updated")
			}
		}
		t.savetoCsv()
		return list
	}
	fmt.Println("please input a valid serial no above 0")
	return list
}

func (t *ListItem) UnDone(serNo string) []ListItem {
	serialNo, err := strconv.Atoi(serNo)
	if err != nil {
		panic(err)
	}
	if serialNo != 0 {
		for i := range list {
			if i == (serialNo-1) && list[i].Status == true {
				list[i].Status = false
			}
			t.savetoCsv()
			fmt.Println("item status updated")
			return list
		}
		fmt.Println("item status is already Undone please enter the correct serial no")
		return list
	}
	fmt.Println("please input a valid serial no above 0")
	return list
}

func (t *ListItem) CleanUp() bool {
	for i, value := range list {
		if value.Status == true {
			list = append(list[:i], list[1+1:]...)
			t.savetoCsv()
			fmt.Println("Item Successfully Added")
			return true
		}
	}
	fmt.Println("please enter a valid input string")
	return false
}

func (t *ListItem) PrintList() {
	for i, value := range list {
		if value.Status == false {
			fmt.Printf("%v, %v\n", i+1, value.Item)
			continue
		}
	}
}
