package twitter

type Twitter struct {
	URL     string
	Name    string
	Friends int
}

func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (t *Twitter) Fame() int {
	return t.Friends
}