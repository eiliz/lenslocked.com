package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

// NewUsers is used to ceeate a new Users controller. This function will panic if the templates are not parsed correctly and should only be used during the initial setup
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// Users is the users controller
type Users struct {
	NewView *views.View
}

// New is used to render the form to signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

// Create is usee to process the signup form and create a new user account
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	// r.PostForm = map[string][]string
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
	// fmt.Fprintln(w, r.PostFormValue("email"))

}
