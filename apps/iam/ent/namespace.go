// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"infra-kit/apps/iam/ent/namespace"
	"infra-kit/apps/iam/ent/org"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Namespace is the model entity for the Namespace schema.
type Namespace struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// Ctime holds the value of the "ctime" field.
	Ctime time.Time `json:"ctime,omitempty"`
	// Mtime holds the value of the "mtime" field.
	Mtime time.Time `json:"mtime,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted bool `json:"deleted,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NamespaceQuery when eager-loading is set.
	Edges          NamespaceEdges `json:"edges"`
	org_namespaces *string
	selectValues   sql.SelectValues
}

// NamespaceEdges holds the relations/edges for other nodes in the graph.
type NamespaceEdges struct {
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NamespaceEdges) OrgOrErr() (*Org, error) {
	if e.loadedTypes[0] {
		if e.Org == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: org.Label}
		}
		return e.Org, nil
	}
	return nil, &NotLoadedError{edge: "org"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Namespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case namespace.FieldDeleted:
			values[i] = new(sql.NullBool)
		case namespace.FieldID, namespace.FieldCode, namespace.FieldName, namespace.FieldDescription, namespace.FieldLabel:
			values[i] = new(sql.NullString)
		case namespace.FieldCtime, namespace.FieldMtime:
			values[i] = new(sql.NullTime)
		case namespace.ForeignKeys[0]: // org_namespaces
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Namespace fields.
func (n *Namespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case namespace.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				n.ID = value.String
			}
		case namespace.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				n.Code = value.String
			}
		case namespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case namespace.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				n.Description = value.String
			}
		case namespace.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				n.Label = value.String
			}
		case namespace.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				n.Ctime = value.Time
			}
		case namespace.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				n.Mtime = value.Time
			}
		case namespace.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				n.Deleted = value.Bool
			}
		case namespace.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field org_namespaces", values[i])
			} else if value.Valid {
				n.org_namespaces = new(string)
				*n.org_namespaces = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Namespace.
// This includes values selected through modifiers, order, etc.
func (n *Namespace) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryOrg queries the "org" edge of the Namespace entity.
func (n *Namespace) QueryOrg() *OrgQuery {
	return NewNamespaceClient(n.config).QueryOrg(n)
}

// QueryUsers queries the "users" edge of the Namespace entity.
func (n *Namespace) QueryUsers() *UserQuery {
	return NewNamespaceClient(n.config).QueryUsers(n)
}

// Update returns a builder for updating this Namespace.
// Note that you need to call Namespace.Unwrap() before calling this method if this Namespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Namespace) Update() *NamespaceUpdateOne {
	return NewNamespaceClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Namespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Namespace) Unwrap() *Namespace {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Namespace is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Namespace) String() string {
	var builder strings.Builder
	builder.WriteString("Namespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("code=")
	builder.WriteString(n.Code)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(n.Description)
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(n.Label)
	builder.WriteString(", ")
	builder.WriteString("ctime=")
	builder.WriteString(n.Ctime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(n.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted=")
	builder.WriteString(fmt.Sprintf("%v", n.Deleted))
	builder.WriteByte(')')
	return builder.String()
}

// Namespaces is a parsable slice of Namespace.
type Namespaces []*Namespace
