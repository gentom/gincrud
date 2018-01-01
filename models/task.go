package task

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:@/gincrud")
	if err != nil {
		panic(err)
	}
}

type Task struct {
	ID   int    `json:"id" xorm:"'id'"`
	Text string `json:"text" xorm:"'text'"`
}

type Tasks []Task

type Repository struct {
}

func TaskRepository() Repository {
	return Repository{}
}

func (m Repository) Create(text string) {
	var task = Task{Text: text}
	engine.Insert(&task)
}

func (m Repository) GetAll() Tasks {
	var task = Task{}
	var tasks = Tasks{}
	rows, err := engine.Rows(task)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&task)
		tasks = append(tasks, task)
	}
	return tasks
}
