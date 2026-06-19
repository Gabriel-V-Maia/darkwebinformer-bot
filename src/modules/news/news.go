package news

type Entries struct {
	Title       string
	Description string
	Company     string
	Topic       string
}

type Feed struct {
	URL     string
	Entries []Entries
}
