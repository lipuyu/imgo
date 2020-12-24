package conf

import (
	"fmt"
	"github.com/whiteshtef/clockwork"
)

func Run() {
	sched := clockwork.NewScheduler()
	sched.Schedule().Every(1800).Seconds().Do(something)
	sched.Run()
}

func something() {
	fmt.Println("hello")
}
