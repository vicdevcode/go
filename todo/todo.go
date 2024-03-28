package todo

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TodoUpdate struct {
	Id int `json:"id"`
}

type TodoAdd struct {
	Title string `json:"title"`
}

type Data struct {
	count int
	Todos []Todo
}

func NewData() *Data {
	return &Data{Todos: []Todo{}, count: 0}
}

func (data *Data) AddTodo(title string) {
	data.count += 1
	data.Todos = append(data.Todos, Todo{
		Id:    data.count,
		Title: title,
		Done:  false,
	})
}

func (data *Data) UpdateTodo(id int) {
	for i, todo := range data.Todos {
		if todo.Id == id {
			data.Todos[i].Done = !data.Todos[i].Done
		}
	}
}
