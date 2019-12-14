package blog

import (
	"fmt"

	"github.com/Alkemic/gowroc-di-project/repository"
)

type BlogService interface {
	List() ([]repository.Post, error)
	Get(id int) (repository.Post, error)
}

type blogService struct {
	postRepository repository.PostRepository
}

func NewBlogService(postRepository repository.PostRepository) BlogService {
	return &blogService{
		postRepository: postRepository,
	}
}

func (s blogService) List() ([]repository.Post, error) {
	posts, err := s.postRepository.FetchEntries()
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %w", err)
	}
	return posts, nil
}

func (s blogService) Get(id int) (repository.Post, error) {
	post, err := s.postRepository.GetEntry(id)
	if err != nil {
		return repository.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return post, nil
}
