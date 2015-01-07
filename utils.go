package main

import (
  "strconv"
)

func HexToRGB(h string) (uint8, uint8, uint8) {
  if len(h) > 0 && h[0] == '#' {
    h = h[1:]
  }
  if len(h) == 3 {
    h = h[:1] + h[:1] + h[1:2] + h[1:2] + h[2:] + h[2:]
  }
  if len(h) == 6 {
    if rgb, err := strconv.ParseUint(string(h), 16, 32); err == nil {
      return uint8(rgb >> 16), uint8((rgb >> 8) & 0xFF), uint8(rgb & 0xFF)
    }
  }
  return 0, 0, 0
}