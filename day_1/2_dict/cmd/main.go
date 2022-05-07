package main

import (
	"context"
	"fmt"
	dict "go_dance/day_1/2_dict"
	"go_dance/day_1/2_dict/source"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	s := dict.NewSource(&source.Caiyun{}, &source.Sougou{})
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	responseStream, errStream := trans(ctx, s, word)
	for response := range responseStream {
		if err := <-errStream; err != nil {
			log.Println(err)
			continue
		}
		name, res := response[len(response)-1], response[:len(response)-1]
		fmt.Println(name, ":")
		for i := range res {
			fmt.Println(res[i])
		}
	}
}

func trans(ctx context.Context, s *dict.Source, word string) (<-chan []string, <-chan error) {
	resultStream := make(chan []string)
	errStream := make(chan error)
	go func() {
		defer close(resultStream)
		defer close(errStream)
		wg := &sync.WaitGroup{}
		for _, t := range s.List() {
			wg.Add(1)
			go func(d dict.Dictionary) {
				defer wg.Done()
				response, err := d.Transform(word, "en", "zh")
				response = append(response, d.Name())
				var out1, out2 = resultStream, errStream
				for i := 0; i < 2; i++ {
					select {
					case <-ctx.Done():
						return
					case out1 <- response:
						out1 = nil
					case out2 <- err:
						out2 = nil
					}
				}
			}(t)
		}
		wg.Wait()
	}()
	return resultStream, errStream
}
