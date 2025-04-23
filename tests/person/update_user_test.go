package person_test

import (
	"errors"
	"testing"

	"github.com/abdullinmm/grok-lean-go/internal/models"
)

func TestUpdateUser(t *testing.T) {

	test := []struct {
		person   *models.Person
		testName string
		name     string
		age      int
		expected error
	}{
		{
			testName: "Update User Name and Age",
			person:   &models.Person{Name: "Marsel", Age: 44},
			name:     "Marat",
			age:      50,
			expected: nil,
		},
	}
	for _, tt := range test {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.person.Validate()
			if !errors.Is(err, tt.expected) {
				t.Errorf("Expected %s, got %s", tt.expected, err)
			}
		})
	}
}
