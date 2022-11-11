package idbuilder

import (
	"strconv"

	"github.com/sony/sonyflake"
)

type snowflake struct {
	snowflake *sonyflake.Sonyflake
}

func (s snowflake) Generate() string {
	sfid, err := s.snowflake.NextID()
	if err != nil {
		return ""
	}
	return strconv.FormatUint(sfid, 10)
}
