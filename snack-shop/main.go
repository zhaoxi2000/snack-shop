package main

import "fmt"
import "time"

type Task interface {
	finish() int
	isDone() bool
}

type BubbleTeaOrder struct {
	orderedTime  time.Time
	finishedTime time.Time
}

type JuiceOrder struct {
	orderedTime  time.Time
	finishedTime time.Time
}
type HotDogOrder struct {
	orderedTime  time.Time
	finishedTime time.Time
}

func main() {
	fmt.Println("App is starting...")

}
