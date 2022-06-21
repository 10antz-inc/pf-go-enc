package enc

import (
	"math/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

var (
	ulid_t       = time.Now()
	ulid_ts      = ulid.Timestamp(ulid_t)
	ulid_entropy = ulid.Monotonic(rand.New(rand.NewSource(ulid_t.UnixNano())), 0)
)

func ULID() string {
	return ulid.MustNew(ulid_ts, ulid_entropy).String()
}
