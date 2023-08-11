package models

import (
	"reflect"
	"testing"

	"private_test/db"
)

func Test_CreateOrder(t *testing.T) {
	// The test work for the first time, after that since the order_id is unique and primary key the column will be duplicate and cannot be inserted.
	// You could test it again with this test by change the order_id uuid in this code :
	// arrobj = append(arrobj, DBCheck{"order_id : CHANGE_HERE_WITH_36_CHAR; product_id : 4a0cb611-bcf7-49cb-b7cd-565d6383d6f1;"})

	var arrobj []DBCheck
	arrobj = append(arrobj, DBCheck{"customer_id : eb83059b-8187-4c3a-ad36-923608911787; status : fresh_order; "})
	arrobj = append(arrobj, DBCheck{"order_id : 4a0cb611-bcf7-49cb-b7cd-565d6383test; product_id : 4a0cb611-bcf7-49cb-b7cd-565d6383d6f1;"})

	type args struct {
		data map[string]interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected Response
	}{
		{
			name: "insert order 1 product",
			args: args{
				data: map[string]interface{}{
					"customer_id": "eb83059b-8187-4c3a-ad36-923608911787",
					"product_id":  "4a0cb611-bcf7-49cb-b7cd-565d6383d6f1",
					"order_id":    "4a0cb611-bcf7-49cb-b7cd-565d6383test",
				},
			},
			expected: Response{Status: 200, Messages: "Create Order Success!", Data: arrobj},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db.Init()
			result, err := CreateOrders(tt.args.data["customer_id"].(string), tt.args.data["order_id"].(string), tt.args.data["product_id"].(string))
			if err != nil {
				t.Error("error when CreateOrder, " + err.Error())
			}
			if !reflect.DeepEqual(tt.expected, result) {
				t.Error("test expected not equal result")
			}
		})
	}
}
