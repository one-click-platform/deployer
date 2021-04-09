package helpers

import (
	"strconv"
)

// TODO change Payload to struct
func ParsePayload(payload string) (int64, error) {
	return strconv.ParseInt(payload, 10, 64)
}
