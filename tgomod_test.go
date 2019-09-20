package tgomod

import "testing"

func TestSayHello(t *testing.T) {
  want := "Hello, world."
  if got := SayHello(" world."); got != want {
    t.Errorf("SayHello(' world.') = %q, want %q", got, want)
  }
}

func TestSayQuoteHello(t *testing.T) {
  want := "Hello, world."
  if got := SayQuoteHello(); got != want {
    t.Errorf("SayQuoteHello() = %q, want %q", got, want)
  }
}

