package main

func main() {
	cfg := newDefaultConfig()
	e := NewServer(cfg)
	e.Start(":8000")
}
