package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
)

type Vector struct {
	vector [2]float64
}

func NewVector() Vector {
	return Vector{vector: [2]float64{0, 0}}
}

func NewVectorFromValues(x, y float64) Vector {
	return Vector{vector: [2]float64{x, y}}
}

func (v Vector) X() float64 {
	return v.vector[0]
}

func (v *Vector) SetX(x float64) {
	v.vector[0] = x
}

func (v Vector) Y() float64 {
	return v.vector[1]
}

func (v *Vector) SetY(y float64) {
	v.vector[1] = y
}

func (v Vector) Clone() Vector {
	return Vector{vector: [2]float64{v.vector[0], v.vector[1]}}
}

func (v Vector) String() string {
	return fmt.Sprintf("%s,%s", strconv.FormatFloat(v.vector[0], 'f', 5, 64), strconv.FormatFloat(v.vector[1], 'f', 5, 64))
}

func (v Vector) Normalized() Vector {
	distance := math.Sqrt(v.vector[0]*v.vector[0] + v.vector[1]*v.vector[1])
	return Vector{vector: [2]float64{v.vector[0] / distance, v.vector[1] / distance}}
}

func (v Vector) Dot(w Vector) float64 {
	return v.vector[0]*w.vector[0] + v.vector[1]*w.vector[1]
}

func (v Vector) DistanceTo(w Vector) float64 {
	return Distance(v, w)
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.vector[0]*v.vector[0] + v.vector[1]*v.vector[1])
}

func Distance(a, b Vector) float64 {
	dist := math.Sqrt((a.vector[0]-b.vector[0])*(a.vector[0]-b.vector[0]) + (a.vector[1]-b.vector[1])*(a.vector[1]-b.vector[1]))
	return dist
}

func (v Vector) Abs() Vector {
	return Vector{vector: [2]float64{math.Abs(v.vector[0]), math.Abs(v.vector[1])}}
}

func (v Vector) MulScaler(s float64) Vector {
	return NewVectorFromValues(v.X()*s, v.Y()*s)
}

func (v Vector) Reflect(normal Vector) Vector {
	dot := v.Dot(normal) * 2
	reflected := v.Sub(normal.MulScalar(dot))
	return reflected
}

func (v Vector) Lerp(b Vector, t float64) Vector {
	return v.Add(b.Sub(v).MulScalar(t))
}

func Lerp(a, b Vector, t float64) Vector {
	return a.Add(b.Sub(a).MulScalar(t))
}

func (v Vector) Add(w Vector) Vector {
	return Vector{vector: [2]float64{v.vector[0] + w.vector[0], v.vector[1] + w.vector[1]}}
}

func (v Vector) Sub(w Vector) Vector {
	return Vector{vector: [2]float64{v.vector[0] - w.vector[0], v.vector[1] - w.vector[1]}}
}

func (v Vector) Mul(w Vector) Vector {
	return Vector{vector: [2]float64{v.vector[0] * w.vector[0], v.vector[1] * w.vector[1]}}
}

func (v Vector) MulScalar(m float64) Vector {
	return Vector{vector: [2]float64{v.vector[0] * m, v.vector[1] * m}}
}

func (v Vector) DivScalar(d float64) Vector {
	return Vector{vector: [2]float64{v.vector[0] / d, v.vector[1] / d}}
}

func (v Vector) ToPosition() fyne.Position {
	return fyne.Position{
		X: float32(v.vector[0]),
		Y: float32(v.vector[1]),
	}
}
