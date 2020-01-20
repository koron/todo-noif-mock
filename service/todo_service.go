package service

import "github.com/koron/todo-noif-mock/do"

type Service struct {
	repos *Repository // Repositoryは自空間にエイリアスで定義してるのでpackageの指定は不要
}

func (srv *Service) PostTodo(task string) (*do.Todo, error) {
	// なにかしら必要な手続きをしたあとに…Repository のコードを呼び出す
	d, err := srv.repos.Insert(&do.Todo{Task: task})
	if err != nil {
		return nil, err
	}
	return d, nil
}
