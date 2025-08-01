package model

import "strings"

type Verse []string

type Verses []Verse

func (v Verses) String() string {
	sb := strings.Builder{}

	for i, verse := range v {
		sb.WriteString(strings.Join(verse, "\n"))

		if i != len(v)-1 {
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}

type RawVerses string

func (v RawVerses) ToVerses() Verses {
	lines := strings.Split(string(v), "\n")
	var verses Verses

	for i := 0; i < len(lines); i += 4 {
		end := min(i+4, len(lines))

		verse := lines[i:end]
		if len(verse) > 0 {
			verses = append(verses, verse)
		}
	}

	return verses
}
