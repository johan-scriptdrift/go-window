package window

func Slide[T any](data []T, size int) <-chan []T {
	windowChan := make(chan []T, 1)
	go func() {
		defer close(windowChan)
		if size <= 0 {
			return
		}
		if len(data) < size {
			return
		}
		for i := 0; i+size <= len(data); i++ {
			win := make([]T, size)
			copy(win, data[i:i+size])
			windowChan <- win
		}
	}()

	return windowChan
}
