package bitbucket

// import (
// 	"net/http"
// 	"testing"
// )
//
// func Test_BasicAuth(t *testing.T) {
// 	a := &BasicAuth{
// 		Username: "foo",
// 		Password: "bar",
// 	}
//
// 	r, err := http.NewRequest("GET", "/test", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	a.authenticate(r)
//
// 	username, password, _ := r.BasicAuth()
//
// 	if username != a.Username {
// 		t.Errorf("Username not set, expected: %s, got: %s", a.Username, username)
// 	}
//
// 	if password != a.Password {
// 		t.Errorf("Password set incorrectly, expected: %s, got: %s", a.Password, password)
// 	}
// }
