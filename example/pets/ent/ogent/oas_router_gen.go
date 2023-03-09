// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListCategoryRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateCategoryRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteCategoryRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadCategoryRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateCategoryRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/pets"
						if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListCategoryPetsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'd': // Prefix: "db-health"
				if l := len("db-health"); len(elem) >= l && elem[0:l] == "db-health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleDBHealthRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListPetRequest([0]string{}, w, r)
					case "POST":
						s.handleCreatePetRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeletePetRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadPetRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdatePetRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "categories"
							if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListPetCategoriesRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'f': // Prefix: "friends"
							if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListPetFriendsRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 'o': // Prefix: "owner"
							if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadPetOwnerRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListUserRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateUserRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteUserRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadUserRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateUserRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/pets"
						if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListUserPetsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "categories"
				if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListCategory"
						r.operationID = "listCategory"
						r.pathPattern = "/categories"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateCategory"
						r.operationID = "createCategory"
						r.pathPattern = "/categories"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteCategory"
							r.operationID = "deleteCategory"
							r.pathPattern = "/categories/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadCategory"
							r.operationID = "readCategory"
							r.pathPattern = "/categories/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateCategory"
							r.operationID = "updateCategory"
							r.pathPattern = "/categories/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/pets"
						if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListCategoryPets
								r.name = "ListCategoryPets"
								r.operationID = "listCategoryPets"
								r.pathPattern = "/categories/{id}/pets"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'd': // Prefix: "db-health"
				if l := len("db-health"); len(elem) >= l && elem[0:l] == "db-health" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: DBHealth
						r.name = "DBHealth"
						r.operationID = "DBHealth"
						r.pathPattern = "/db-health"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			case 'p': // Prefix: "pets"
				if l := len("pets"); len(elem) >= l && elem[0:l] == "pets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListPet"
						r.operationID = "listPet"
						r.pathPattern = "/pets"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreatePet"
						r.operationID = "createPet"
						r.pathPattern = "/pets"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeletePet"
							r.operationID = "deletePet"
							r.pathPattern = "/pets/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadPet"
							r.operationID = "readPet"
							r.pathPattern = "/pets/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdatePet"
							r.operationID = "updatePet"
							r.pathPattern = "/pets/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'c': // Prefix: "categories"
							if l := len("categories"); len(elem) >= l && elem[0:l] == "categories" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ListPetCategories
									r.name = "ListPetCategories"
									r.operationID = "listPetCategories"
									r.pathPattern = "/pets/{id}/categories"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'f': // Prefix: "friends"
							if l := len("friends"); len(elem) >= l && elem[0:l] == "friends" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ListPetFriends
									r.name = "ListPetFriends"
									r.operationID = "listPetFriends"
									r.pathPattern = "/pets/{id}/friends"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 'o': // Prefix: "owner"
							if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadPetOwner
									r.name = "ReadPetOwner"
									r.operationID = "readPetOwner"
									r.pathPattern = "/pets/{id}/owner"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListUser"
						r.operationID = "listUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateUser"
						r.operationID = "createUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteUser"
							r.operationID = "deleteUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadUser"
							r.operationID = "readUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateUser"
							r.operationID = "updateUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/pets"
						if l := len("/pets"); len(elem) >= l && elem[0:l] == "/pets" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListUserPets
								r.name = "ListUserPets"
								r.operationID = "listUserPets"
								r.pathPattern = "/users/{id}/pets"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
