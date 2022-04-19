package tag

import (
	"github.com/dotdevgo/nest/pkg/crud"
)

type TagCrud struct {
	*crud.Service[*Tag]
}

// GetTags godoc
func (s *TagCrud) GetTags(tags []string) []Tag {
	var rows []Tag

	for _, tagName := range tags {
		tag := Tag{Name: tagName}
		s.Stmt().Model(&tag).Where("name = ?", tagName).First(&tag)
		rows = append(rows, tag)
	}

	return rows
}
