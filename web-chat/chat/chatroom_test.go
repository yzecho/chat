package chat

import (
	"fmt"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	room := NewRoom()
	jack := room.Join("jack")
	go func() {
		for v := range jack.Pipe {
			fmt.Println(v)
		}
	}()

	tom := room.Join("tom")
	go func() {
		for v := range tom.Pipe {
			fmt.Println(v)
		}
	}()
	jack.Say("hello tom")
	tom.Say("hello jack")

	jack.Leave()
	tom.Leave()
	time.Sleep(time.Second)
}
