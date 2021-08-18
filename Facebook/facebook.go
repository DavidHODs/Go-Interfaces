package facebook

type Facebook struct {
	URL     string
	Name    string
	Friends int
}

func (f *Facebook) Feed() []string {
	return []string{
		"Facebook feeds",
		"Hey, here is my cool selfie",
		"What is code?",
	}
}

func (f *Facebook) Fame() int {
	return f.Friends
}