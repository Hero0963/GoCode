package p33

import (
	"math/rand"
	"time"
)

func getLuckyBall() string {

	List := getCommentList()
	l := len(List)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(l)

	return List[idx]

}

func getCommentList() []string {

	return []string{"Yes.", "No.", "Maybe.", "Ask ask later."}
}
