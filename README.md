# Bus
Simplistic UDP Multicast bus system in Go.

Try if i can provide following functionality with a small go package.

```
import fmt
import time
import "github.com/linecker/bus"

func receive(data []byte) {
	fmt.Printf("received ", data);
}

func main(argc int, argv []String) {
	bus.Serve(callback)
	for {
        bus.Send("example message")
        time.Sleep(10 * time.Second)
    }
}
```
