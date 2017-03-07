// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package response

// Error internal response error type
type Error struct {
	Message string `json:"message"`
}
