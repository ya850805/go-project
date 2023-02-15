package model

type student struct {
	Name  string
	score float64
}

//因為student結構體首字母小寫，因此只能在model包使用
//通過工廠模式來解決

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

// GetScore 如果score字段首字母小寫，則在其他包不可以直接使用
// 此時我們可以提供一個方法
func (s *student) GetScore() float64 {
	return s.score
}
