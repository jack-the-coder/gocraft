package main

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	opensimplex "github.com/ojrac/opensimplex-go"
)

var (
	sim = opensimplex.NewWithSeed(0)
)

// func customNoise (x, y float32) float32 {
//         return abs(float32(10*sim.Eval2(float64(x), float64(y))))
// }
// 
// func simplexHeight(x, y float32) float32 {
//         e := 1 * customNoise(1*x, 1*y) + 0.5 * customNoise(2*x, 2*y) + 0.25 * customNoise(4*x,4*y)
//         print("\n", math.Pow(float64(e), 1.5))
//         return float32(math.Pow(float64(e), 1.5))
// }

func pow(x, y float32) float32 {
        return float32(math.Pow(float64(x), float64(y)))
}

func abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

func sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func radian(angle float32) float32 {
	return mgl32.DegToRad(angle)
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func mix(a, b, factor float32) float32 {
	return a*(1-factor) + factor*b
}

func noise2(x, y float32, octaves int, persistence, lacunarity float32) float32 {
	var (
		freq  float32 = 1
		amp   float32 = 1
		max   float32 = 1
		total         = sim.Eval2(float64(x), float64(y))
	)
	for i := 0; i < octaves; i++ {
		freq *= lacunarity
		amp *= persistence
		max += amp
		total += sim.Eval2(float64(x*freq), float64(y*freq)) * float64(amp)
	}
	return (1 + float32(total)/max) / 2
}

func noise3(x, y, z float32, octaves int, persistence, lacunarity float32) float32 {
	var (
		freq  float32 = 1
		amp   float32 = 1
		max   float32 = 1
		total         = sim.Eval3(float64(x), float64(y), float64(z))
	)
	for i := 0; i < octaves; i++ {
		freq *= lacunarity
		amp *= persistence
		max += amp
		total += sim.Eval3(float64(x*freq), float64(y*freq), float64(z*freq)) * float64(amp)
	}
	return (1 + float32(total)/max) / 2
}
