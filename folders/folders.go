package folders

import (
	"fmt"
	"github.com/gofrs/uuid"
)


// GetAllFolders retrieves all folders for a given organisation
// Paramaters:
// - req: A pointer to a FetchFolderRequest containing the organisation id
// Returns:
// - A pointer to FetchFolderResponse containing the fetched folders
// - An error, if any errors occured
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// input validation check
	if req == nil || req.OrgID == uuid.Nil {
		return nil, fmt.Errorf("invalid request: missing organization ID")
	}

	// Fetch all folders for the given organisation
	folders, err := FetchAllFoldersByOrgID(req.OrgID)

	// Error Handling
	if err != nil {
		return nil, fmt.Errorf("error fetching folders: %w", err)
	}

	return &FetchFolderResponse{Folders: folders}, nil
}

// FetchAllFoldersByOrgID retrieves all folders for a specific organization ID
// Parameters:
//   - orgID: The UUID of the organization to fetch folders for
// Returns:
//   - A slice of pointers to Folder objects belonging to the specified organization
//   - An error, if any errors occurred 
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {

	if orgID == uuid.Nil {
        return nil, fmt.Errorf("invalid organization ID")
    }

	var recoveryError error

	defer func() {
        if r := recover(); r != nil {
            recoveryError = fmt.Errorf("panic occurred: %v", r)
        }
    }()

	if recoveryError != nil {
		return nil, recoveryError
	}

	// Pointers to Folder objects by fetching data from the JSON file, sample.JSON
	folders := GetSampleData()
	
	resFolder := make([]*Folder, 0, len(folders))
	for _, folder := range folders {
		// Filter folders by the given organisiation
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
