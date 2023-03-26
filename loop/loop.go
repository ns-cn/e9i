package loop

const (
	NONE = iota
	SINGLE
	LOOP_QUEUE
	LOOP_RANDOM
)

var (
	Title = []string{"不循环", "单曲循环", "队列", "随机"}
)
