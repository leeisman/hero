package enums

const (
	//總參與人數：完成遊戲的總人數，跳出不計算
	RedisFinishedGameCount = "RedisFinishedGameCount"

	//跳出人數：未完成遊戲的跳出人數
	RedisUnfinishedGameCount = "RedisUnfinishedGameCount"

	//重複玩人數：同一玩家重複完整玩完的人數
	RedisRepeatUserCount   = "RedisRepeatUserCount"

	//不重複玩人數：同一玩家重複完整玩完的人數
	RedisNotRepeatUserCount   = "RedisNotRepeatUserCount"
)
