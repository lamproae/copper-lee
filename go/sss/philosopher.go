package main

type Philosopher struct {
	leftHand  bool
	rightHand bool
	status    int
	name      string
}

func main() {
	philosophers := [...]Philospher{"Kant", "Turing", "Descartes", "Kiekegaard", "Wittgenstein"}

	evaluate := func() {
		for {
			select {
			case <-forkUp:
				//philosophers think!
			case <-forkDown:
				//next philospher eats in round robin
			}
		}
	}
}
