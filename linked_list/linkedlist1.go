package main

import "fmt"

//Linked list struct
type storypage struct {
	text     string
	nextpage *storypage
}

// Print linked list
func (page *storypage) playstory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextpage
	}
}

func main() {
	page1 := storypage{"It was a dark stormy night on Day 1 !!!", nil}
	page2 := storypage{"It was a dark stormy night on Day 2 !!!", nil}
	page3 := storypage{"It was a dark stormy night on Day 3!!!", nil}
	page1.nextpage = &page2
	page2.nextpage = &page3
	//fmt.Println(page1, "\n", page2, "\n", page3)
	page1.playstory()
}
