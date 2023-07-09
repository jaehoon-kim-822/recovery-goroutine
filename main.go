package main

var count int = 1

type TaskList struct {
	tasks    []int
	progress int
}

var jobList []*TaskList

func main() {
	println("Hello, World!")

	f := func(c chan int, i int, taskList *TaskList) {
		if taskList.progress != 0 {
			println("2nd Execute task: ", i, "call count: ", count, "task length: ", len(taskList.tasks), " progress: ", taskList.progress)
		} else {
			println("1st Execute task: ", i, "call count: ", count, "task length: ", len(taskList.tasks), " progress: ", taskList.progress)
		}
		count++
		println("Send channel: ", i)
		c <- i
	}
	c := make(chan int, 10)

	for i := 0; i < 10; i++ {
		taskList := TaskList{
			tasks: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		}
		jobList = append(jobList, &taskList)
		println("Create task: ", i)
		go f(c, i, &taskList)
	}

	for i := 0; i < 20; i++ {
		index := <-c
		println("Receive channel: ", index)
		jobList[index].progress = index
		println("Recreate task: ", index, " progress: ", jobList[index].progress)
		go f(c, index, jobList[index])
	}

	println("Bye, World!")
}
