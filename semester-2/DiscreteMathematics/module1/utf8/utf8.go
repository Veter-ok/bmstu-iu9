package main

import (
	"fmt"
)

func encode(utf32 []rune) []byte {
  var utf8Bytes []byte
	for _, r := range utf32 {
		switch {
		case r <= 0x7F:
			utf8Bytes = append(utf8Bytes, byte(r))
		case r <= 0x7FF:
			utf8Bytes = append(utf8Bytes, byte(0xC0 | ((r >> 6) & 0x1F)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | (r & 0x3F)))
		case r <= 0xFFFF:
			utf8Bytes = append(utf8Bytes, byte(0xE0 | ((r >> 12) & 0x0F)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | ((r >> 6) & 0x3F)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | (r & 0x3F)))
		case r <= 0x10FFFF:
			utf8Bytes = append(utf8Bytes, byte(0xF0 | ((r >> 18) & 0x07)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | ((r >> 12) & 0x3F)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | ((r >> 6) & 0x3F)))
			utf8Bytes = append(utf8Bytes, byte(0x80 | (r & 0x3F)))
		default:
			utf8Bytes = append(utf8Bytes, byte(0xEF), byte(0xBF), byte(0xBD))
		}
	}

	return utf8Bytes
}

func decode(utf8Bytes []byte) []rune {
  var utf32 []rune
  for i := 0; i < len(utf8Bytes); {
      switch {
      case utf8Bytes[i]&0xF0 == 0xF0:
        u := rune(utf8Bytes[i] & 0x07) << 18
        z := rune(utf8Bytes[i+1] & 0x3F) << 12
        y := rune(utf8Bytes[i+2] & 0x3F) << 6
        x := rune(utf8Bytes[i+3] & 0x3F)
        char := u | z | y | x
        utf32 = append(utf32, char)
        i += 4
      case utf8Bytes[i]&0xE0 == 0xE0:
        z := rune(utf8Bytes[i] & 0x0F) << 12
        y := rune(utf8Bytes[i+1] & 0x3F) << 6
        x := rune(utf8Bytes[i+2] & 0x3F)
        char := z | y | x
        utf32 = append(utf32, char)
        i += 3
      case utf8Bytes[i] & 0xC0 == 0xC0:
        y := rune(utf8Bytes[i] & 0x1F) << 6
        x := rune(utf8Bytes[i+1] & 0x3F)
        char := y | x
        utf32 = append(utf32, char)
        i += 2
      default:
        char := rune(utf8Bytes[i])
        utf32 = append(utf32, char)
        i++
      }
  }
  return utf32
}

func main() {
	var message string = "Hello, Мир"
  arr := []rune(message)
  fmt.Println(arr)
  arr2 := []byte(message)
  fmt.Println(arr2)
  utf8 := encode(arr)
  fmt.Println(utf8)
  utf32 := decode(arr2)
  fmt.Println(utf32)
}