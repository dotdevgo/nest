package tag

import (
	"dotdev/nest/pkg/crud"
)

type TagCrud struct {
	*crud.Crud[*Tag]
}

// GetTags godoc
func (s *TagCrud) GetTags(tags []string) []Tag {
	var rows []Tag

	for _, tagName := range tags {
		tag := Tag{Name: tagName}
		s.Stmt().
			Model(&tag).
			Where("name = ?", tagName).
			First(&tag)
		rows = append(rows, tag)
		// if err := s.Stmt().
		// 	Model(&tag).
		// 	Where("name = ?", tagName).
		// 	First(&tag).Error; err == nil {
		// 	rows = append(rows, tag)
		// }
	}

	return rows
}