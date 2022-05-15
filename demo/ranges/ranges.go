package main

func main() {
	slice := []string{"hello", "world"}
	for _, s := range slice {
		println(s)
		for _, c := range s {
			println(c)
		}
	}
}
