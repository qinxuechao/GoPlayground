package model

type MysqlParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string // Database to use
}

type HTTPServerParams struct {
	Host string
	Port string
}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Role       string `json:"role"`
}
