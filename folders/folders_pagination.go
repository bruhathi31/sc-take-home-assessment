package folders

import (
	"fmt"
	"github.com/gofrs/uuid"
)

// My Approach:
// - I used token-based navigation, where the token represents the starting point for the next page
// - I have a constant PageSize to determine the number of items per page, ensuring consitent response sizes. This can also be adjusted depending on the needs
// - The pagination state is entierign contianed in the request token, eliminating any need for a server-side management
// - Data retrieval is done efficiently, it relies on the FetchAllFoldersByOrgID to find all the folders, but it immediately slices the result based on the token
// - I used simple token encoding, where a string representation of an integer index is used, but it can easily be enhanced with encryption or hashing for additional security in a production setting

// This can be adjusted based on desire
const PageSize = 10

// Represents the request structure for the paginated folder fetching
type PaginatedFetchFolderRequest struct {
	OrgID uuid.UUID
	Token string
}

// Represents the response structure upon successful paginated folder fetching
type PaginatedFetchFolderResponse struct {
	Folders []*Folder
	Token   string
}

// GetPaginatedFolders retrieves a paginated folders for a specific organization ID
// Parameters:
//   - OrgID: The UUID of the organization to fetch folders for
// 	 - Token: The token for the next page, empty if it's the first page
// Returns:
//   - A slice of folders for the current page, with a maximum of folders being within the PageSize
//   - An error, if any errors occurred 
func GetPaginatedFolders(req *PaginatedFetchFolderRequest) (*PaginatedFetchFolderResponse, error) {
	// Input validation
	if req == nil || req.OrgID == uuid.Nil {
		return nil, fmt.Errorf("invalid request: missing organization ID")
	}

	// Fetch all folders for the given organisation
	allFolders, err := FetchAllFoldersByOrgID(req.OrgID)

	if err != nil {
		return nil, fmt.Errorf("error fetching folders: %w", err)
	}

	// Determine the starting index for the pagination
	startIndex := 0
	if req.Token != "" {
		startIndex = decodeToken(req.Token)
	}

	endIndex := startIndex + PageSize
	if endIndex > len(allFolders) {
		endIndex = len(allFolders)
	}

	// Slice the folders for the current page
	paginatedFolders := allFolders[startIndex:endIndex]

	// Generate token fro the next page
	var nextToken string
	if endIndex < len(allFolders) {
		nextToken = encodeToken(endIndex)
	}

	return &PaginatedFetchFolderResponse{
		Folders: paginatedFolders,
		Token:   nextToken,
	}, nil
}

// encodeToken converts an index to a string token
func encodeToken(index int) string {
	return fmt.Sprintf("%d", index)
}

// decodeToken converts an string to a index token
func decodeToken(token string) int {
	var index int
	fmt.Sscanf(token, "%d", &index)
	return index
}