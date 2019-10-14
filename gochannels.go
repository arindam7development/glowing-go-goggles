// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into channels from one
// goroutine and receive those values into another
// goroutine.

package main

import "fmt"

func main() {

    // Create a new channel with `make(chan val-type)`.
    // Channels are typed by the values they convey.
    network_ch := make(chan string)
    server_ch := make(chan string)
   

    // _Send_ a value into a channel using the `channel <-`
    // syntax. Here we send `"ping"`  to the `messages`
    // channel we made above, from a new goroutine.
    go func() { network_ch <- "net ping" }()
    go func() { server_ch <- "server ping"}()

    // The `<-channel` syntax _receives_ a value from the
    // channel. Here we'll receive the `"ping"` message
    // we sent above and print it out.
    net_msg := <-network_ch
    server_msg := <-server_ch

    fmt.Println("net_msg : ")
    fmt.Println(net_msg)
    fmt.Println("server_msg :")
    fmt.Println(server_msg)
}
