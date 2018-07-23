package main

import (
	runner "./x"
	"log"
	"os"
	"time"
)

// 这个 timeout 是 r.Add 里所有方法的总时长，跟我一开始想的不一样。。。
const timeout = 11 * time.Second

func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("Processor Done - Task #%d.", id)
	}
}
