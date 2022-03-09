package cmd

import "testing"

func TestListItem_Add(t *testing.T) {
	var add = []struct {
		input1 ListItem
		input  string
		want   string
	}{
		{ListItem{"change oil", false}, "change oil", "item successfully added"},
		{ListItem{"", false}, "", "invalid input string, please enter a valid"},
	}
	for _, test := range add {
		got := test.input1.Add(test.input)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
func TestListItem_Done(t *testing.T) {
	var done = []struct {
		input1 ListItem
		input  string
		want   string
	}{
		{ListItem{"change oil", false}, "1", "item status updated"},
		{ListItem{"", true}, "0", "please input a valid serial no above 0"},
	}
	for _, test := range done {
		got := test.input1.Done(test.input)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
func TestListItem_UnDone(t *testing.T) {
	var undone = []struct {
		input1 ListItem
		input  string
		want   string
	}{
		{ListItem{"change oil", false}, "2", "item status updated"},
		//{ListItem{"change oil", true}, "3", "item status is already Undone please enter the correct serial no"},
		{ListItem{"", false}, "0", "please input a valid serial no above 0"},
	}
	for _, test := range undone {
		got := test.input1.UnDone(test.input)
		if got != test.want {
			t.Errorf("expected %v\n but got %v", test.want, got)
		}
	}
}
func TestListItem_CleanUp(t *testing.T) {
	var cleanUp = []struct {
		input1 ListItem

		want bool
	}{
		{ListItem{"change oil", true}, true},
		{ListItem{"", false}, false},
	}
	for _, test := range cleanUp {
		got := test.input1.CleanUp()
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
func TestListItem_PrintList(t *testing.T) {
	var printList = []struct {
		input1 ListItem
		want   string
	}{
		{ListItem{"change oil", false}, "The Todo List"},
		{ListItem{"change oil", false}, "The Todo List"},
	}
	for _, test := range printList {
		got := test.input1.PrintList()
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
