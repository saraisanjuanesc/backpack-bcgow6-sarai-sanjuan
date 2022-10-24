package users

import "errors"

type MockStorage struct {
	dataMock    []User
	readCalled  bool
	writeCalled bool
}

func (m *MockStorage) Read(data interface{}) error {
	m.readCalled = true
	dat, ok := data.(*[]User)
	if !ok {
		return errors.New("Error in method Read")
	}
	*dat = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	m.writeCalled = true
	dat, ok := data.([]User)
	if !ok {
		return errors.New("Error in method Write")
	}
	m.dataMock = dat
	return nil
}
