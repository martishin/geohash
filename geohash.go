package geohash

import "strings"

const (
	base32Codes = "0123456789bcdefghjkmnpqrstuvwxyz"
	// Each base32 character represents exactly 5 bits of a sequence.
	bitsPerChar = 5
	// Base32 mask 0x1F, equivalent to 00011111 in binary.
	bitsMask  = (1 << bitsPerChar) - 1
	coordBits = 30
	totalBits = 60
)

func GenerateGeohash(x, y float64) string {
	var xBits, yBits uint32

	xs := []float64{-180.0, 0.0, 180.0}
	ys := []float64{-90.0, 0.0, 90.0}

	for i := coordBits - 1; i >= 0; i-- {
		processCoordinate(x >= xs[1], &xBits, xs, i)
		processCoordinate(y >= ys[1], &yBits, ys, i)
	}

	resultBits := interleaveBits(xBits, yBits)

	return encodeGeohash(resultBits)
}

func processCoordinate(rightOfMedian bool, bits *uint32, coords []float64, bitPosition int) {
	if rightOfMedian {
		*bits |= 1 << bitPosition
		coords[0] = coords[1]
	} else {
		coords[2] = coords[1]
	}

	coords[1] = (coords[0] + coords[2]) / 2.0
}

func interleaveBits(xBits, yBits uint32) uint64 {
	var resultBits uint64

	for i := coordBits - 1; i >= 0; i-- {
		resultBits |= uint64((xBits>>i)&1) << (2*i + 1)
		resultBits |= uint64((yBits>>i)&1) << (2 * i)
	}

	return resultBits
}

func encodeGeohash(resultBits uint64) string {
	var geohash strings.Builder
	geohash.Grow(totalBits / bitsPerChar)

	for i := bitsPerChar; i <= totalBits; i += bitsPerChar {
		pos := (resultBits >> (totalBits - i)) & bitsMask
		geohash.WriteByte(base32Codes[pos])
	}

	return geohash.String()
}
