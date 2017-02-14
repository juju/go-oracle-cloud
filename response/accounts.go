package response

type Account struct {
	Credentials      map[string]string `json:"credentials,omitempty"`
	Description      string            `json:"description,omitempty"`
	Accounttype      string            `json:"accounttype,omitempty"`
	Name             string            `json:"name,omitempty"`
	Uri              string            `json:"uri,omitempty"`
	Objectproperties map[string]string `json:"objectproperties,omitempty"`
}

// DirectoryNames are names of all the accounts
// in the specified container.
type DirectoryNames struct {
	Result []string `json:"result,omitempty"`
}

type AllAccount struct {
	Result []Account `json:"result,omitempty"`
}
