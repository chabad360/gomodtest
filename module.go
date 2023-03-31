package gomodtest

func Add(a int, b ...int) int {
  for _, v := range b {
    a += v
  }
  return a
}
