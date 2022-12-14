package panichand

import "log"

func Divide(a int, b int) int {
	res := a / b

	log.Printf("[Divide] %d / %d => %d", a, b, res)
	return a / b
}

func defereable(i string) {
	log.Printf("[defereable] %s", i)
}

type Block struct {
	Flag int
}

func Run(wantFail bool) (int, *Block, int) {

	var block1 *Block = &Block{}
	state := 0

	defer defereable("beta")

	defer func() {

		if r := recover(); r != nil {
			log.Printf("[recover] defer() that are defined after a recover will not be invoked") // d1 d2 d3 d4 dR d6 d7
		}
	}()
	i := 10

	var inferior int

	if wantFail {
		inferior = 0

	} else {
		inferior = 1
	}
	for i > inferior {

		i -= 1
		Divide(10-i, i)
	}

	defer func() {
		log.Printf("[inline-defer state init] primitive: %d pointer val: %d %p", state, block1.Flag, block1)

		state += 1
		block1.Flag += 1

	}()

	defer func() {
		log.Printf("[inline-defer state after] primitive: %d pointer val: %d %p", state, block1.Flag, block1)
	}()

	defer defereable("omega")

	defer defereable("charlie. should be executed last")

	return state, block1, block1.Flag
}
