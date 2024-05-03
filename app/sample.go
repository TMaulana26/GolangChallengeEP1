package app

import "fmt"

func Multiply(x float64, y int) float64 {
	return x * float64(y)
}

func IsInteger(value interface{}) (bool, error) {
	_, ok := value.(int)
	if !ok {
		return false, fmt.Errorf("bukan integer : %v", value)
	}

	return true, fmt.Errorf("ini integer : %v", value)
}

func IntToDayName(a int) string {
	hari := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	if a > 0 && a <= len(hari) {
		return hari[a-1]
	}

	return fmt.Sprintln("error tidak ada hari ke ", a)
}

func DayNameToInt(b string) int {
	hari := map[string]int{
		"senin":  1,
		"selasa": 2,
		"rabu":   3,
		"kamis":  4,
		"jumat":  5,
		"sabtu":  6,
		"minggu": 7,
	}

	hariKe, ok := hari[b]
	if !ok {
		fmt.Println("error tidak ada nama hari lain", b, " ,ini bukan nama hari")
	}

	return hariKe
}

func IsiHello(c *string) {
	if *c == "" {
		*c = "hello world"
	}
}
