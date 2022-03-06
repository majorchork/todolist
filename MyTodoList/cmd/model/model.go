package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type listItem = struct {
	SN     int
	Item   string
	Status bool
}

var list []listItem
var CsvData, err = ioutil.ReadFile("todo.csv")

//func CsvCreate() {
//	csvfile, err := os.Create("data.csv")
//	if err != nil {
//		log.Fatalf("csv create failed")
//	}
//	csvwriter := csv.NewWriter(csvfile)
//	for _, Input := range list {
//		_ = csvwriter.Write(Input)
//	}
//	csvwriter.Flush()
//	csvwriter.Close()
//}
func unmarshalCsv() {
	err := json.Unmarshal(CsvData, &list)
	if err != nil {
		log.Fatal(err)
	}
}
func savetoCsv() {
	jsonData, err := json.Marshal(list)
	if err != nil {
		log.Fatal(err)
	}

	CsvData = append(CsvData, jsonData...)
	ioutil.WriteFile("todo.csv", CsvData, 0777)
}
func (t *listItem) Add(Item string, index int) []listItem {
	ItemCheck := strings.Trim(Item, " ")
	if ItemCheck != "" {
		list = append(list, listItem{
			SN:     index,
			Item:   Item,
			Status: false,
		})
		//jsonData, err := json.Marshal(list)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//CsvData = append(CsvData, jsonData...)
		//ioutil.WriteFile("todo.csv", CsvData, 0777)
		savetoCsv()
		return list
		fmt.Println("Item Successfully Added")
	}
	fmt.Println("please enter a valid input string")
	return list
}
func (t *listItem) Done(serialNo int) bool {
	//err := json.Unmarshal(CsvData, &list)
	//if err != nil {
	//	log.Fatal(err)
	//}
	unmarshalCsv()
	if serialNo != 0 {
		for i, _ := range list {
			if i == (serialNo - 1) {
				t.Status = true
			}
		}
		savetoCsv()
		fmt.Println("item status updated")
		return true
	}
	fmt.Println("please input a valid serial no above 0")
	return false
}
func (t *listItem) UnDone(serialNo int) bool {
	unmarshalCsv()
	if serialNo != 0 {
		for i, _ := range list {
			if i == (serialNo-1) && t.Status == true {
				t.Status = false
			}
			savetoCsv()
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
	unmarshalCsv()
	for i, value := range list {
		if value.Status == true {
			list = append(list[:i], list[1+1:]...)
			savetoCsv()
			return true
			fmt.Println("Item Successfully Added")
		}
	}
	fmt.Println("please enter a valid input string")
	return false
}
func (t *listItem) PrintList() {
	unmarshalCsv()
	for i, value := range list {
		if value.Status == false {
			fmt.Printf("%v, %v/n", i+1, value.Item)
			continue
		}
	}
}
