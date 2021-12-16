package cracker

import (
	"sync"
)

type Result interface{}

type Resource interface{}

type Argument interface{}

type Proc func(resource Resource, arg Argument) (Result, error)

func ProcessBatch(cmds []Command, resources []Resource) []Command {
	var results []positionedCommand
	cq := make(chan positionedCommand, len(cmds))
	rq := make(chan positionedCommand, len(cmds))
	var wg sync.WaitGroup

	pcmds := newCommandArray(cmds)
	wg.Add(len(pcmds))

	for _, cmd := range pcmds {
		cq <- cmd
	}

	// wait till all commands processed
	go func() {
		wg.Wait()
		close(rq)
	}()

	// process commands, reschedule if re-try required
	for _, resource := range resources {
		go func(resource Resource) {
			for cmd := range cq {
				cmd.run(resource)
				wg.Done()
				if cmd.shouldTryAgain() {
					wg.Add(1)
					cq <- cmd
				} else {
					rq <- cmd
				}
			}
		}(resource)
	}

	for val := range rq {
		results = append(results, val)
	}

	return commandArray(results).toCommandSlice()
}
