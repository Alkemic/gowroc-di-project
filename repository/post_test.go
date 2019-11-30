package repository

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestRepositoryFetchEntries(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &repository{
				db: tt.fields.db,
			}
			got, err := s.FetchEntries()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.FetchEntries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.FetchEntries() = %v, want %v", got, tt.want)
			}
		})
	}
}
