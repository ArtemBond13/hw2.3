package stats

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

//
type Transaction struct {
	Id      string
	From    string
	To      string
	Amount  int64
	Created int64
}

// "Часть" хранит номер месяца и указатель на транкзакцию
type part struct {
	monthTimestamp int64 // Метка времени месяца
	transactions   []*Transaction
}

func Sum(transactions []int64) int64 {
	result := int64(0)
	for _, transaction := range transactions {
		result += transaction
	}
	return result
}

func SumConcurrently(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := len(transactions) / goroutines
	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			total += Sum(part) // FIXME: shared memory bug, discuss later
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}

func SumConcurrentlyMonth(start, finish time.Time) int64 {
	// слайс с указателем на структуру part, len и cap = 0
	months := make([]*part, 0)

	// Опрделить сколько месяцев
	next := start
	for next.Before(finish) { // пока next != finish
		months = append(months, &part{monthTimestamp: next.Unix()}) // добавить в мвссив отметку времени
		next = next.AddDate(0, 1, 0)                                // next прибавляем 1 месяц
	}
	months = append(months, &part{monthTimestamp: finish.Unix()})

	// создаем для транкзаций
	transactions := make([]*Transaction, 0)
	// TODO: заполняете transactions
	for i, transaction := range transactions {
		// TODO: находите нужный месяц для совпадения
		month := months[i]
		month.transactions = append(month.transactions, transaction)
	}

	wg := sync.WaitGroup{}

	total := int64(0)
	for _, i := range transactions {
		wg.Add(1)
		i := i
		go func() {
			fmt.Println("start")
			total += i.Amount // FIXME: shared memory bug, discuss later
			wg.Done()
		}()
	}
	wg.Wait()
	return total
}

func SortSlice(transactions []int64) []int64 {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i] > transactions[j]

	})
	return transactions
}

func SortSliceStable(transactions []int64) []int64 {
	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i] > transactions[j]

	})
	return transactions
}
