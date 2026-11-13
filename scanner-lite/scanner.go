package scanner

type Table struct {
	Name     string
	RowСount int
}

type Result struct {
	TableName string
	HasPII    bool
}

func Scan(tables []Table, workes int) []Result {

}
