package ckks

import (
	"math"
)

type BootParams struct {
	Parameters

	SinType   SinType // Chose betwenn [Sin(2*pi*x)] or [cos(2*pi*x/r) with double angle formula]
	SinRange  uint64  // K parameter (interpolation in the range -K to K)
	SinDeg    uint64  // Degree of the interpolation
	SinRescal uint64  // Number of rescale and double angle formula (only applies for cos)
	BabySplit uint64  // L parameter of the Baby-step giant-step algorithm (the smallest the more precision but the more non-scalr multiplications)

	CtSLevel []uint64 // Level of the Coeffs To Slots
	StCLevel []uint64 // Level of the Slots To Coeffs

	SinDepth uint64 // Automatically set
}

type SinType uint64

const (
	Sin = SinType(0)
	Cos = SinType(1)
)

func (b *BootParams) Gen() {

	if b.SinType == SinType(Sin) && b.SinRescal != 0 {
		panic("BootParams: cannot use double angle formul for SinType = Sin -> must use SinType = Cos")
	}

	b.SinDepth = uint64(math.Ceil(math.Log2(float64(b.SinDeg))) + float64(b.SinRescal) + 1)

	b.Parameters.Gen()
}

var BootstrappParams = []*BootParams{

	// 1390 Sin - 550
	{Parameters: Parameters{
		LogN:     16,
		LogSlots: 15,
		LogModuli: LogModuli{
			LogQi: []uint64{55, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 56, 28, 55, 55, 55, 55, 55, 55, 55, 55, 53, 53, 53},
			LogPi: []uint64{55, 55, 55, 55},
		},
		Scale: 1 << 45,
		Sigma: 3.2,
	},
		SinType:   Sin,
		SinRange:  15,
		SinDeg:    124,
		SinRescal: 0,
		CtSLevel:  []uint64{24, 23, 22},
		StCLevel:  []uint64{13, 12, 12},
	},

	// 1399 Sin - 505
	// 15 : 19.55
	{Parameters: Parameters{
		LogN:     16,
		LogSlots: 15,
		LogModuli: LogModuli{
			LogQi: []uint64{55, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 56, 28, 55, 55, 55, 55, 55, 55, 55, 55, 53, 53, 53},
			LogPi: []uint64{55, 55, 55, 55},
		},
		Scale: 1 << 45,
		Sigma: 3.2,
	},
		SinType:   Cos,
		SinRange:  15,
		SinDeg:    38,
		SinRescal: 2,
		CtSLevel:  []uint64{23, 22, 21},
		StCLevel:  []uint64{12, 12, 11},
	},

	// 1439 Sin - 505
	{Parameters: Parameters{
		LogN:     16,
		LogSlots: 15,
		LogModuli: LogModuli{
			LogQi: []uint64{55, 60, 60, 60, 60, 60, 60, 60, 60, 60, 55, 55, 55, 55, 55, 55, 55, 55, 53, 53, 53},
			LogPi: []uint64{61, 61, 61, 61},
		},
		Scale: 1 << 30,
		Sigma: 3.2,
	},
		SinType:   Sin,
		SinRange:  15,
		SinDeg:    127,
		SinRescal: 0,
		CtSLevel:  []uint64{20, 19, 18},
		StCLevel:  []uint64{9, 9, 8},
	},

	// 1439 cos - 505
	// 10 : 21.44
	// 11 : 21.00
	// 12 : 20.46
	// 13 : 19.89
	// 14 : 19.31
	// 15 : 19.51
	{Parameters: Parameters{
		LogN:     16,
		LogSlots: 15,
		LogModuli: LogModuli{
			LogQi: []uint64{55, 60, 60, 60, 60, 60, 60, 60, 60, 60, 55, 55, 55, 55, 55, 55, 55, 55, 53, 53, 53},
			LogPi: []uint64{61, 61, 61, 61},
		},
		Scale: 1 << 30,
		Sigma: 3.2,
	},
		SinType:   Cos,
		SinRange:  15,
		SinDeg:    38,
		SinRescal: 2,
		CtSLevel:  []uint64{20, 19, 18},
		StCLevel:  []uint64{9, 9, 8},
	},
}
