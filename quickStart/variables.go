package main

    // 11 целочисленных типов: int8, int16, int32, int64, uint8, uint16, uint32, uint64
    //      int, uint, byte (uint8), rune (uint32), uintptr
    // 2 типа с плавающей точкой: float32, float64
    // 3 типа с неограниченной точностью (пакет big): big.Int, big.Rat, big.Float
    // 2 типа комплексных чисел: complex64, complex128
    // 1 логический тип: bool
    // 1 строка: string

func main() {
    i := 4
    f := 4.5
    l := false
    s := "string"
    println(i)
    println(f)
    println(l)
    println(s)
}