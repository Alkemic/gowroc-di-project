package blog

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Alkemic/gowroc-di-project/repository"
)

type mockPostRepository struct {
	fetchEntriesResult []repository.Post
	fetchEntriesErr    error
	getEntryResult     repository.Post
	getEntryErr        error
	getEntryParams     []int
}

func (m *mockPostRepository) FetchEntries() ([]repository.Post, error) {
	return m.fetchEntriesResult, m.fetchEntriesErr
}

func (m *mockPostRepository) GetEntry(id int) (repository.Post, error) {
	m.getEntryParams = append(m.getEntryParams, id)
	return m.getEntryResult, m.getEntryErr
}

func TestBlogServiceList(t *testing.T) {
	type check func([]repository.Post, error, *mockPostRepository, *testing.T)
	checks := func(cs ...check) []check { return cs }
	hasNoError := func(_ []repository.Post, err error, _ *mockPostRepository, t *testing.T) {
		t.Helper()
		if err != nil {
			t.Errorf("Expected err to be nil, but got '%v'", err)
		}
	}
	hasPosts := func(expectedResult []repository.Post) check {
		return func(result []repository.Post, _ error, _ *mockPostRepository, t *testing.T) {
			t.Helper()
			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected result to be '%v', but got '%v'", expectedResult, result)
			}
		}
	}
	hasError := func(expectedErr error) check {
		return func(_ []repository.Post, err error, _ *mockPostRepository, t *testing.T) {
			t.Helper()
			if !errors.Is(err, expectedErr) {
				t.Errorf("Expected error to be '%v', but got '%v'", expectedErr, err)
			}
		}
	}

	mockedErr := errors.New("mocked error")
	tests := []struct {
		name               string
		fetchEntriesResult []repository.Post
		fetchEntriesErr    error
		checks             []check
	}{{
		name: "success",
		fetchEntriesResult: []repository.Post{{
			Id:        11,
			Title:     "title 11",
			CreatedAt: time.Date(2016, 3, 17, 7, 56, 35, 0, time.UTC),
			Content:   "content 11",
		}, {
			Id:        12,
			Title:     "title 12",
			CreatedAt: time.Date(2016, 3, 19, 7, 56, 35, 0, time.UTC),
			Content:   "content 12",
		}},
		checks: checks(
			hasNoError,
			hasPosts([]repository.Post{{
				Id:        11,
				Title:     "title 11",
				CreatedAt: time.Date(2016, 3, 17, 7, 56, 35, 0, time.UTC),
				Content:   "content 11",
			}, {
				Id:        12,
				Title:     "title 12",
				CreatedAt: time.Date(2016, 3, 19, 7, 56, 35, 0, time.UTC),
				Content:   "content 12",
			}}),
		),
	}, {
		name:            "return nil on error",
		fetchEntriesErr: mockedErr,
		checks: checks(
			hasError(mockedErr),
			hasPosts(nil),
		),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedPostRepository := &mockPostRepository{
				fetchEntriesResult: tt.fetchEntriesResult,
				fetchEntriesErr:    tt.fetchEntriesErr,
			}
			s := blogService{
				postRepository: mockedPostRepository,
			}
			result, err := s.List()
			for _, ch := range tt.checks {
				ch(result, err, mockedPostRepository, t)
			}
		})
	}
}

func TestBlogServiceGet(t *testing.T) {
	type check func(repository.Post, error, *mockPostRepository, *testing.T)
	checks := func(cs ...check) []check { return cs }
	hasNoError := func(_ repository.Post, err error, _ *mockPostRepository, t *testing.T) {
		t.Helper()
		if err != nil {
			t.Errorf("Expected err to be nil, but got '%v'", err)
		}
	}
	hasPosts := func(expectedResult repository.Post) check {
		return func(result repository.Post, _ error, _ *mockPostRepository, t *testing.T) {
			t.Helper()
			if !reflect.DeepEqual(result, expectedResult) {
				t.Errorf("Expected result to be '%v', but got '%v'", expectedResult, result)
			}
		}
	}
	hasError := func(expectedErr error) check {
		return func(_ repository.Post, err error, _ *mockPostRepository, t *testing.T) {
			t.Helper()
			if !errors.Is(err, expectedErr) {
				t.Errorf("Expected error to be '%v', but got '%v'", expectedErr, err)
			}
		}
	}

	mockedErr := errors.New("mocked error")
	tests := []struct {
		name           string
		getEntryResult repository.Post
		getEntryErr    error
		checks         []check
	}{{
		name: "success",
		getEntryResult: repository.Post{
			Id:        12,
			Title:     "title 12",
			CreatedAt: time.Date(2016, 3, 19, 7, 56, 35, 0, time.UTC),
			Content:   "content 12",
		},
		checks: checks(
			hasNoError,
			hasPosts(repository.Post{
				Id:        12,
				Title:     "title 12",
				CreatedAt: time.Date(2016, 3, 19, 7, 56, 35, 0, time.UTC),
				Content:   "content 12",
			}),
		),
	}, {
		name: "return empty struct on error",
		getEntryResult: repository.Post{
			Id:        12,
			Title:     "title 12",
			CreatedAt: time.Date(2016, 3, 19, 7, 56, 35, 0, time.UTC),
			Content:   "content 12",
		},
		getEntryErr: mockedErr,
		checks: checks(
			hasError(mockedErr),
			hasPosts(repository.Post{}),
		),
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedPostRepository := &mockPostRepository{
				getEntryResult: tt.getEntryResult,
				getEntryErr:    tt.getEntryErr,
			}
			s := blogService{
				postRepository: mockedPostRepository,
			}
			result, err := s.Get(16)
			for _, ch := range tt.checks {
				ch(result, err, mockedPostRepository, t)
			}
		})
	}
}
