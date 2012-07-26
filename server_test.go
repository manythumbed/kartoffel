package kartoffel

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	statusChannel := make(chan ServerStatus)

	go func() {
		fmt.Println("Running listen")
		err := Listen(0, statusChannel)
		fmt.Println(err)
	}()

	fmt.Println("Listener started")
	status := <-statusChannel
	if status.State != Accepting {
		t.Errorf("Server not ready")
	}
}
