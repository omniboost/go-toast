package toast_test

import (
	"encoding/json"
	"testing"
)

func EmployeesPostRequest(t *testing.T) {
	b := []byte(`{"first_name": "John", "last_name": "Doe", "email": "TJ7gW@example.com", "phone": "1234567890", "role": "waiter", "password": "password"}`)
	req := client.NewEmployeesPostRequest()
	json.Unmarshal(b, req.RequestBody())
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ = json.MarshalIndent(resp, "", "  ")
	t.Log(string(b))
}
