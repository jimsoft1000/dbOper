package model

func (Version) TableName() string {
	return "sys_version"
}

type Version struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
}
