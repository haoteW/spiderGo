package sqlclient

import (
	"good-spider/model"
	"testing"
)

func TestMultiInsert(t *testing.T) {
	movieInfo := model.MovieInfo{"2338160", "https://movie.douban.com/subject/2338160/", "", "王树忱", "", "动画/短片", "中国大陆", "无对白", "1988", "19分钟", "Feelings of Mountains and Waters", "无翼鸟"}

	if err := Insert(movieInfo); err != nil {
		t.Errorf("Insert() error = %v", err)
	}
	//if err := MultiInsert(movieInfo); err != nil {
	//	t.Errorf("MultiInsert() error = %v", err)
	//}
}
