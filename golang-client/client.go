package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

/* signature of Context interface
https://medium.com/nerd-for-tech/learn-intermediate-go-how-does-context-work-1898704c649b
type Context interface {
    Done() <-chan struct{}
    Err() error
    Deadline() (deadline time.Time, ok bool)
    Value(key interface{}) interface{}
}
For now, pay special attention to the Done function.
It returns a channel of empty structs.

*/
//https://www.youtube.com/watch?v=BkzgYfygDy8

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	url := "http://localhost:8080/movies"
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print("Req failed\n")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}

}
