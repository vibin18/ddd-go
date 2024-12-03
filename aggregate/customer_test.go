package aggregate

import (
	"errors"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T){
	type testCase struct{
		test string
		name string
		expecterErr error
	}

	testCases := []testCase{
		{	
			test: "Empty name validation",
			name: "",
			expecterErr: ErrInvalidPerson,
		},{
			test: "Valid name",
			name: "Vibin",
			expecterErr: nil,
		},
	}

	for _, tc := range testCases{
		t.Run(tc.test, func(t *testing.T){
			_, err := NewCustomer(tc.name)
				if !errors.Is(err, tc.expecterErr){
					t.Errorf("extected error %v. but got %v", tc.expecterErr, err)
				}			

		})
	}

}