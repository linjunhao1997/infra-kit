package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return xid.New().String()
			}),
		field.String("name"),
		field.String("email"),
		field.Bool("disabled"),
		field.Int("ver"),
		field.Time("ctime"),
		field.Time("mtime"),
		field.Bool("deleted"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("org", Org.Type).Ref("users").Unique(),
		edge.From("groups", Group.Type).Ref("users"),
		edge.From("namespace", Namespace.Type).Ref("users"),
		edge.To("credential", Credential.Type).Unique(),
	}
}
