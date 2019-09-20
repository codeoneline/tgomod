package tgomod

import "fmt"
import "rsc.io/quote"

func SayHello(name string) string {
  return fmt.Sprintf("Hello,%s", name)
}

func SayQuoteHello() string {
  return quote.Hello()
}

