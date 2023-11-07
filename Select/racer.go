package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTime = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTime)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("time out waiting for %s and %s", a, b)
	}
}

func ping(URL string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(URL)
		close(ch)
	}()
	return ch
}
