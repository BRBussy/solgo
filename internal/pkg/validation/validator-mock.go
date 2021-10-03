package validation

import (
	"sync"
	"testing"
)

type MockValidator struct {
	T                              *testing.T
	m                              sync.Mutex
	ValidateRequestFuncInvocations int64
	ValidateRequestFunc            func(t *testing.T, m *MockValidator, i interface{}) error
}

func (m *MockValidator) ValidateRequest(i interface{}) error {
	m.m.Lock()
	m.ValidateRequestFuncInvocations++
	m.m.Unlock()
	if m.ValidateRequestFunc == nil {
		return nil
	}
	return m.ValidateRequestFunc(m.T, m, i)
}
