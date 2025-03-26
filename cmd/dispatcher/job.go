package dispatcher

import (
	"fmt"
	"time"
)

func start_batch() {
	fmt.Println("Batch job started...")
	time.Sleep(5 * time.Second)
	fmt.Println("Batch finished...")
}
