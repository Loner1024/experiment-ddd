package aggregate

import "testing"

func TestNewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "Unix",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewCustomer(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error:%v, but got %v", tc.expectedErr, err)
			}
		})
	}
}
