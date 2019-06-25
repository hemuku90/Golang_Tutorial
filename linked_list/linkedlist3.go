package main

import "fmt"

type Item struct {
	pNext *Item
	val   rune
}

func createList() *Item {

	pHead := &Item{nil, 'a'}

	pCurr := pHead
	for i := 'b'; i <= 'z'; i++ {
		pItem := &Item{nil, i}
		pCurr.pNext = pItem
		pCurr = pItem
	}

	return pHead
}

func printList(pList *Item) {

	pCurr := pList
	for {
		fmt.Printf("%c", pCurr.val)

		if pCurr.pNext != nil {
			pCurr = pCurr.pNext
		} else {
			break
		}
	}
	fmt.Println("")
}

func reverseList(pList *Item) *Item {
	pCurr := pList
	var pTop *Item = nil
	for {
		if pCurr == nil {
			break
		}
		//fmt.Println("pcurr", pCurr)
		pTemp := pCurr.pNext
		//fmt.Println("pTemp", pTemp)
		//fmt.Println("pcurrent.pnext", pCurr.pNext)
		pCurr.pNext = pTop
		//fmt.Println("pTop", pTop)
		pTop = pCurr
		//fmt.Println("pTop", pTop)
		pCurr = pTemp
	}

	return pTop
}

func main() {
	var pList = createList()
	printList(pList)
	printList(reverseList(pList))
}
