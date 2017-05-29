package kong

import "testing"

type FakeApiInfoer struct {
	Name string
	Err  error
}

// ()   ...
func (f FakeApiInfoer) GetApiInfo(name string) (string, error) {

	if f.Err != nil {
		return "", f.Err
	}
	return f.Name, nil
}

func TestGetInfoApi(t *testing.T) {
	f := FakeApiInfoer{
		Name: "MyApi",
		Err:  nil,
	}

	expectedName := "MyApi"

	name, err := GetApiInfo(f, "MyApi")

	if err != nil {
		t.Fatalf("Expected error to be nil but it was %s", err)
	}

	if name != expectedName {
		t.Fatalf("Expected %s but got %s", expectedName, name)
	}

}
