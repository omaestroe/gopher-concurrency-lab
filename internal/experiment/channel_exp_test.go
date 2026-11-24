package experiment

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannelCommunication(t *testing.T) {
	ch := make(chan string)

	go SendMessage(ch)

	message := <-ch

	assert.Equal(t, "Hello from goroutine", message)
}

func TestTwoChannelCommunication(t *testing.T) {
	ch := make(chan string)
	ch2 := make(chan string)

	go SendMessage(ch)

	message := <-ch

	go func() {
		ch2 <- message
	}()

	message2 := <-ch2
	assert.Equal(t, "Hello from goroutine", message2)
}

func TestTwoChannelCommunicationBuffered(t *testing.T) {
	ch := make(chan string)
	ch2 := make(chan string, 1)

	go SendMessage(ch)

	message := <-ch

	ch2 <- message
	message2 := <-ch2

	assert.Equal(t, "Hello from goroutine", message2)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan string, 2)

	ch <- "message 1"
	ch <- "message 2"

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		message := <-ch

		t.Logf("Received: %s", message)
	}()

	ch <- "message 3"

	//wg.Wait()

	t.Log("Third message sent successfully")
}
