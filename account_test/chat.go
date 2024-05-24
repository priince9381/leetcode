package account_merge

import (
	"fmt"
	"sort"
)

// mergeList merges the accounts.
func mergeList(accounts [][]string) [][]string {
	emailToName := make(map[string]string)
	graph := make(map[string][]string)

	// Build the graph
	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			graph[email] = append(graph[email], account[1])
			graph[account[1]] = append(graph[account[1]], email)
			emailToName[email] = name
		}
	}

	// Traverse the graph to find connected components
	visited := make(map[string]bool)
	var result [][]string
	for email := range graph {
		if !visited[email] {
			component := make([]string, 0)
			dfs(graph, visited, email, &component)
			sort.Strings(component)
			component = append([]string{emailToName[email]}, component...)
			result = append(result, component)
		}
	}

	return result
}

// dfs performs depth-first search on the graph to find connected components.
func dfs(graph map[string][]string, visited map[string]bool, email string, component *[]string) {
	visited[email] = true
	*component = append(*component, email)
	for _, neighbor := range graph[email] {
		if !visited[neighbor] {
			dfs(graph, visited, neighbor, component)
		}
	}
}

// compareLists compares two lists of lists and returns true if they contain the same values, false otherwise.
func compareLists(list1, list2 [][]string) bool {
	if len(list1) != len(list2) {
		return false
	}

	for i := range list1 {
		if !equalInnerList(list1[i], list2[i]) {
			return false
		}
	}

	return true
}

// equalInnerList compares if two inner lists are equal.
func equalInnerList(innerList1, innerList2 []string) bool {
	if len(innerList1) != len(innerList2) {
		return false
	}

	for i := range innerList1 {
		if innerList1[i] != innerList2[i] {
			return false
		}
	}

	return true
}

func accountsMerge(accounts [][]string) [][]string {
	//logic
	// Create a map ["john":true, "mary": true]
	// Unique map ["john_0": [email1,email2, email3], "mary_0": [mary1], "john_1":[mail]]
	tempFinalList := mergeList(accounts)
	flag := true
	for flag {
		tempFinalList2 := mergeList(tempFinalList)
		if compareLists(tempFinalList, mergeList(tempFinalList2)) {
			return tempFinalList
		}
		tempFinalList = tempFinalList2
	}

	return mergeList(tempFinalList)
}

func AccountTest() {
	account := [][]string{
		{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
		{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
		{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
		{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
	}
	t := accountsMerge(account)
	fmt.Println(t)
}
