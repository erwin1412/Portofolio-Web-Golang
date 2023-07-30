package models

type Project struct {
	ID           int
	Title        string
	Content      string
	Author       string
	PostDate     string
	StartDate    string
	EndDate      string
	ImagePath    string
	Technologies []string
	User_id      int
}
