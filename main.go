package main

import plusone "test/plus_one"

// "test/findallanagram"

// "lengthOfLastWord"

func contain(key int, list []int) bool {
	for _, value := range list {
		if key == value {
			return true
		}
	}
	return false
}

func VisitiedRoom(arrayKey []int, visitedroom []int, rooms [][]int) []int {
	if len(visitedroom) == 0 {
		return arrayKey
	}
	kevValue := visitedroom[0]
	visitedroom = visitedroom[1:]
	tempKeys := rooms[kevValue]
	for _, tempKey := range tempKeys {
		if !contain(tempKey, arrayKey) {
			arrayKey = append(arrayKey, tempKey)
			visitedroom = append(visitedroom, tempKey) // [1,2,3]
		}
	}
	VisitiedRoom(arrayKey, visitedroom, rooms)
	return arrayKey
}

func canVisitAllRooms(rooms [][]int) bool {
	arrayKey := []int{}
	arrayKey = append(arrayKey, rooms[0]...) // [1]
	visitedroom := append(arrayKey, rooms[0]...)
	// for _, keyValue := range arrayKey {
	// 	tempKeys := rooms[keyValue] // [2] [3]
	// 	fmt.Println(tempKeys, arrayKey)
	// 	for _, tempKey := range tempKeys {
	// 		if !contain(tempKey, arrayKey) {
	// 			arrayKey = append(arrayKey, tempKey) // [1,2,3]
	// 		}
	// 	}

	// }
	arrayKey = VisitiedRoom(arrayKey, visitedroom, rooms)
	if len(arrayKey) != len(rooms)-1 {
		return false
	}
	return true

}

func main() {
	// rooms := [][]int{}
	// rooms = append(rooms, []int{1})
	// rooms = append(rooms, []int{2})
	// rooms = append(rooms, []int{3})
	// rooms = append(rooms, []int{})
	// canVisitAllRooms(rooms)
	// account_merge.AccountTest()
	// pivot.PivotTest()
	// timebasedkey.TimesBaseMain()

	// inordertraversal.TestInorderTraversal()
	// palindronstring.TestPalindron()
	// lengthOfLastWord.TestLastWord()
	// lastworldlength.TestLastWord()
	// kthsmallestelement.TestKthSmallestElement()
	// wordsearch.Testexist()
	// findallanagram.TestFindAllAnagras()
	// searchinsetpostion.TestSearchInsert()
	plusone.TestplusOne()
}
