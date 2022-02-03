package dto

type (
	TeamDto struct {
		UUID string `form:"id" json:"id" param:"id" gqlgen:"id"`
		Name string `form:"name" json:"name" param:"name" graphql:"name"`
	}
)
