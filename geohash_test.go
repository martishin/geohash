package geohash_test

import (
	"testing"

	"github.com/martishin/geohash"
)

type TestCase struct {
	expectedInt uint64
	expectedStr string
	lat         float64
	lng         float64
}

func TestGenerateGeohash(t *testing.T) {
	var testcases = []TestCase{
		{0xc28a4d93b20a22f8, "sb54v4xk18jg", 0.497818518, 38.198505253},
		{0x003558b7d15148f1, "00upjeyjb54g", -84.529178182, -174.125057287},
		{0x949dcd034ca43b30, "kkfwu0udnhxm", -17.090238388, 14.947853282},
		{0x7d44be93dc8c3d1f, "gp2cx4ywjhyj", 86.06108453, -43.628546008},
		{0x801e4ccef502f590, "h0g4tmrp0cut", -85.311745894, 4.459114168},
		{0xd90e166bb45673b6, "v471duxnbttv", 57.945830289, 49.349241965},
		{0x81d1418fe43c871f, "h78n33z47k3j", -69.203844118, 11.314685805},
		{0x7ef3c3f9b722a109, "gvtw7yer4bhh", 77.073040753, -3.346243298},
		{0x03adcf02afef7916, "0fqwy0pgxxwj", -76.156584583, -136.834730089},
		{0x644a3e6ab5fbafd6, "dj53wuppzfrx", 28.411988257, -85.123100792},
		{0xbcd540239bd297fb, "rmbn08wvubcz", -11.597823607, 146.281448853},
	}

	for _, tc := range testcases {
		result := geohash.GenerateGeohash(tc.lng, tc.lat)
		if result != tc.expectedStr {
			t.Errorf("GenerateGeohash(%f, %f) = %s; expected %s", tc.lng, tc.lat, result, tc.expectedStr)
		}
	}
}
