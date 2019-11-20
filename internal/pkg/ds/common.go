package ds

// Config read in from cmd line
type Config struct {
	DBtype    string
	DBhost    string
	DBname    string
	DBport    int
	DBuser    string
	Tablename string
	Password  string
}
