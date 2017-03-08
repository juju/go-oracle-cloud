// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api_test

import (
	enc "encoding/json"
	"fmt"
	"net/http"

	gc "gopkg.in/check.v1"
)

var cookie = `nimbula=eyJpZGVudGl0eSI6ICJ7XCJyZWFsbVwiOiBcInVzY29tLWNlbnRyYWwtMVwiLCBcInZhbHVlXCI6IFwie1xcXCJjdXN0b21lclxcXCI6IFxcXCJDb21wdXRlLWE0MzIxMDBcXFwiLCBcXFwicmVhbG1cXFwiOiBcXFwidXNjb20tY2VudHJhbC0xXFxcIiwgXFxcImVudGl0eV90eXBlXFxcIjogXFxcInVzZXJcXFwiLCBcXFwic2Vzc2lvbl9leHBpcmVzXFxcIjogMTQ4ODkwMzI1Mi43MjY2MywgXFxcImV4cGlyZXNcXFwiOiAxNDg4ODk0MjUyLjcyNjY3NzksIFxcXCJ1c2VyXFxcIjogXFxcIi9Db21wdXRlLWE0MzIxMDAvc2dpdWxpdHRpQGNsb3VkYmFzZS5jb21cXFwiLCBcXFwiZ3JvdXBzXFxcIjogW1xcXCIvQ29tcHV0ZS1hNDMyMTAwL0NvbXB1dGUuQ29tcHV0ZV9PcGVyYXRpb25zXFxcIiwgXFxcIi9Db21wdXRlLWE0MzIxMDAvQ29tcHV0ZS5Db21wdXRlX01vbml0b3JcXFwiLCBcXFwiL0NvbXB1dGUtYTQzMjEwMC9Db21wdXRlLkNvbXB1dGVfT3BlcmF0aW9uc1xcXCIsIFxcXCIvQ29tcHV0ZS1hNDMyMTAwL0NvbXB1dGUuQ29tcHV0ZV9Nb25pdG9yXFxcIl19XCIsIFwic2lnbmF0dXJlXCI6IFwiWGRhTFo2WHV5K2t0SkVVZmppbzh6VmlHNVlIbGdORkZIS2JpaGdzVjd6Nkx3NXZiMTlaT1lXck8zZ3VIV25USWZOaWY0SG9sQzhsa3gvWjE0eHVNcVJiMklraVQrelIzalVzbUYyb1I2QWVrclVyaTcxOXN5RFNxZ1V6Y3d0TXYxbXl2aVJON2xMMzhsSHFnbVZ1Wkw3QUxSaTBKRDhGQ3ZWa3U5WFhDT014QmM0QnVMaDViUVg1dk1NTW1HYXZMNjNwODgzM2MxdVRyaWNlNHBzT1ZlOWg2dDBDeGJTUzExYkkrUS9IamtHQjExWUg0cFpUZm5QMW9KcjdUelV0YWY2R2RhTzl5M1FiczFsMXFwTEtSWDBjSEFzclgvdWFjZ2tOc21rREpUa25nTEhsR0dLMnZEMlBpdW1Eb2lNczBlU1hGdndyUndsNUtuYWRuNDNBLzNBPT1cIn0ifQ==; Path=/; Max-Age=1800`

func (cl clientTest) TestAuthentication(c *gc.C) {

	ts, client := cl.StartTestServer(httpParams{
		check:              c,
		manualHeaderStatus: true,
		handler: func(w http.ResponseWriter, r *http.Request) {
			c.Assert(r.Method, gc.Equals, http.MethodPost)
			c.Assert(r.Header.Get("Content-Type"), gc.DeepEquals, json)
			c.Assert(r.Header.Get("Accept"), gc.DeepEquals, json)

			auth := struct {
				User     string `json:"user"`
				Password string `json:"password"`
			}{}

			err := enc.NewDecoder(r.Body).Decode(&auth)
			c.Assert(err, gc.IsNil)
			c.Assert(auth.User, gc.DeepEquals, fmt.Sprintf("/Compute-%s/%s",
				identify, username))
			c.Assert(auth.Password, gc.DeepEquals, password)

			// give the client a new cookie
			w.Header().Set("Set-Cookie", cookie)
			w.Header().Set("Content-Type", json)
			w.WriteHeader(http.StatusNoContent)
		},
	})
	defer ts.Close()

	// authenticate to the oracle api
	err := client.Authenticate()
	c.Assert(err, gc.IsNil)
}
