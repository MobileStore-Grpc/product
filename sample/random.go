package sample

import (
	"math/rand"
	"time"

	"github.com/MobileStore-Grpc/product/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomStringFromSet(names ...string) string {
	n := len(names)
	if n == 0 {
		return ""
	}
	return names[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomMobileBrand() string {
	return randomStringFromSet("Apple", "Oneplus", "Samsung")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomMobileName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Apple 12 Plus", "Apple 12 Plus Pro")
	case "Oneplus":
		return randomStringFromSet("Oneplus 8", "Oneplus 8T")
	default:
		return randomStringFromSet("Samsung Galaxy S21", "Samsung Galaxy S21 Ultra")
	}
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9",
			"Core i7",
			"Core i5",
			"Core i3",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO",
		"Ryzen 5 PRO",
		"Ryzon 3 PRO",
	)
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1670",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(height),
		Height: uint32(width),
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}
