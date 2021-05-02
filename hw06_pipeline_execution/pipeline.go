package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = func(in In) Out {
			bufferChannel := make(Bi, 1)
			go func() {
				defer close(bufferChannel)
				for {
					select {
					case <-done:
						return
					default:
					}

					select {
					case <-done:
						return
					case v, ok := <-in:
						if !ok {
							return
						}
						bufferChannel <- v
					}
				}
			}()
			return bufferChannel
		}(stage(out))
	}
	return out
}
