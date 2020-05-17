// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package launcher

// Struct for creating simple version JSONs.
type Version struct {
	ID string                   `json:"id"`
	Type string                 `json:"type"`
	InheritsFrom string         `json:"inheritsFrom"`
	MainClass string            `json:"mainClass,omitempty"`
	Libraries []*VersionLibrary `json:"libraries"`
}

type VersionLibrary struct {
	Name string `json:"name"`
	URL string `json:"url,omitempty"`
}