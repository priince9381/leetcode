package gobasics

import (
	"fmt"
	"sync"
	"time"
)

var PrimeArray []int
var mu sync.Mutex

func isPrime(number int) {
	count := 0
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count += 1
		}
	}
	if count <= 2 {
		mu.Lock()
		PrimeArray = append(PrimeArray, number)
		mu.Unlock()
	}
}

func isPrimeByGoRoutine() {
	wg := &sync.WaitGroup{}
	checkPrimeTill := 1000
	for i := 1; i <= checkPrimeTill; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			isPrime(n)
		}(i)
	}
	wg.Wait()
}

func isPrimeNormal() {
	checkPrimeTill := 1000
	for i := 1; i <= checkPrimeTill; i++ {
		isPrime(i)
	}
}

func TestMainFunction() {
	startGoRoutineTime := time.Now()
	isPrimeByGoRoutine()
	endGoRoutineTimese := time.Since(startGoRoutineTime)
	fmt.Println("Go Routine Execution Time", endGoRoutineTimese)
	fmt.Println(PrimeArray)
	PrimeArray = []int{}
	normalisRunTime := time.Now()
	isPrimeNormal()
	endNormalRunTime := time.Since(normalisRunTime)
	fmt.Println("Normal Run Time", endNormalRunTime)
	fmt.Println(PrimeArray)
}
