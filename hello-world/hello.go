package main

import (
	"fmt"
)

const (
  spanish = "spanish"
  french = "french"
  englishHelloPrefix = "Hello, "
  spanishHelloPrefix = "Hola, "
  frenchHelloPrefix = "Bonjour, "
)

func Hello(w string, language string) string {
  if (w == "") {
    w = "World"
  }
  return greetingPrefix(language) + w
}

func greetingPrefix(language string) string {
  switch language {
  case french:
    return frenchHelloPrefix
  case spanish:
    return spanishHelloPrefix
  default:
    return englishHelloPrefix
  }
}

func main() {
  fmt.Println(Hello("Captain", ""))
}

