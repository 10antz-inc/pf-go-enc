package enc

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// RandomNumber 関数は 18 桁まで.
func RandomNumber(digits uint) string {
	if 18 < digits {
		panic("Do not exceed 18 digits")
	}
	t := time.Now()
	max := math.Pow(10, float64(digits))
	n := rand.New(rand.NewSource(t.UnixNano())).Int63n(int64(max))
	return fmt.Sprintf("%0"+strconv.Itoa(int(digits))+"v", n)
}
