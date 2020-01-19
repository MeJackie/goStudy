package main

import (
	"fmt"
	"io"
	"os"
	"time"
	"flag"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	var s io.Writer
	s = os.Stdout
	_, _ = s.Write([]byte("hello"))
	flag.Parse()
	fmt.Printf("sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// 实现flag.Value接口，实现自定义符号
type Celsius float64    //摄氏度
type Fahrenheit float64  //华氏度

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var uint string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &uint)
	switch uint {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("Invalid temperature %v", s)
}

