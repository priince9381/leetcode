package account_merge

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// func contains(slice []string, value string) bool {
// 	for _, email := range slice {
// 		if email == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// func appendUniqueEmail(arr1, arr2 []string) []string {
// 	// Create a map to store unique strings
// 	uniqueMap := make(map[string]bool)

// 	// Add strings from arr1 to the map
// 	for _, str := range arr1 {
// 		uniqueMap[str] = true
// 	}

// 	// Add strings from arr2 to the map
// 	for _, str := range arr2 {
// 		uniqueMap[str] = true
// 	}

// 	// Extract unique strings from the map
// 	var uniqueStrings []string
// 	for str := range uniqueMap {
// 		uniqueStrings = append(uniqueStrings, str)
// 	}

// 	return uniqueStrings
// }

// func mergelist(accounts [][]string) [][]string {
// 	var accountHolderName string
// 	var accountHolderEmails []string
// 	uniqueAccountHolder := make(map[string][]string)
// 	checkListHolder := make(map[string]bool)
// 	for _, accountDetails := range accounts {
// 		accountHolderName = accountDetails[0]
// 		accountHolderEmails = accountDetails[1:]
// 		if _, ok := checkListHolder[accountHolderName]; ok {
// 			sameAccount := true
// 			i := 0
// 			for sameAccount {
// 				exist := false
// 				if useremails, ok := uniqueAccountHolder[accountHolderName+"_"+strconv.Itoa(i)]; ok {
// 					for _, email := range accountHolderEmails {
// 						exist = contains(useremails, email)
// 						if exist {
// 							uniqueAccountHolder[accountHolderName+"_"+strconv.Itoa(i)] = append(useremails, accountHolderEmails...)
// 							break
// 						}
// 					}
// 					if exist {
// 						break
// 					}
// 					i++
// 					continue
// 				}
// 				if exist {
// 					break
// 				}
// 				uniqueAccountHolder[accountHolderName+"_"+strconv.Itoa(i)] = accountHolderEmails
// 				sameAccount = false
// 			}
// 		} else {
// 			checkListHolder[accountHolderName] = true
// 			uniqueAccountHolder[accountHolderName+"_"+"0"] = accountHolderEmails
// 		}

// 	}
// 	final_list := [][]string{}

// 	for key, value := range uniqueAccountHolder {
// 		var userdetails []string
// 		name := strings.Split(key, "_")
// 		value := appendUniqueEmail(value, []string{})
// 		sort.Strings(value)
// 		userdetails = append(userdetails, name[0])
// 		userdetails = append(userdetails, value...)
// 		final_list = append(final_list, userdetails)

// 	}
// 	return final_list
// }

// func compareLists(list1, list2 [][]string) bool {
// 	// Sort the inner lists in both lists
// 	sortedList1 := sortAndCopy(list1)
// 	sortedList2 := sortAndCopy(list2)

// 	// Compare the sorted lists
// 	return equal(sortedList1, sortedList2)
// }

// func equal(list1, list2 [][]string) bool {
// 	if len(list1) != len(list2) {
// 		return false
// 	}

// 	for i := range list1 {
// 		if !equalInnerList(list1[i], list2[i]) {
// 			return false
// 		}
// 	}

// 	return true
// }

// // equalInnerList compares if two inner lists are equal.
// func equalInnerList(innerList1, innerList2 []string) bool {
// 	if len(innerList1) != len(innerList2) {
// 		return false
// 	}

// 	for i := range innerList1 {
// 		if innerList1[i] != innerList2[i] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func sortAndCopy(list [][]string) [][]string {
// 	sortedList := make([][]string, len(list))
// 	for i, innerList := range list {
// 		sortedInnerList := make([]string, len(innerList))
// 		copy(sortedInnerList, innerList)
// 		sort.Strings(sortedInnerList)
// 		sortedList[i] = sortedInnerList
// 	}
// 	return sortedList
// }

// func accountsMerge(accounts [][]string) [][]string {
// 	//logic
// 	// Create a map ["john":true, "mary": true]
// 	// Unique map ["john_0": [email1,email2, email3], "mary_0": [mary1], "john_1":[mail]]
// 	tempFinalList := mergelist(accounts)
// 	flag := true
// 	for flag {
// 		tempFinalList2 := mergelist(tempFinalList)
// 		if compareLists(tempFinalList, mergelist(tempFinalList2)) {
// 			return tempFinalList
// 		}
// 		tempFinalList = tempFinalList2
// 	}

// 	return mergelist(tempFinalList)
// }

// func AccountTest() {
// 	account := [][]string{
// 		{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
// 		{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
// 		{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
// 		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
// 		{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
// 	}
// 	t := accountsMerge(account)
// 	fmt.Println(t)
// }
