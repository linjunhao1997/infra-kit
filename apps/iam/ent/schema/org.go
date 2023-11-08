package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

var (
	ErrUserCodeLen = errors.New("user code len err")
)

type Org struct {
	ent.Schema
}

func (Org) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("code"),
		field.String("name"),
		field.Time("ctime"),
		field.Time("mtime"),
		field.Bool("deleted"),
	}
}

func (Org) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("users", User.Type),
		edge.To("namespaces", Namespace.Type),
	}
}
