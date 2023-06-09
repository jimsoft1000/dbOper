package model

func (Version) TableName() string {
	return "sys_version"
}

type Version struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
