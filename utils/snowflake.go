package utils

import (
	"github.com/bwmarrin/snowflake"
)

var snowflakeNode *snowflake.Node

func InitSnowflakeNode(number int64) {
	var err error
	snowflakeNode, err = snowflake.NewNode(number)
	if err != nil {
		panic(err)
	}
}
func GenOrderNo() string {
	return "O" + snowflakeNode.Generate().String()
}
func GenVipFeeNo() string {
	return "V" + snowflakeNode.Generate().String()
}
func GenWithdrawWords() string {
	return Str.Base62Encode(snowflakeNode.Generate().Int64())
}
