package geohash

import (
	"fmt"
	"strings"
)

const (
	base32Codes = "0123456789bcdefghjkmnpqrstuvwxyz"
	// Each base32 character represents exactly 5 bits of a sequence.
	bitsPerChar = 5
	// Base32 mask 0x1F, equivalent to 00011111 in binary.
	bitsMask     = (1 << bitsPerChar) - 1
	oneCoordBits = 30
	totalBits    = 60
)

// GenerateGeohash encodes the (longitude, latitude) into a 12-character geohash.
func GenerateGeohash(lng, lat float64) string {
	var lngBits, latBits uint32

	lngLimits := []float64{-180.0, 0.0, 180.0}
	latLimits := []float64{-90.0, 0.0, 90.0}

	for i := oneCoordBits - 1; i >= 0; i-- {
		processCoordinate(lng >= lngLimits[1], &lngBits, lngLimits, i)
		processCoordinate(lat >= latLimits[1], &latBits, latLimits, i)
	}

	resultBits := interleaveBits(lngBits, latBits)

	return encodeGeohash(resultBits)
}

func processCoordinate(rightOfMedian bool, coordBits *uint32, coordLimits []float64, bitPosition int) {
	if rightOfMedian {
		*coordBits |= 1 << bitPosition
		coordLimits[0] = coordLimits[1]
	} else {
		coordLimits[2] = coordLimits[1]
	}

	coordLimits[1] = (coordLimits[0] + coordLimits[2]) / 2.0
}

func interleaveBits(lngBits, latBits uint32) uint64 {
	var resultBits uint64

	for i := oneCoordBits - 1; i >= 0; i-- {
		resultBits |= uint64((lngBits>>i)&1) << (2*i + 1)
		resultBits |= uint64((latBits>>i)&1) << (2 * i)
	}

	return resultBits
}

func encodeGeohash(resultBits uint64) string {
	var geohash strings.Builder
	geohash.Grow(totalBits / bitsPerChar)

	for i := bitsPerChar; i <= totalBits; i += bitsPerChar {
		idx := (resultBits >> (totalBits - i)) & bitsMask
		geohash.WriteByte(base32Codes[idx])
	}

	return geohash.String()
}

// DecodeGeohash decodes a 12-character geohash string back to the (longitude, latitude)
func DecodeGeohash(hash string) (lng, lat float64, err error) {
	if len(hash) != totalBits/bitsPerChar {
		return 0, 0, fmt.Errorf("invalid geohash length: got %d, expected %d", len(hash), totalBits/bitsPerChar)
	}
	var inputBits uint64
	for i := 0; i < len(hash); i++ {
		idx := strings.IndexByte(base32Codes, hash[i])
		if idx < 0 {
			return 0, 0, fmt.Errorf("invalid character in geohash: %c", hash[i])
		}
		inputBits = (inputBits << bitsPerChar) | uint64(idx)
	}

	lngBits, latBits := deInterleaveBits(inputBits)

	lng = decodeCoordinate(-180.0, 180.0, lngBits)
	lat = decodeCoordinate(-90.0, 90.0, latBits)

	return lng, lat, nil
}

func deInterleaveBits(resultBits uint64) (uint32, uint32) {
	var lngBits, latBits uint32
	for i := oneCoordBits - 1; i >= 0; i-- {
		lngBits |= uint32((resultBits>>(2*i+1))&1) << i
		latBits |= uint32((resultBits>>(2*i))&1) << i
	}

	return lngBits, latBits
}

func decodeCoordinate(coordMin float64, coordMax float64, coordBits uint32) float64 {
	// Decode coordinate using binary subdivision.
	for i := oneCoordBits - 1; i >= 0; i-- {
		coordMid := (coordMin + coordMax) / 2
		if ((coordBits >> i) & 1) == 1 {
			coordMin = coordMid
		} else {
			coordMax = coordMid
		}
	}
	return (coordMin + coordMax) / 2
}
