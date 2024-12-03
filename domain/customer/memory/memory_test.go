package memory

import (
	"ddd/aggregate"
	"ddd/domain/customer"
	"errors"
	"testing"

	"github.com/google/uuid"
)


func TestMemory_getCustomer(t *testing.T){
	type testCase struct {
		name string
		id uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("Vibin")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := Repository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name: "no customer by id",
			id: uuid.MustParse("3e94e7f1-e8ce-4339-a544-8a6f99dad7a4"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name: "customer by id",
			id: id,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases{
		t.Run(tc.name, func(t *testing.T){
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr){
				t.Errorf("expected error is %v, got %v", tc.expectedErr, err)
			}
		})
	}
}