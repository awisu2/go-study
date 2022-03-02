package main

func callback(f func()) {
	f()
}

func counter(f func(i int)) func(j int) {
	count := 0
	return func(j int) {
		count += j
		f(count)
	}
}
