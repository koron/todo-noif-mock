// +build !mock

//go:generate mockgo -package ../repository Repository

package service

import "github.com/koron/todo-noif-mock/repository"

type Repository = repository.Repository
