// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"time"

	"github.com/go-faster/jx"
)

// Ref: #/components/schemas/CategoryCreate
type CategoryCreate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *CategoryCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *CategoryCreate) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *CategoryCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *CategoryCreate) SetName(val string) {
	s.Name = val
}

func (*CategoryCreate) createCategoryRes() {}

// Ref: #/components/schemas/CategoryList
type CategoryList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *CategoryList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *CategoryList) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *CategoryList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *CategoryList) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/Category_PetsList
type CategoryPetsList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *CategoryPetsList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *CategoryPetsList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *CategoryPetsList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *CategoryPetsList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *CategoryPetsList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *CategoryPetsList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *CategoryPetsList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *CategoryPetsList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/CategoryRead
type CategoryRead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *CategoryRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *CategoryRead) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *CategoryRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *CategoryRead) SetName(val string) {
	s.Name = val
}

func (*CategoryRead) readCategoryRes() {}

// Ref: #/components/schemas/CategoryUpdate
type CategoryUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *CategoryUpdate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *CategoryUpdate) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *CategoryUpdate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *CategoryUpdate) SetName(val string) {
	s.Name = val
}

func (*CategoryUpdate) updateCategoryRes() {}

type CreateCategoryReq struct {
	Name string `json:"name"`
	Pets []int  `json:"pets"`
}

// GetName returns the value of Name.
func (s *CreateCategoryReq) GetName() string {
	return s.Name
}

// GetPets returns the value of Pets.
func (s *CreateCategoryReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *CreateCategoryReq) SetName(val string) {
	s.Name = val
}

// SetPets sets the value of Pets.
func (s *CreateCategoryReq) SetPets(val []int) {
	s.Pets = val
}

type CreatePetReq struct {
	Name       string      `json:"name"`
	Weight     OptInt      `json:"weight"`
	Birthday   OptDateTime `json:"birthday"`
	Categories []int       `json:"categories"`
	Owner      int         `json:"owner"`
	Friends    []int       `json:"friends"`
}

// GetName returns the value of Name.
func (s *CreatePetReq) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *CreatePetReq) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *CreatePetReq) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *CreatePetReq) GetCategories() []int {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *CreatePetReq) GetOwner() int {
	return s.Owner
}

// GetFriends returns the value of Friends.
func (s *CreatePetReq) GetFriends() []int {
	return s.Friends
}

// SetName sets the value of Name.
func (s *CreatePetReq) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *CreatePetReq) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *CreatePetReq) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *CreatePetReq) SetCategories(val []int) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *CreatePetReq) SetOwner(val int) {
	s.Owner = val
}

// SetFriends sets the value of Friends.
func (s *CreatePetReq) SetFriends(val []int) {
	s.Friends = val
}

type CreateUserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Pets []int  `json:"pets"`
}

// GetName returns the value of Name.
func (s *CreateUserReq) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *CreateUserReq) GetAge() int {
	return s.Age
}

// GetPets returns the value of Pets.
func (s *CreateUserReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *CreateUserReq) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *CreateUserReq) SetAge(val int) {
	s.Age = val
}

// SetPets sets the value of Pets.
func (s *CreateUserReq) SetPets(val []int) {
	s.Pets = val
}

// DBHealthNoContent is response for DBHealth operation.
type DBHealthNoContent struct{}

func (*DBHealthNoContent) dBHealthRes() {}

// DBHealthServiceUnavailable is response for DBHealth operation.
type DBHealthServiceUnavailable struct{}

func (*DBHealthServiceUnavailable) dBHealthRes() {}

// DeleteCategoryNoContent is response for DeleteCategory operation.
type DeleteCategoryNoContent struct{}

func (*DeleteCategoryNoContent) deleteCategoryRes() {}

// DeletePetNoContent is response for DeletePet operation.
type DeletePetNoContent struct{}

func (*DeletePetNoContent) deletePetRes() {}

// DeleteUserNoContent is response for DeleteUser operation.
type DeleteUserNoContent struct{}

func (*DeleteUserNoContent) deleteUserRes() {}

type ListCategoryOKApplicationJSON []CategoryList

func (*ListCategoryOKApplicationJSON) listCategoryRes() {}

type ListCategoryPetsOKApplicationJSON []CategoryPetsList

func (*ListCategoryPetsOKApplicationJSON) listCategoryPetsRes() {}

type ListPetCategoriesOKApplicationJSON []PetCategoriesList

func (*ListPetCategoriesOKApplicationJSON) listPetCategoriesRes() {}

type ListPetFriendsOKApplicationJSON []PetFriendsList

func (*ListPetFriendsOKApplicationJSON) listPetFriendsRes() {}

type ListPetOKApplicationJSON []PetList

func (*ListPetOKApplicationJSON) listPetRes() {}

type ListUserOKApplicationJSON []UserList

func (*ListUserOKApplicationJSON) listUserRes() {}

type ListUserPetsOKApplicationJSON []UserPetsList

func (*ListUserPetsOKApplicationJSON) listUserPetsRes() {}

// NewOptDateTime returns new OptDateTime with value set to v.
func NewOptDateTime(v time.Time) OptDateTime {
	return OptDateTime{
		Value: v,
		Set:   true,
	}
}

// OptDateTime is optional time.Time.
type OptDateTime struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDateTime was set.
func (o OptDateTime) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDateTime) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDateTime) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDateTime) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDateTime) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pet_CategoriesList
type PetCategoriesList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *PetCategoriesList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCategoriesList) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *PetCategoriesList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCategoriesList) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/PetCreate
type PetCreate struct {
	ID         int                   `json:"id"`
	Name       string                `json:"name"`
	Weight     OptInt                `json:"weight"`
	Birthday   OptDateTime           `json:"birthday"`
	Categories []PetCreateCategories `json:"categories"`
	Owner      PetCreateOwner        `json:"owner"`
}

// GetID returns the value of ID.
func (s *PetCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreate) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetCreate) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetCreate) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *PetCreate) GetCategories() []PetCreateCategories {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *PetCreate) GetOwner() PetCreateOwner {
	return s.Owner
}

// SetID sets the value of ID.
func (s *PetCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreate) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetCreate) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetCreate) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *PetCreate) SetCategories(val []PetCreateCategories) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *PetCreate) SetOwner(val PetCreateOwner) {
	s.Owner = val
}

func (*PetCreate) createPetRes() {}

// Ref: #/components/schemas/PetCreate_Categories
type PetCreateCategories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetID returns the value of ID.
func (s *PetCreateCategories) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreateCategories) GetName() string {
	return s.Name
}

// SetID sets the value of ID.
func (s *PetCreateCategories) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreateCategories) SetName(val string) {
	s.Name = val
}

// Ref: #/components/schemas/PetCreate_Owner
type PetCreateOwner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *PetCreateOwner) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetCreateOwner) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *PetCreateOwner) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *PetCreateOwner) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetCreateOwner) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *PetCreateOwner) SetAge(val int) {
	s.Age = val
}

// Ref: #/components/schemas/Pet_FriendsList
type PetFriendsList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetFriendsList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetFriendsList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetFriendsList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetFriendsList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetFriendsList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetFriendsList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetFriendsList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetFriendsList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/PetList
type PetList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/Pet_OwnerRead
type PetOwnerRead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *PetOwnerRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetOwnerRead) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *PetOwnerRead) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *PetOwnerRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetOwnerRead) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *PetOwnerRead) SetAge(val int) {
	s.Age = val
}

func (*PetOwnerRead) readPetOwnerRes() {}

// Ref: #/components/schemas/PetRead
type PetRead struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetRead) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetRead) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetRead) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetRead) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetRead) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetRead) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

func (*PetRead) readPetRes() {}

// Ref: #/components/schemas/PetUpdate
type PetUpdate struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *PetUpdate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *PetUpdate) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *PetUpdate) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *PetUpdate) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *PetUpdate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *PetUpdate) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *PetUpdate) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *PetUpdate) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

func (*PetUpdate) updatePetRes() {}

type R400 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R400) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R400) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R400) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R400) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R400) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R400) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R400) createCategoryRes()    {}
func (*R400) createPetRes()         {}
func (*R400) createUserRes()        {}
func (*R400) deleteCategoryRes()    {}
func (*R400) deletePetRes()         {}
func (*R400) deleteUserRes()        {}
func (*R400) listCategoryPetsRes()  {}
func (*R400) listCategoryRes()      {}
func (*R400) listPetCategoriesRes() {}
func (*R400) listPetFriendsRes()    {}
func (*R400) listPetRes()           {}
func (*R400) listUserPetsRes()      {}
func (*R400) listUserRes()          {}
func (*R400) readCategoryRes()      {}
func (*R400) readPetOwnerRes()      {}
func (*R400) readPetRes()           {}
func (*R400) readUserRes()          {}
func (*R400) updateCategoryRes()    {}
func (*R400) updatePetRes()         {}
func (*R400) updateUserRes()        {}

type R404 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R404) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R404) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R404) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R404) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R404) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R404) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R404) deleteCategoryRes()    {}
func (*R404) deletePetRes()         {}
func (*R404) deleteUserRes()        {}
func (*R404) listCategoryPetsRes()  {}
func (*R404) listCategoryRes()      {}
func (*R404) listPetCategoriesRes() {}
func (*R404) listPetFriendsRes()    {}
func (*R404) listPetRes()           {}
func (*R404) listUserPetsRes()      {}
func (*R404) listUserRes()          {}
func (*R404) readCategoryRes()      {}
func (*R404) readPetOwnerRes()      {}
func (*R404) readPetRes()           {}
func (*R404) readUserRes()          {}
func (*R404) updateCategoryRes()    {}
func (*R404) updatePetRes()         {}
func (*R404) updateUserRes()        {}

type R409 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R409) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R409) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R409) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R409) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R409) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R409) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R409) createCategoryRes()    {}
func (*R409) createPetRes()         {}
func (*R409) createUserRes()        {}
func (*R409) deleteCategoryRes()    {}
func (*R409) deletePetRes()         {}
func (*R409) deleteUserRes()        {}
func (*R409) listCategoryPetsRes()  {}
func (*R409) listCategoryRes()      {}
func (*R409) listPetCategoriesRes() {}
func (*R409) listPetFriendsRes()    {}
func (*R409) listPetRes()           {}
func (*R409) listUserPetsRes()      {}
func (*R409) listUserRes()          {}
func (*R409) readCategoryRes()      {}
func (*R409) readPetOwnerRes()      {}
func (*R409) readPetRes()           {}
func (*R409) readUserRes()          {}
func (*R409) updateCategoryRes()    {}
func (*R409) updatePetRes()         {}
func (*R409) updateUserRes()        {}

type R500 struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors jx.Raw `json:"errors"`
}

// GetCode returns the value of Code.
func (s *R500) GetCode() int {
	return s.Code
}

// GetStatus returns the value of Status.
func (s *R500) GetStatus() string {
	return s.Status
}

// GetErrors returns the value of Errors.
func (s *R500) GetErrors() jx.Raw {
	return s.Errors
}

// SetCode sets the value of Code.
func (s *R500) SetCode(val int) {
	s.Code = val
}

// SetStatus sets the value of Status.
func (s *R500) SetStatus(val string) {
	s.Status = val
}

// SetErrors sets the value of Errors.
func (s *R500) SetErrors(val jx.Raw) {
	s.Errors = val
}

func (*R500) createCategoryRes()    {}
func (*R500) createPetRes()         {}
func (*R500) createUserRes()        {}
func (*R500) deleteCategoryRes()    {}
func (*R500) deletePetRes()         {}
func (*R500) deleteUserRes()        {}
func (*R500) listCategoryPetsRes()  {}
func (*R500) listCategoryRes()      {}
func (*R500) listPetCategoriesRes() {}
func (*R500) listPetFriendsRes()    {}
func (*R500) listPetRes()           {}
func (*R500) listUserPetsRes()      {}
func (*R500) listUserRes()          {}
func (*R500) readCategoryRes()      {}
func (*R500) readPetOwnerRes()      {}
func (*R500) readPetRes()           {}
func (*R500) readUserRes()          {}
func (*R500) updateCategoryRes()    {}
func (*R500) updatePetRes()         {}
func (*R500) updateUserRes()        {}

type UpdateCategoryReq struct {
	Name OptString `json:"name"`
	Pets []int     `json:"pets"`
}

// GetName returns the value of Name.
func (s *UpdateCategoryReq) GetName() OptString {
	return s.Name
}

// GetPets returns the value of Pets.
func (s *UpdateCategoryReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *UpdateCategoryReq) SetName(val OptString) {
	s.Name = val
}

// SetPets sets the value of Pets.
func (s *UpdateCategoryReq) SetPets(val []int) {
	s.Pets = val
}

type UpdatePetReq struct {
	Name       OptString   `json:"name"`
	Weight     OptInt      `json:"weight"`
	Birthday   OptDateTime `json:"birthday"`
	Categories []int       `json:"categories"`
	Owner      OptInt      `json:"owner"`
	Friends    []int       `json:"friends"`
}

// GetName returns the value of Name.
func (s *UpdatePetReq) GetName() OptString {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *UpdatePetReq) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *UpdatePetReq) GetBirthday() OptDateTime {
	return s.Birthday
}

// GetCategories returns the value of Categories.
func (s *UpdatePetReq) GetCategories() []int {
	return s.Categories
}

// GetOwner returns the value of Owner.
func (s *UpdatePetReq) GetOwner() OptInt {
	return s.Owner
}

// GetFriends returns the value of Friends.
func (s *UpdatePetReq) GetFriends() []int {
	return s.Friends
}

// SetName sets the value of Name.
func (s *UpdatePetReq) SetName(val OptString) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *UpdatePetReq) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *UpdatePetReq) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// SetCategories sets the value of Categories.
func (s *UpdatePetReq) SetCategories(val []int) {
	s.Categories = val
}

// SetOwner sets the value of Owner.
func (s *UpdatePetReq) SetOwner(val OptInt) {
	s.Owner = val
}

// SetFriends sets the value of Friends.
func (s *UpdatePetReq) SetFriends(val []int) {
	s.Friends = val
}

type UpdateUserReq struct {
	Name OptString `json:"name"`
	Age  OptInt    `json:"age"`
	Pets []int     `json:"pets"`
}

// GetName returns the value of Name.
func (s *UpdateUserReq) GetName() OptString {
	return s.Name
}

// GetAge returns the value of Age.
func (s *UpdateUserReq) GetAge() OptInt {
	return s.Age
}

// GetPets returns the value of Pets.
func (s *UpdateUserReq) GetPets() []int {
	return s.Pets
}

// SetName sets the value of Name.
func (s *UpdateUserReq) SetName(val OptString) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *UpdateUserReq) SetAge(val OptInt) {
	s.Age = val
}

// SetPets sets the value of Pets.
func (s *UpdateUserReq) SetPets(val []int) {
	s.Pets = val
}

// Ref: #/components/schemas/UserCreate
type UserCreate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *UserCreate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *UserCreate) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *UserCreate) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *UserCreate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *UserCreate) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *UserCreate) SetAge(val int) {
	s.Age = val
}

func (*UserCreate) createUserRes() {}

// Ref: #/components/schemas/UserList
type UserList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *UserList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *UserList) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *UserList) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *UserList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *UserList) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *UserList) SetAge(val int) {
	s.Age = val
}

// Ref: #/components/schemas/User_PetsList
type UserPetsList struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Weight   OptInt      `json:"weight"`
	Birthday OptDateTime `json:"birthday"`
}

// GetID returns the value of ID.
func (s *UserPetsList) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *UserPetsList) GetName() string {
	return s.Name
}

// GetWeight returns the value of Weight.
func (s *UserPetsList) GetWeight() OptInt {
	return s.Weight
}

// GetBirthday returns the value of Birthday.
func (s *UserPetsList) GetBirthday() OptDateTime {
	return s.Birthday
}

// SetID sets the value of ID.
func (s *UserPetsList) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *UserPetsList) SetName(val string) {
	s.Name = val
}

// SetWeight sets the value of Weight.
func (s *UserPetsList) SetWeight(val OptInt) {
	s.Weight = val
}

// SetBirthday sets the value of Birthday.
func (s *UserPetsList) SetBirthday(val OptDateTime) {
	s.Birthday = val
}

// Ref: #/components/schemas/UserRead
type UserRead struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *UserRead) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *UserRead) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *UserRead) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *UserRead) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *UserRead) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *UserRead) SetAge(val int) {
	s.Age = val
}

func (*UserRead) readUserRes() {}

// Ref: #/components/schemas/UserUpdate
type UserUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetID returns the value of ID.
func (s *UserUpdate) GetID() int {
	return s.ID
}

// GetName returns the value of Name.
func (s *UserUpdate) GetName() string {
	return s.Name
}

// GetAge returns the value of Age.
func (s *UserUpdate) GetAge() int {
	return s.Age
}

// SetID sets the value of ID.
func (s *UserUpdate) SetID(val int) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *UserUpdate) SetName(val string) {
	s.Name = val
}

// SetAge sets the value of Age.
func (s *UserUpdate) SetAge(val int) {
	s.Age = val
}

func (*UserUpdate) updateUserRes() {}
