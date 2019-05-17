# topdesk-go

https://developers.topdesk.com

Partial REST and WEBDAV access to Topdesk API.

## Known Issues

### REST

The API is very incomplete, requests have been filed in https://tip.topdesk.com

A The swagger definition exists for some modules but fails to generate. Supporting files as of v1.29.0 missing missing currency DTOs
https://developers.topdesk.com/explorer/?page=supporting-files

#### Resources

The REST API is missing a lot of POST, PUT, DELETE & GET verbs for resources:
* Suppliers missing INDEX, GET only.
* Locations INDEX, GET only.
* Extra fields INDEX, POST only.

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
    branch := branches.Branch()
    log.Print("branch: %+v", branch)
  }

  branch := branches.Branch() // Last branch in list.
  person := &topdesk.Person{
    FirstName:     "Apple",
    FirstInitials: "A",
    Surname:       "Arthurton",
    Email:         "apple@arthurton.local",
    Branch:        branch.Ref(), // Topdesk requires an {"id":"..."} structure rather than a straight ID string.
    ShowBranch:    true,
  }

  id, err := client.SavePerson(context.Background(), person)
  log.Printf("person id: %s", id)
}
```

## WebDAV

There is basic PUT support for file uploads.

For example our Topdesk consultant created import scripts to get around missing REST API.
