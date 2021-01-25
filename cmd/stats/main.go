package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.3.git/pkg/stats"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(4)

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	const users = 10               // количество клиентов
	const transactionsPerUser = 10 // каждый совершил в месяце транкзакций
	const transactionAmount = 3_00 // в таком количестве

	// 01.01.2020 00:00 в локальном часовом поясе
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local) // Time
	finish := time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local)
	months := make([]int64, 0)
	next := start
	for next.Before(finish) {
		months = append(months, next.Unix())
		next = next.AddDate(0, 1, 0)
	}
	fmt.Printf("%#v\n", months)

	transactions := make([]int64, users*transactionsPerUser)
	for index := range transactions {
		if index%9 == 0 {
			transactions[index] = transactionAmount + 2
		} else if index%10 == 0 {
			transactions[index] = transactionAmount + 10
		} else {
			transactions[index] = transactionAmount
		}
	}

	//stats.SumConcurrentlyMonth(start, finish)

	total := int64(0)
	const partsCount = 10
	partSize := len(transactions) / partsCount
	for i := 0; i < partsCount; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			fmt.Println("start")
			total += stats.Sum(part) // FIXME: shared memory bug, discuss later
		}()
	}
	time.Sleep(time.Minute)
	fmt.Println(total)

	sortSlice := stats.SortSlice(transactions)
	fmt.Printf("%v\n", sortSlice)
	stableSlice := stats.SortSliceStable(transactions)
	fmt.Printf("%v\n", stableSlice)
}
