package constants

const HashCost = 10

type UserLabelsType struct {
	Name     string
	Password string
}

var UserLabels = UserLabelsType{
	Name:     "名前",
	Password: "パスワード",
}
