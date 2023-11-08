package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type Group struct {
	ent.Schema
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("code"),
		field.String("name"),
		field.String("description"),
		field.Time("ctime"),
		field.Time("mtime"),
		field.Bool("deleted"),
	}
}

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("groups").Unique(),
		edge.To("users", User.Type),
		edge.To("authorities", Authority.Type),
	}
}
