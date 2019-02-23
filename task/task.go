// Package task is a simple scalable work system. It uses channels to queue and distribute work, and collate results.
package task

import (
	"bufio"
	"log"
	"os"
	"sync"
)

// A Task is an abstract interface for real tasks.
type Task interface {
	Process()
	Output()
}

// A Factory is an abstract interface to create Tasks.
type Factory interface {
	Create(line string) Task
}

// Run takes in a Factory, creates a goroutine to read from Stdin and create tasks, and use multiple goroutines to process them. Then it outputs the results.
func Run(f Factory, count int) {
	var wg sync.WaitGroup

	in := make(chan Task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.Create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan Task)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.Output()
	}
}
