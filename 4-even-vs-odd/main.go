package main

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range s {
		if v%2 == 0 {
			println(v, ": Is Even")
		} else {
			println(v, ": Is Odd")
		}
	}
}
