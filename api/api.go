package api

type Load struct {
	Key string
}

type Store struct {
	Key   string
	Value string
}

type ValueResult struct {
	value string
}

type IntResult struct {
	result int
}
