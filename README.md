<h1 align="center">
    <img src=".github/alligator.png" width="240" />
</h1>

# 🐊 Alligator Go Client for Pterodactyl
> This fork attempts to follow [Official Pterodactyl API Docs @Dashflo](https://dashflo.net/docs/api/pterodactyl/v1)
> as closely as possible, but is not maintained by Pterodactyl team.

### Installation
```bash
go get -u github.com/Coneriys/alligator
```
---


## 🔥 So, what's new?

### ✨ Support for Filters, Include and Endpoint-specific parameters
*Including support for extended struct fields
- **Example for [List Users](https://dashflo.net/docs/api/pterodactyl/v1/#req_5703791f721f4b50bb0318cf19e2262d) endpoint**
```go
package main
 import (
	 gator "github.com/Coneriys/alligator"
	 "github.com/Coneriys/alligator/options"
 )

 func main() {
	 app, _ := gator.NewApp("https://panel.pterodactyl.io", "ApplicationKeyYouCreated")
	 
	 // Fetch some users
	 users, err := app.ListUsers(options.ListUsersOptions{
		 Include: options.IncludeUsers{
			 Servers: true,
		 },
		 Filters: options.FiltersUsers{
			 Username: "example",
		 },
		 SortBy: options.ListUsersSort_UUID_DESC, // Same as "-uuid"
	 })
	 if err != nil {
		 fmt.Printf("%#v", err)
		 return
	 }
	 
	 // options are optional btw
	 // users, err := app.ListUsers() // <- is also valid

	 for _, u := range users {
		 fmt.Printf("%d: %s\n", u.ID, u.Username)
	 }
 }
```

**More examples at [📁 _examples](_examples)**

## 📝 What's done?
- [X] App API
  - [X] Options
  - [X] Database endpoint support [TESTING]
    - [X] Extended databases details (password, host) [TESTING]
  - [X] Nests endpoint support
    - [X] Extended nest details (eggs, servers)
  - [X] Eggs endpoint support
    - [X] Extended eggs details (nest, servers, variables)
  - [X] Extended user details (servers)
  - [X] Extended nodes details (allocations, location, servers)
  - [X] Extended allocations details (node, server)
  - [X] Extended location details (nodes, servers)
  - [X] Extended servers details (allocations+, user+, subusers+, nest+, egg+, variables+, location+, node+)
  - [X] Extended servers details (databases) [TESTING]
  - [X] Additional methods like `/{server}/reinstall` and `/{server}/force`
- [ ] Client API
  - [ ] What is this goofy ahh infinite documentation...
- [ ] Pagination (50 servers limit is a pain tbh)
- [X] Godoc (App)
- [ ] Godoc (Client)