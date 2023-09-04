package generator

import (
	"github.com/brainlabs/snowflake"
	"github.com/rs/zerolog/log"
)

var (
	sfGenerator *snowflake.Node
)

// New initiated snowflake
func New(node uint64) {
	s, err := snowflake.NewNode(int64(node))

	if err != nil {
		log.Err(err).Msg("snowflake generator error")
	}

	sfGenerator = s
}

// GenerateInt64 generate id int64
func GenerateInt64() int64 {
	return sfGenerator.Generate().Int64()
}
