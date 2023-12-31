package entity

type CreateConnectionParam struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

type CreateConnectionResult struct {
	ConnectionId int
}

type DeleteConnectionResult struct {
	ConnectionId int
}

type GetTableRecordsParam struct {
	TableName  string
	Limit      int
	Offset     int
	SortBy     string
	OrderBy    string
	Conditions []Condition
}

type GetTableRecordsResult struct {
	Data   []map[string]any
	Total  int
	Limit  int
	Offset int
}
