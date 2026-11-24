package experiment

func SendMessage(ch chan string) {
	ch <- "Hello from goroutine"
}
