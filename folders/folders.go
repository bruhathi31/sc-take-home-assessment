package folders

import (
	"github.com/gofrs/uuid"
)


// GetAllFolders retrieves all folders for a given organisation
// Paramaters:
// - req: A pointer to a FetchFolderRequest containing the organisation id
// Returns:
// - A pointer to FetchFolderResponse containing the fetched folders
// - An error, if any errors occured
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Variable declation block(unused variables)
	var (
		err error
		f1  Folder
		fs  []*Folder
	)

	// Initialise an empty slice of Folders
	f := []Folder{}

	// Fetch all folders for the given organisation
	r, _ := FetchAllFoldersByOrgID(req.OrgID)

	// Appending all the fetched folders (converting []*Folder to []Folder)
	for k, v := range r {
		f = append(f, *v)
	}

	// converting []Folder to []*Folder
	var fp []*Folder
	for k1, v1 := range f {
		fp = append(fp, &v1)
	}

	// Intialising and returning the FetchFolderResponse
	var ffr *FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	return ffr, nil
}

// FetchAllFoldersByOrgID retrieves all folders for a specific organization ID
// Parameters:
//   - orgID: The UUID of the organization to fetch folders for
// Returns:
//   - A slice of pointers to Folder objects belonging to the specified organization
//   - An error, if any errors occurred 
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	// Pointers to Folder objects by fetching data from the JSON file, sample.JSON
	folders := GetSampleData()

	
	resFolder := []*Folder{}
	for _, folder := range folders {
		// Filter folders by the given organisiation
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
