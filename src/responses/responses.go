package responses

type GradingScale struct {
	ID        string	`json:"id"`
	CreatedAt string	`json:"createdAt"`
	UpdatedAt string	`json:"updatedAt"`
	Name      string	`json:"name"`
}

type GradingScaleMeta struct{
	ID        string					`json:"id"`
	CreatedAt string					`json:"createdAt"`
	UpdatedAt string					`json:"updatedAt"`
	Name string							`json:"name"`
	MinimumPercentage int				`json:"minPercentage"`
	MaximumPercentage int				`json:"maxPercentage"`
	Grade string						`json:"grade"`
	GpaValue int						`json:"gpaValue"`
}

type School struct {
	ID        string			`json:"id"`
	CreatedAt string			`json:"createdAt"`
	UpdatedAt string			`json:"updatedAt"`
	Name string					`json:"name"`
	GradingScale GradingScale	`json:"gradingScale"`
}

type Course struct {
	ID        string	`json:"id"`
	CreatedAt string	`json:"createdAt"`
	UpdatedAt string	`json:"updatedAt"`
	Name string			`json:"name"`
	Credit int			`json:"credit"`
}

type Result struct {
	ID        string	`json:"id"`
	CreatedAt string	`json:"createdAt"`
	UpdatedAt string	`json:"updatedAt"`
	Course     Course
	Percentage int		`json:"percentage"`
	Grade string		`json:"grade"`
}

type Student struct {
	ID        string	`json:"id"`
	CreatedAt string	`json:"createdAt"`
	UpdatedAt string	`json:"updatedAt"`
	FullName string		`json:"fullName"`
	Email    string		`json:"email"`
	Gpa float64			`json:"gpa"`
	School School		`json:"school"`
	Results []Result	`json:"results"`
}

