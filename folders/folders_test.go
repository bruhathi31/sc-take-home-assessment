package folders_test

import (
	"testing"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/stretchr/testify/assert"
	"github.com/gofrs/uuid"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Success Test - shorter length", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("52214b35-f4da-461a-9f93-fbd3590e700f")),
		}

		expected := &folders.FetchFolderResponse{
			Folders: []*folders.Folder{
				{
					Id:      uuid.Must(uuid.FromString("9fc98418-7039-4443-a82b-84049ed25d38")),
					Name:    "fitting-talisman",
					OrgId:   uuid.Must(uuid.FromString("52214b35-f4da-461a-9f93-fbd3590e700f")),
					Deleted: false,
				},
			},
		}

		response, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.Equal(t, expected, response)
	})

	t.Run("Success Test - high number of occurances", func(t *testing.T) {
		req := &folders.FetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
		}
		response, err := folders.GetAllFolders(req)
		assert.NoError(t, err)
		assert.Equal(t, 666, len(response.Folders))
	})

	t.Run("Invalid Input - empty request", func(t *testing.T) {
		response, err := folders.GetAllFolders(nil)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("Invalid Input - empty OrgId", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: uuid.Nil}
		response, err := folders.GetAllFolders(req)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}

func Test_GetPaginatedFolders(t *testing.T) {
	t.Run("Success Test - First Page", func(t *testing.T) {
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			Token: "",
		}

		response, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, folders.PageSize, len(response.Folders))
		assert.NotEmpty(t, response.Token)
	})

	t.Run("Success Test - Second Page", func(t *testing.T) {
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			Token: "10",
		}

		response, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, folders.PageSize, len(response.Folders))
		assert.NotEmpty(t, response.Token)
	})

	t.Run("Success Test - Last Page", func(t *testing.T) {
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
			Token: "660",
		}

		response, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 6, len(response.Folders))
		assert.Empty(t, response.Token)
	})

	t.Run("Success Test - Org with Few Folders", func(t *testing.T) {
		req := &folders.PaginatedFetchFolderRequest{
			OrgID: uuid.Must(uuid.FromString("52214b35-f4da-461a-9f93-fbd3590e700f")),
			Token: "",
		}

		response, err := folders.GetPaginatedFolders(req)
		assert.NoError(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, 1, len(response.Folders))
		assert.Empty(t, response.Token)
	})

	t.Run("Invalid Input - empty request", func(t *testing.T) {
		response, err := folders.GetPaginatedFolders(nil)
		assert.Error(t, err)
		assert.Nil(t, response)
	})

	t.Run("Invalid Input - empty OrgId", func(t *testing.T) {
		req := &folders.PaginatedFetchFolderRequest{OrgID: uuid.Nil}
		response, err := folders.GetPaginatedFolders(req)
		assert.Error(t, err)
		assert.Nil(t, response)
	})
}