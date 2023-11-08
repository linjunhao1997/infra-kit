package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Authority struct {
	ent.Schema
}

func (Authority) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("code"),
		field.String("name"),
		field.Time("ctime"),
		field.Time("mtime"),
	}
}

func (Authority) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).Ref("authorities"),
	}
}
