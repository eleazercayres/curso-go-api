package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClient_SetFirstName(t *testing.T) {
	var _uuid = uuid.NewV4()

	type fields struct {
		ID          int
		UUID        uuid.UUID
		FirstName   string
		LastName    string
		Age         int
		PhoneNumber PhoneNumber
		Address     Address
	}
	type args struct {
		n string
	}
	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *Client
		wantErr error
	}{
		{
			name: "happy path",
			fields: fields{
				ID:          1,
				UUID:        _uuid,
				FirstName:   "A",
				LastName:    "B",
				Age:         20,
				PhoneNumber: PhoneNumber("987654321"),
				Address:     Address{},
			},
			args: args{n: "X"},
			want: &Client{
				ID:          1,
				UUID:        _uuid,
				FirstName:   "X",
				LastName:    "B",
				Age:         20,
				PhoneNumber: PhoneNumber("987654321"),
				Address:     Address{},
			},
			wantErr: nil,
		},
		{
			name: "error: empty value",
			fields: fields{
				ID:          1,
				UUID:        _uuid,
				FirstName:   "A",
				LastName:    "B",
				Age:         20,
				PhoneNumber: PhoneNumber("987654321"),
				Address:     Address{},
			},
			args: args{n: ""},
			want: &Client{
				ID:          1,
				UUID:        _uuid,
				FirstName:   "A",
				LastName:    "B",
				Age:         20,
				PhoneNumber: PhoneNumber("987654321"),
				Address:     Address{},
			},
			wantErr: errors.New("empty value"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				ID:          tt.fields.ID,
				UUID:        tt.fields.UUID,
				FirstName:   tt.fields.FirstName,
				LastName:    tt.fields.LastName,
				Age:         tt.fields.Age,
				PhoneNumber: tt.fields.PhoneNumber,
				Address:     tt.fields.Address,
			}

			err := c.SetFirstName(tt.args.n)

			if err != nil && tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error(), "Should be the same")
			} else {
				if !reflect.DeepEqual(c, tt.want) {
					t.Errorf("SetFirstName() = %v, want %v", c, tt.want)
				}
			}
		})
	}
}
