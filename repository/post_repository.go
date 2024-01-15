package repository

import "github.com/zulfianfreza/test-sharing-vision-backend/entity"

type PostRepository interface {
	FindAll() ([]entity.Post, error)
	FindById(postId int) (entity.Post, error)
	Save(post entity.Post) (entity.Post, error)
	Update(post entity.Post) (entity.Post, error)
	Delete(post entity.Post) (entity.Post, error)
}
