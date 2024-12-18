package main

import "sync"

type post struct {
	views  int
	mutext sync.Mutex
}

func (p *post) incrementViews(wg *sync.WaitGroup) {
	defer func() {
		p.mutext.Unlock()
		wg.Done()
	}()

	// lock mutext in which line you need. don't use it at the top of the function.
	p.mutext.Lock()
	p.views++
}

func main() {
	var wg sync.WaitGroup
	p := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go p.incrementViews(&wg)
	}

	wg.Wait()
	println(p.views) // Output: 1
}
