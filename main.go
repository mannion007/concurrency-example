package main

import (
    "log"
    "fmt"
    "sync"
    "time"
    "net/http"
)

func benchmark(wg *sync.WaitGroup, durations chan<- float64, url string) {
    start := time.Now()

    wg.Add(1)

    go func() {
        _, err := http.Get(url)
    
        if err != nil {
            log.Printf("%s took too long to respond, giving up", url)
        }

        durations<- time.Since(start).Seconds()

        wg.Done()

    }()
}

func main() {
    urls := make([]string, 0)
    urls = append(urls, "https://www.google.com")
    urls = append(urls, "https://www.youtube.com")
    urls = append(urls, "https://www.sainsburys.co.uk")
    urls = append(urls, "https://www.bbc.co.uk")

    durations := make(chan float64, len(urls))
    defer close(durations)

    var wg sync.WaitGroup

    for i:=0; i<len(urls); i++{
        benchmark(&wg, durations, urls[i])
    }

    for i:=0; i<len(urls); i++{
        fmt.Println(fmt.Sprintf("%s took %f seconds to respond", urls[i], <-durations))
    }

    wg.Wait()
}