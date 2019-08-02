package main

import (
    "sync"

    "github.com/vietwow/user-management-grpc/client"
    "github.com/vietwow/user-management-grpc/server"
)

func main() {
    // Create the waitgroup and add the total number of goroutines we're going to use
    var wg sync.WaitGroup
    wg.Add(2)

    go client.StartClient(&wg)
    go server.StartServer(&wg)

    wg.Wait()
}
