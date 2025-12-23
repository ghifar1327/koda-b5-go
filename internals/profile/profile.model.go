package profile

type profile struct {
	name         string
	img          string
	email        string
	age          int
	phone_number int64
	isMarrid     bool
	education    []education
}
type education struct {
	name     string
	isPassed bool
}