package lobsang

import (
	"crypto/sha1"
	"encoding/base64"
	"time"
)

type Entry struct {
	Comment  string
	Duration time.Duration
}

func (e Entry) Hash() string {
	hasher := sha1.New()
	hasher.Write([]byte(e.Duration.String()))
	hasher.Write([]byte(e.Comment))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha[:8]
}
