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
		fmt.Println("Page", page)
		page = page.nextpage
	}
}

// Insertion into a linked list at the END
func (page *storypage) addToend(text string) {
	pageToAdd := &storypage{text, nil}
	for page.nextpage != nil {
		page = page.nextpage
	}
	page.nextpage = pageToAdd

}

// Inserting into a randome position into a linked list

func (page *storypage) addAfter(text string) {
	newPage := &storypage{text, page.nextpage}
	page.nextpage = newPage
}

func main() {
	page1 := storypage{"It was a dark stormy night on Day 1 !!!", nil}
	page1.addToend("It was a dark stormy night on Day 2 !!!")
	page1.addToend("It was a dark stormy night on Day 3 !!!")
	page1.addAfter("Testing addAfter")
	page1.playstory()
}
