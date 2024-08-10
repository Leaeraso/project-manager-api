package main

// Mocks

type MockStore struct {
}

func (m *MockStore) CreateUser(u *User) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) GetUserByID(id string) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) GetUserByEmail(email string) (*User, error) {
	return &User{}, nil
}

func (m *MockStore) CreateTask(t *Task) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) GetTask(id string) (*Task, error) {
	return &Task{}, nil
}

func (m *MockStore) CreateProject(p *Project) (*Project, error) {
	return &Project{}, nil
}

func (m *MockStore) GetProjectById(id string) (*Project, error) {
	return &Project{}, nil
}
