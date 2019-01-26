package main

func main() {
	// dataStream := make(<-chan interface{}) // compilable

	// close(dataStream)
	//	<-dataStream

	var ch chan interface{}
	<-ch // Allow receive from nil chanel

	// go run ./strange/main.go
	// fatal error: all goroutines are asleep - deadlock!
	//
	// goroutine 1 [chan receive (nil chan)]:
	// main.main()
	// 	/Users/budougumi0617/go/src/github.com/budougumi0617/cig/ch03/strange/main.go:9 +0x29
}
