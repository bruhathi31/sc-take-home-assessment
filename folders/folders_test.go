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

	// t.Run("Error - getSampleData has panic and function recovers", func(t *testing.T) {
    //     // Override the getSampleData for this test
    //     folders.getSampleData = func() []*folders.Folder {
    //         panic("simulated panic")
    //     }

    //     req := &folders.FetchFolderRequest{
    //         OrgID: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")),
    //     }

    //     response, err := folders.GetAllFolders(req)
    //     assert.Error(t, err)
    //     assert.Nil(t, response)
    //     assert.Contains(t, err.Error(), "panic occurred")
    // })
}