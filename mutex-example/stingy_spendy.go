package main

import (
	"sync"
	"time"
)

var money = 100
var lock = sync.Mutex{}

func stingy() {
	for i := 0; i <= 1000; i++ {
		lock.Lock()
		vMoney := money
		time.Sleep(1 * time.Millisecond)
		money = vMoney + 10
		lock.Unlock()
	}
	println("Stingy completed")
}

func spendy() {
	for i := 0; i <= 1000; i++ {
		lock.Lock()
		vMoney := money
		time.Sleep(1 * time.Millisecond)
		money = vMoney - 10
		lock.Unlock()
	}
	println("Spendy completed")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}
