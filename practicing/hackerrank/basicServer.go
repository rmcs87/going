package hackerrank

var serverChan chan in

type in struct {
	arr      int32
	oddChan  chan int32
	evenChan chan int32
}

func Sever() {
	serverChan = make(chan in)

	input := <-serverChan
	values := []int32{}

	tam := input.arr

	for i := 0; i < int(tam); i++ {
		v := <-serverChan
		values = append(values, v.arr)
	}

	for _, v := range values {
		if v%2 != 0 {
			input.oddChan <- v
		}
	}
	for _, v := range values {
		if v%2 == 0 {
			input.evenChan <- v
		}
	}
}
