// Package topdesk is a partial implementation of the Topdesk REST/WebDAV API.
//
// The Topdesk REST API surface area is large, poorly designed and definitely not API first. Most resources are missing
// one or more verbs and it tends to 500 a lot. Requests were filed on https://tip.topdesk.com years ago if you feel like
// adding a +1 but our guess is the developers are buried under an avalanche of technical debt from the 90s, send help!
//
// This package implements endpoints and fields as required to integrate Topdesk so will probably never will be complete.
// Our hope is given enough time Topdesk will release a more complete API with valid schema that can be used for code
// generation replacing the need for this library.
//
// A swagger definition exists for some modules but fails to generate. E.g. supporting files as of v1.29.0 is missing
// currency DTOs https://developers.topdesk.com/explorer/?page=supporting-files
//
// The API has been implemented with `database/sql.Rows` iterator style list access making pagination transparent.
package topdesk
