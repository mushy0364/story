package parts

import s "github.com/skilstak/storyeng-go"

func init() {

	s.Parts["Start"] = s.Part{
		"onenter": "Welcome to this story game.",
		"oninput": "Name",
	}

	s.Parts["Welcome"] = s.Part{
		"onenter": "My name is Normal. What is your name?",
		"oninput": func(e s.InputEvent) {
			s.Data["name"] = e.Line
		},
		"onleave": "Nice to meet you {{.name}}.",
	}
}
