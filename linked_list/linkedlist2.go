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

// Insertion into a linked list
func (page *storypage) addToend(text string) {
	pageToAdd := &storypage{text, nil}
	for page.nextpage != nil {
		page = page.nextpage
	}
	page.nextpage = pageToAdd

}

func main() {
	page1 := storypage{"It was a dark stormy night on Day 1 !!!", nil}
	page1.addToend("It was a dark stormy night on Day 2 !!!")
	page1.addToend("It was a dark stormy night on Day 3 !!!")
	page1.playstory()
}
