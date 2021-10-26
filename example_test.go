package jsoncat_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/clarify/jsoncat"
)

// Model assumed a base model for an API client.
type Model struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"` // read-only field
	CreatedBy string    `json:"createdBy"` // read-only field
}

func (m Model) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID string `json:"id"`
	}{
		ID: m.ID,
	})
}

type AttrUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type User struct {
	Model
	AttrUser
}

func (a User) MarshalJSON() ([]byte, error) {
	return jsoncat.MarshalObject(a.Model, a.AttrUser)
}

func Example_embedded_struct_unamrshal() {
	data := []byte(`{
		"id": "kari",
		"firstName": "Kari",
		"lastName": "Normann",
		"createdAt": "2020-01-01T10:20:30Z",
		"createdBy": "john"
	}`)
	var u User
	err := json.Unmarshal(data, &u)
	fmt.Printf("err is %v\nu is %+v\n", err, u)
	// Output:
	// err is <nil>
	// u is {Model:{ID:kari CreatedAt:2020-01-01 10:20:30 +0000 UTC CreatedBy:john} AttrUser:{FirstName:Kari LastName:Normann}}
}

func Example_embedded_struct_marshal() {
	u := User{
		Model: Model{
			ID:        "kari",
			CreatedAt: time.Date(2020, 01, 01, 10, 20, 30, 0, time.UTC),
			CreatedBy: "john",
		},
		AttrUser: AttrUser{
			FirstName: "Kari",
			LastName:  "Normann",
		},
	}
	data, err := json.Marshal(u)
	fmt.Printf("err is %v\ndata is %s", err, data)
	// Output:
	// err is <nil>
	// data is {"id":"kari","firstName":"Kari","lastName":"Normann"}
}
