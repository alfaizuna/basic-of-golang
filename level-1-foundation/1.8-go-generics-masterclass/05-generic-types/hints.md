# Hints

### Hint 1
Deklarasikan `type Pair[T, U any] struct { First T; Second U }`.

### Hint 2
Tulis fungsi: `func NewPair[T, U any](first T, second U) Pair[T, U] { return Pair[T, U]{First: first, Second: second} }`.

### Hint 3
Di main: panggil `p := NewPair("HTTP", 200)`.
