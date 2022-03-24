package go_parallel

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestParallelForLoop(t *testing.T) {
	const sliceSize int = 111
	var data = make([]string, 0, sliceSize)
	for i := 1; i <= sliceSize; i++ {
		data = append(data, strconv.Itoa(i))
	}

	newData := make([]interface{}, len(data))

	for i, _ := range data {
		newData[i] = data[i]
	}

	ForLoop(newData, 20, func(item interface{}, workerNumber int) {
		number := item.(string)
		println(fmt.Sprintf(`Run %d`, workerNumber))
		time.Sleep(time.Duration(rand.Int31n(10)) * time.Second)
		println(fmt.Sprint(number))
		println(fmt.Sprintf(`Stop %d`, workerNumber))
	})

	time.Sleep(time.Minute)
}