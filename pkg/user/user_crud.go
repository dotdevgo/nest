package user

import (
	"dotdev/nest/pkg/crud"
	"dotdev/nest/pkg/paginator"
)

type UserCrud struct {
	*crud.Crud[*User]
}

type UserList []*User
type UserPaginator *paginator.Result[*UserList]

// Paginate godoc
func (s UserCrud) Paginate(result interface{}, pagination []paginator.Option, options ...crud.Option) (UserPaginator, error) {
	var stmt = s.Stmt(options...)

	return paginator.Paginate[*UserList](stmt, result, pagination...)
}

// FindByIdentity godoc
func (c UserCrud) FindByIdentity(identity string) (User, error) {
	var u User

	result := c.Stmt(ScopeByIdentity(identity)).First(&u)
	if err := result.Error; err != nil {
		return u, err
	}

	if u.Pk == 0 {
		return u, ErrorInvalidIdentity
	}

	return u, nil
}

// // UnmarshalGQL implements the graphql.Unmarshaler interface
// func (list *UserList) UnmarshalGQL(v interface{}) error {
// 	data, ok := v.(string)
// 	if !ok {
// 		return fmt.Errorf("UserList must be a string")
// 	}

// 	json.Unmarshal([]byte(data), list)

// 	return nil
// }

// // MarshalGQL implements the graphql.Marshaler interface
// func (list UserList) MarshalGQL(w io.Writer) {
// 	data, err := json.Marshal(list)
// 	if err == nil {
// 		w.Write(data)
// 	}
// }

// // UnmarshalGQLContext implements the graphql.ContextUnmarshaler interface
// func (l *UserList) UnmarshalGQLContext(ctx context.Context, v interface{}) error {
// 	data, ok := v.(string)
// 	if !ok {
// 		return fmt.Errorf("UserList must be a string")
// 	}

// 	json.Unmarshal([]byte(data), l)

// 	// length, err := ParseLength(s)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// *l = length
// 	return nil
// }

// // MarshalGQLContext implements the graphql.ContextMarshaler interface
// func (l UserList) MarshalGQLContext(ctx context.Context, w io.Writer) error {
// 	data, err := json.Marshal(l)
// 	if err == nil {
// 		w.Write(data)
// 	}
// 	// s, err := l.FormatContext(ctx)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// w.Write([]byte(strconv.Quote(s)))
// 	return nil
// }
