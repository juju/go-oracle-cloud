package response

type Account struct {
	// Credentials specific to the account, which may include a username, password or certificate. The credentials are not returned.
	Credentials map[string]string `json:"credentials,omitempty"`
	// Description of this account.
	Description string `json:"description,omitempty"`
	// Name of the account
	Name string `json:"name,omitempty"`
	// Uniform resource identifier of the account
	Uri string `json:"uri,omitempty"`
}
