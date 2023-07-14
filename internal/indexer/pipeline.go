package indexer

type Processor func(any) any

func Pipeline(input <-chan any, processor Processor) <-chan any {
	out := make(chan any)
	go func() {
		for in := range input {
			out <- processor(in)
		}
		close(out)
	}()
	return out
}
