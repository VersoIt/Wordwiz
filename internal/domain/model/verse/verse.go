package verse

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

	filteredLines := make([]string, 0, len(lines))
	for _, line := range lines {
		if len(line) != 0 {
			filteredLines = append(filteredLines, line)
		}
	}

	var verses Verses

	for i := 0; i < len(filteredLines); i += 4 {
		end := min(i+4, len(filteredLines))

		verse := filteredLines[i:end]
		if len(verse) > 0 {
			verses = append(verses, verse)
		}
	}

	return verses
}
