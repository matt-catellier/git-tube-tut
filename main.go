package main

import (
  "fmt"
  "math/rand"
  web "./web"
)



func goRandom() {
 // provides pseudo random number, meaning same sequence each time
 // to vary output, need to provide a seed
 fmt.Println(rand.Intn(100))
 rand.Seed(42)
 fmt.Println(rand.Intn(100))
 // create new rand instance
 r := rand.New(rand.NewSource(46))
 fmt.Println(r.Intn(100))
}

func goTypes() {
  // cannot change typ during runtime
  // NOTE: when using := shorthand for paramenter values
  add := func(x, y float32) float32 {
    return x +y
  }
  fmt.Println(add(5.2, 0.8))
  const num1, num2 int = 12, 3
  fmt.Println(add(float32(num1), float32(num2)))
}

func goPointers() {
  fmt.Println("=== go pointers ===")

  x := 15
  a := &x // memory address, not the value

  fmt.Println(a)
  fmt.Println(*a) // access value by pointer

  *a = 5 // overwrite value in memory
  fmt.Println(x)

  *a = *a * *a
  fmt.Println(x)
  fmt.Println(*a)
}

// struct lesson

const u_16_max uint16 = 65535
const kmh_multiple float64 = 1.6778

type car struct {
  gas_pedal uint16 // 0 - 65535
  brake_pedal uint16
  steering_wheel int16
  top_speed_kmh float64
}

// value reciever - pure, perform calulations
// makes a copy of the value passed in, taking up new space in memory
func (c car) kmh() float64 { // c car is what assocaites it with a struct
  return float64(c.gas_pedal) * (c.top_speed_kmh/float64(u_16_max))
}

// pointer receiver - performs sideeffects
func (c *car) new_top_speed(speed float64){
  c.top_speed_kmh = speed
}

func newer_top_speed(c car, speed float64) car {
  c.top_speed_kmh = speed
  return c
}

func goStruct() {
  a_car := car{2234, 0, 12562, 225.0}

  fmt.Println(a_car.kmh())

  a_car.new_top_speed(500)
  fmt.Println(a_car.kmh())

  a_car = newer_top_speed(a_car, 600)

  fmt.Println(a_car.kmh())
}

func main() {
//  goRandom()
//  goTypes()
//  goPointers()
  web.Run()
//  goStruct()
}
