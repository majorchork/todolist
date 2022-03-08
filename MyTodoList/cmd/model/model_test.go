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
func TestListItem_UnDone(t *testing.T) {
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
func TestListItem_CleanUp(t *testing.T) {
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
func TestListItem_PrintList(t *testing.T) {
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
