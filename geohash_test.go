package geohash_test

import (
	"math"
	"testing"

	"github.com/martishin/geohash"
)

type TestCase struct {
	hash string
	lat  float64
	lng  float64
}

var testCases = []TestCase{
	{"sb54v4xk18jg", 0.497818518, 38.198505253},
	{"00upjeyjb54g", -84.529178182, -174.125057287},
	{"kkfwu0udnhxm", -17.090238388, 14.947853282},
	{"gp2cx4ywjhyj", 86.06108453, -43.628546008},
	{"h0g4tmrp0cut", -85.311745894, 4.459114168},
	{"v471duxnbttv", 57.945830289, 49.349241965},
	{"h78n33z47k3j", -69.203844118, 11.314685805},
	{"gvtw7yer4bhh", 77.073040753, -3.346243298},
	{"0fqwy0pgxxwj", -76.156584583, -136.834730089},
	{"dj53wuppzfrx", 28.411988257, -85.123100792},
	{"rmbn08wvubcz", -11.597823607, 146.281448853},
}

func TestGenerateGeohash(t *testing.T) {
	for _, tc := range testCases {
		result := geohash.GenerateGeohash(tc.lng, tc.lat)
		if result != tc.hash {
			t.Errorf("GenerateGeohash(%f, %f) = %s; expected %s", tc.lng, tc.lat, result, tc.hash)
		}
	}
}

func TestDecodeGeohash(t *testing.T) {
	const tolerance = 1e-6

	for _, tc := range testCases {
		lng, lat, err := geohash.DecodeGeohash(tc.hash)
		if err != nil {
			t.Errorf("DecodeGeohash(%s) returned error: %v", tc.hash, err)
		}
		if math.Abs(lng-tc.lng) > tolerance || math.Abs(lat-tc.lat) > tolerance {
			t.Errorf("DecodeGeohash(%s) = (%f, %f); expected (%f, %f)", tc.hash, lng, lat, tc.lng, tc.lat)
		}
	}
}
