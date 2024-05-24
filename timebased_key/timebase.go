package timebasedkey

import "fmt"

type TimeMap struct {
	head *Node
}

type Node struct {
	Key       string
	Value     string
	Timestamp int
	NextStamp *Node
}

func Constructor() TimeMap {
	return TimeMap{head: &Node{}}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	newNode := &Node{
		Key:       key,
		Value:     value,
		Timestamp: timestamp,
	}

	current := tm.head
	for current.NextStamp != nil && current.NextStamp.Timestamp <= timestamp {
		current = current.NextStamp
	}
	newNode.NextStamp = current.NextStamp
	current.NextStamp = newNode
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	current := tm.head
	for current.NextStamp != nil {
		if current.NextStamp.Timestamp > timestamp || current.NextStamp.Timestamp == timestamp {
			break
		}
		current = current.NextStamp
	}
	for current.NextStamp.Timestamp > timestamp && current.Key == key {
		return current.Value
	}
	return ""
}

func TimesBaseMain() {
	// ["ctondw","ztpearaw",1],["vrobykydll","hwliiq",2],["gszaw","ztpearaw",3],["ctondw","gszaw",4],["gszaw",5]
	obj := Constructor()
	obj.Set("ctondw", "ztpearaw", 1)
	obj.Set("vrobykydll", "hwliiq", 2)
	obj.Set("gszaw", "ztpearaw", 3)
	obj.Set("ctondw", "gszaw", 3)
	fmt.Println(obj.Get("gszaw", 5))
}
