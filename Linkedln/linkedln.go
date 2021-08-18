package linkedln

type Linkedln struct {
	URL     string
	Name    string
	Friends int
}

func (l *Linkedln) Feed() []string {
	return []string{
		"Linkedln feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (l *Linkedln) Fame() int {
	return l.Friends
}