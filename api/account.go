package api

func (c Client) AllAccounts() (respponse.AllAccounts, err error) {
	if !c.isAuth() {
		return ErrNotAuth
	}
	
	url := fmt.Sprintf("%s/%s/Computer-%s/",c.endpoint, "account", c.identify)

}
