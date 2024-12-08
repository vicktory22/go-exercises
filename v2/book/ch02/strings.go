package main

import (
	"fmt"
	"strings"
)

func main() {
  fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAlis")) 
  fmt.Printf("EqualFold: %v\n", strings.EqualFold("Mihalis", "MIHAli"))

  fmt.Printf("Index: %v\n", strings.Index("Mihalis", "ha"))
  fmt.Printf("Index: %v\n", strings.Index("Mihalis", "Ha"))

  fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "Mi"))
  fmt.Printf("Prefix: %v\n", strings.HasPrefix("Mihalis", "mi"))

  fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "is"))
  fmt.Printf("Suffix: %v\n", strings.HasSuffix("Mihalis", "IS"))

  t := strings.Fields("This is a string")
  fmt.Printf("Fields: %v\n", len(t))

  t = strings.Fields("ThisIs a\tstring!")
  fmt.Printf("Fields: %v\n", len(t))
}
