package go_parallel

func ForLoop(data []interface{}, goroutineCount int, f func(interface{}, int)) {
	myChan := make(chan interface{}, len(data))

	for i := 0; i < goroutineCount; i++ {
		go func(i int) {
			for {
				f(<-myChan, i)
			}
		}(i)
	}

	for _, item := range data {
		myChan <- item
	}
}