package snowflakelib

import (
	snowflake2 "github.com/bwmarrin/snowflake"
	"log"
)

var snowflake *snowflake2.Node

func init() {
	var err error
	snowflake, err = snowflake2.NewNode(1)
	if err != nil {
		log.Panicln(err)
	}
}

func Uint64ID() uint64 {
	return uint64(snowflake.Generate().Int64())
}
