package models

type Note struct {
	Id      int      `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (n *Note) IsTagContained(tag string) bool {
	var tagFound = false
	for _, tagNote := range n.Tags {
		if tag == tagNote {
			tagFound = true
			break
		}
	}

	return tagFound
}
