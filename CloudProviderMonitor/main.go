package main

import (
	"CloudProviderMonitor/dto"
	"CloudProviderMonitor/task"
)

func main() {
	taskChan := make(chan dto.Task, 100)
	go task.SubscriptEvent(taskChan)
	task.Deploy(taskChan)
}
