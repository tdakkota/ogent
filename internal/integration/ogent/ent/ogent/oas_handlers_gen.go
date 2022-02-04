// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

// HandleCreateCategoryRequest handles createCategory operation.
//
// POST /categories
func (s *Server) handleCreateCategoryRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateCategory",
		trace.WithAttributes(otelogen.OperationID("createCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateCategoryRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateCategory(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreateCategoryPetsRequest handles createCategoryPets operation.
//
// POST /categories/{id}/pets
func (s *Server) handleCreateCategoryPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateCategoryPets",
		trace.WithAttributes(otelogen.OperationID("createCategoryPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreateCategoryPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreateCategoryPetsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateCategoryPets(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateCategoryPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetRequest handles createPet operation.
//
// POST /pets
func (s *Server) handleCreatePetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePet",
		trace.WithAttributes(otelogen.OperationID("createPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePet(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetCategoriesRequest handles createPetCategories operation.
//
// POST /pets/{id}/categories
func (s *Server) handleCreatePetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetCategories",
		trace.WithAttributes(otelogen.OperationID("createPetCategories")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetCategoriesParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetCategoriesRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetCategories(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetFriendsRequest handles createPetFriends operation.
//
// POST /pets/{id}/friends
func (s *Server) handleCreatePetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetFriends",
		trace.WithAttributes(otelogen.OperationID("createPetFriends")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetFriendsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetFriendsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetFriends(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreatePetOwnerRequest handles createPetOwner operation.
//
// POST /pets/{id}/owner
func (s *Server) handleCreatePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreatePetOwner",
		trace.WithAttributes(otelogen.OperationID("createPetOwner")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreatePetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreatePetOwnerRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreatePetOwner(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreatePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreateUserRequest handles createUser operation.
//
// POST /users
func (s *Server) handleCreateUserRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateUser",
		trace.WithAttributes(otelogen.OperationID("createUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	request, err := decodeCreateUserRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateUser(ctx, request)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreateUserBestFriendRequest handles createUserBestFriend operation.
//
// POST /users/{id}/best-friend
func (s *Server) handleCreateUserBestFriendRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateUserBestFriend",
		trace.WithAttributes(otelogen.OperationID("createUserBestFriend")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreateUserBestFriendParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreateUserBestFriendRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateUserBestFriend(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateUserBestFriendResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleCreateUserPetsRequest handles createUserPets operation.
//
// POST /users/{id}/pets
func (s *Server) handleCreateUserPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateUserPets",
		trace.WithAttributes(otelogen.OperationID("createUserPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeCreateUserPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeCreateUserPetsRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.CreateUserPets(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeCreateUserPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeleteCategoryRequest handles deleteCategory operation.
//
// DELETE /categories/{id}
func (s *Server) handleDeleteCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteCategory",
		trace.WithAttributes(otelogen.OperationID("deleteCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeleteCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeletePetRequest handles deletePet operation.
//
// DELETE /pets/{id}
func (s *Server) handleDeletePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeletePet",
		trace.WithAttributes(otelogen.OperationID("deletePet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeletePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeletePet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeletePetOwnerRequest handles deletePetOwner operation.
//
// DELETE /pets/{id}/owner
func (s *Server) handleDeletePetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeletePetOwner",
		trace.WithAttributes(otelogen.OperationID("deletePetOwner")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeletePetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeletePetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeletePetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeleteUserRequest handles deleteUser operation.
//
// DELETE /users/{id}
func (s *Server) handleDeleteUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteUser",
		trace.WithAttributes(otelogen.OperationID("deleteUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeleteUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleDeleteUserBestFriendRequest handles deleteUserBestFriend operation.
//
// DELETE /users/{id}/best-friend
func (s *Server) handleDeleteUserBestFriendRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteUserBestFriend",
		trace.WithAttributes(otelogen.OperationID("deleteUserBestFriend")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeDeleteUserBestFriendParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.DeleteUserBestFriend(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeDeleteUserBestFriendResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListCategoryRequest handles listCategory operation.
//
// GET /categories
func (s *Server) handleListCategoryRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListCategory",
		trace.WithAttributes(otelogen.OperationID("listCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListCategoryPetsRequest handles listCategoryPets operation.
//
// GET /categories/{id}/pets
func (s *Server) handleListCategoryPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListCategoryPets",
		trace.WithAttributes(otelogen.OperationID("listCategoryPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListCategoryPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListCategoryPets(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListCategoryPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetRequest handles listPet operation.
//
// GET /pets
func (s *Server) handleListPetRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPet",
		trace.WithAttributes(otelogen.OperationID("listPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetCategoriesRequest handles listPetCategories operation.
//
// GET /pets/{id}/categories
func (s *Server) handleListPetCategoriesRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetCategories",
		trace.WithAttributes(otelogen.OperationID("listPetCategories")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetCategoriesParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetCategories(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetCategoriesResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListPetFriendsRequest handles listPetFriends operation.
//
// GET /pets/{id}/friends
func (s *Server) handleListPetFriendsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListPetFriends",
		trace.WithAttributes(otelogen.OperationID("listPetFriends")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListPetFriendsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListPetFriends(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListPetFriendsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListUserRequest handles listUser operation.
//
// GET /users
func (s *Server) handleListUserRequest(args [0]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListUser",
		trace.WithAttributes(otelogen.OperationID("listUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleListUserPetsRequest handles listUserPets operation.
//
// GET /users/{id}/pets
func (s *Server) handleListUserPetsRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ListUserPets",
		trace.WithAttributes(otelogen.OperationID("listUserPets")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeListUserPetsParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ListUserPets(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeListUserPetsResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadCategoryRequest handles readCategory operation.
//
// GET /categories/{id}
func (s *Server) handleReadCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadCategory",
		trace.WithAttributes(otelogen.OperationID("readCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadCategory(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadPetRequest handles readPet operation.
//
// GET /pets/{id}
func (s *Server) handleReadPetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPet",
		trace.WithAttributes(otelogen.OperationID("readPet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPet(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadPetOwnerRequest handles readPetOwner operation.
//
// GET /pets/{id}/owner
func (s *Server) handleReadPetOwnerRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadPetOwner",
		trace.WithAttributes(otelogen.OperationID("readPetOwner")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadPetOwnerParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadPetOwner(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadPetOwnerResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadUserRequest handles readUser operation.
//
// GET /users/{id}
func (s *Server) handleReadUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadUser",
		trace.WithAttributes(otelogen.OperationID("readUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadUser(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleReadUserBestFriendRequest handles readUserBestFriend operation.
//
// GET /users/{id}/best-friend
func (s *Server) handleReadUserBestFriendRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "ReadUserBestFriend",
		trace.WithAttributes(otelogen.OperationID("readUserBestFriend")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeReadUserBestFriendParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.ReadUserBestFriend(ctx, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeReadUserBestFriendResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleUpdateCategoryRequest handles updateCategory operation.
//
// PATCH /categories/{id}
func (s *Server) handleUpdateCategoryRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateCategory",
		trace.WithAttributes(otelogen.OperationID("updateCategory")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdateCategoryParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdateCategoryRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdateCategory(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdateCategoryResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleUpdatePetRequest handles updatePet operation.
//
// PATCH /pets/{id}
func (s *Server) handleUpdatePetRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdatePet",
		trace.WithAttributes(otelogen.OperationID("updatePet")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdatePetParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdatePetRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdatePet(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdatePetResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

// HandleUpdateUserRequest handles updateUser operation.
//
// PATCH /users/{id}
func (s *Server) handleUpdateUserRequest(args [1]string, w http.ResponseWriter, r *http.Request) {
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateUser",
		trace.WithAttributes(otelogen.OperationID("updateUser")),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	params, err := decodeUpdateUserParams(args, r)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}
	request, err := decodeUpdateUserRequest(r, span)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := s.h.UpdateUser(ctx, request, params)
	if err != nil {
		span.RecordError(err)
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	if err := encodeUpdateUserResponse(response, w, span); err != nil {
		span.RecordError(err)
		return
	}
}

func respondError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, writeErr := json.Marshal(struct {
		ErrorMessage string `json:"error_message"`
	}{
		ErrorMessage: err.Error(),
	})
	if writeErr == nil {
		w.Write(data)
	}
}
