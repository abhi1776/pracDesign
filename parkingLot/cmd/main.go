package main

import (
	"github.com/abhi1776/pracDesign/parkingLot/service"
)

func main() {
	service.ParkingLotService()
}

// func main() {
// 	// ch := make(chan struct{}, 2)
// 	var wg, wg1 sync.WaitGroup

// 	people := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
// 	for i := 0; i < len(people); i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("Process:", people[i])
// 			time.Sleep(1 * time.Second)
// 		}()
// 		if i+1 >= len(people) {
// 			break
// 		}
// 		wg1.Add(1)
// 		go func() {
// 			defer wg1.Done()
// 			fmt.Println("Process:", people[i+1])
// 			time.Sleep(1 * time.Second)
// 		}()
// 		i++

// 	}
// 	wg.Wait()
// 	wg1.Wait()
// 	// close(ch)
// }

// func main() {
// 	ch := make(chan struct{})
// 	go func() {
// 		people := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
// 		for i := 0; i < len(people); i++ {
// 			fmt.Println("Process:", people[i])
// 			time.Sleep(100 * time.Millisecond)
// 			ch <- struct{}{}
// 		}
// 		close(ch)
// 	}()

// 	//iteration over messages in the channel, stops only when the channel is closed using close(ch)
// 	for message := range ch {
// 		fmt.Println(message)
// 	}

// 	//another way of iterating
// 	// for {
// 	// 	message, isChannelOpen := <-ch
// 	// 	if !isChannelOpen {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(message)
// 	// }

// }

// When a channel is closed, no more values can be sent on it,
// but you can still receive from it until the channel is empty.
// func main() {
// 	ninja1 := make(chan struct{})
// 	close(ninja1)
// 	ninja2 := make(chan struct{})
// 	close(ninja2)
// 	c1, c2 := 0, 0

// 	//Receiving from a closed channel always yields the zero value for the channel's type and does not block.
// 	for i := 0; i < 1000; i++ {

// 		select {
// 		case <-ninja1:
// 			c1++
// 		case <-ninja2:
// 			c2++
// 		case <-time.After(time.Duration(1) * time.Second): // set a timeout
// 			fmt.Println("timed out")
// 		}

// 	}
// 	fmt.Println(c1, c2)
// }

//var rwMutex sync.RWMutex

// func main() {
// 	go read()
// 	go write()
// 	go read()
// 	go read()
// 	go read()
// 	time.Sleep(10 * time.Second)
// }

// func read() {
// 	rwMutex.RLock()
// 	defer rwMutex.RUnlock()
// 	fmt.Println("read locking")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("read unlocking")
// }

// func write() {
// 	rwMutex.Lock()
// 	defer rwMutex.Unlock()
// 	fmt.Println("write locking")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("write unlocking")
// }

// func main() {
//     var bufferPool = sync.Pool{
//         New: func() interface{} {
//             // Define what to do when the pool is empty
//             fmt.Println("Allocating new buffer")
//             return make([]byte, 1024) // Example: 1KB buffer
//         },
//     }

//     // Retrieve a buffer from the pool
//     buffer := bufferPool.Get().([]byte)

//     // Use the buffer (for demonstration purposes, we just print its length)
//     fmt.Println("Buffer length:", len(buffer))

//     // Return the buffer to the pool
//     bufferPool.Put(buffer)

//     // Retrieve another buffer from the pool (this should reuse the buffer returned earlier)
//     buffer2 := bufferPool.Get().([]byte)
//     fmt.Println("Buffer2 length:", len(buffer2))
// }

// func main() {
// 	ch1 := read(1)
// 	ch2 := read(2)
// 	chM := merge([]chan int{ch1, ch2})

// 	for message := range chM {
// 		fmt.Println(message)
// 	}
// }

// func read(a int) (ch chan int) {
// 	ch = make(chan int)
// 	go func() {
// 		for i := 0; i < 7; i++ {
// 			time.Sleep(300 * time.Millisecond)
// 			ch <- a * (i + 1)
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }

// func merge(ch []chan int) (chM chan int) {
// 	chM = make(chan int)
// 	var wg sync.WaitGroup
// 	for i := 0; i < 2; i++ {
// 		wg.Add(1)
// 		go func() {
// 			for msg := range ch[i] {
// 				chM <- msg
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(chM)
// 	}()

// 	return chM
// }

// func main() {
// 	a := "aaaaabccdddddddddddddddde"
// 	encoded := encode(a)
// 	fmt.Println(encoded)
// 	decoded := decode(encoded)
// 	fmt.Println(decoded)
// }

// func encode(a string) (encoded string) {
// 	if a == "" {
// 		return ""
// 	}
// 	count := 1
// 	for i := 1; i < len(a); i++ {
// 		if a[i] == a[i-1] {
// 			count++
// 		} else {
// 			encoded += string(a[i-1]) + strconv.Itoa(count)
// 			count = 1
// 		}
// 	}
// 	encoded += string(a[len(a)-1]) + strconv.Itoa(count)
// 	return
// }

// func decode(a string) (decoded string) {
// 	l := len(a)
// 	if l < 2 {
// 		return
// 	}
// 	i := 0
// 	for i < l {
// 		count := 0
// 		j := i + 1
// 		for j < l && a[j] <= '9' && a[j] >= '0' {
// 			count = count*10 + int(a[j]-'0')
// 			j++
// 		}
// 		for k := 0; k < count; k++ {
// 			decoded += string(a[i])
// 		}
// 		i = j
// 	}
// 	return
// }
