# topdesk-go

https://developers.topdesk.com

Partial REST and WEBDAV access to Topdesk API.

## Known Issues

### REST

The Topdesk REST API surface area is large, poorly designed and definitely not API first. Most resources are missing
one or more verbs and it tends to 500 a lot. Requests were filed on https://tip.topdesk.com years ago if you feel like
adding a +1 but our guess is the developers are buried under an avalanche of technical debt from the 90s, send help!

This package implements endpoints and fields as required to integrate Topdesk so will probably never will be complete.
Our hope is given enough time Topdesk will release a more complete API with valid schema that can be used for code
generation replacing the need for this library.

A swagger definition exists for some modules but fails to generate. E.g. supporting files as of v1.29.0 is missing
currency DTOs https://developers.topdesk.com/explorer/?page=supporting-files

## Usage

The API has been implemented with `database/sql.Rows` iterator style list access making pagination transparent.

```go
import (
  "log"

  "github.com/techspaceco/topdesk-go"
)

func main() {
  client, err := topdesk.NewRestClient(context.Background(), "https://{company}.topdesk.net/tas/api", "{token}")
  _ = err // Error handling omitted.

  branches, _ := client.ListBranches(context.Background())
  for branches.Next() { // Pagination is transparent.
    branch, err := branches.Branch()
    _ = err // Error handling omitted.

    log.Print("branch: %+v", branch)
  }

  // Last branch in list.
  //
  // Topdesk requires an {"id":"..."} structure rather than a straight ID string.
  // .Ref() creates the correct structure from complete, list or partial objects.
  branchRef := branches.Branch().Ref()

  // Resource endpoints don't share common domain models so `*Request` structs are used to allow for differences.
  person, err := client.CreatePerson(
    context.Background(),
    &topdesk.CreatePersonRequest{
      FirstName:     "Apple",
      FirstInitials: "A",
      Surname:       "Arthurton",
      Email:         "apple@arthurton.local",
      Branch:        branchRef,
      ShowBranch:    true,
    },
  )
  _ = err // Error handling omitted.

  log.Printf("person id: %s", person.ID)
}
```

## WebDAV

There is basic PUT support for file uploads.

For example our Topdesk consultant created import scripts to get around missing REST API.
