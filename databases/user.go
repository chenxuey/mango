package databases

type Users struct {
	Id   int64
	Name string
	Sex  int
	Age  int
}

type UsersFac struct {
	Table Users
}

func (uf *UsersFac) Get() error {
	uf.Table.Id = 1
	return nil
}