package gomobilebug

import (
	"fmt"

	"github.com/google/netstack/tcpip"
)

// ExampleMethod prints the value of a variable defined in github.com/google/netstack/tcpip
func ExampleMethod() {
	fmt.Println(tcpip.ErrUnknownProtocol)
}
