package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.PaginatedFetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		Token: "660",
	}

	res, err := folders.GetPaginatedFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
}
