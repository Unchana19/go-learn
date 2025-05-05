package main

import (
	"net/http"
	"strconv"

	"github.com/Unchana19/go-learn/internal/store"
	"github.com/go-chi/chi/v5"
)

type userContextKey string

const userCtxKey userContextKey = "user"

// ActivateUser godoc
//
//	@Summary	Activate user account
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		token	path		string	true	"User activation token"
//	@Success	204		{object}	nil
//	@Failure	404		{object}	error
//	@Failure	500		{object}	error
//	@Router		/v1/users/activate/{token} [put]
func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	ctx := r.Context()

	err := app.store.Users.Activate(ctx, token)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFound(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// GetUser godoc
//
//	@Summary	Get user by ID
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		userID	path		int	true	"User ID"
//	@Success	200		{object}	store.User
//	@Failure	404		{object}	error
//	@Router		/v1/users/{userID} [get]
func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

// FollowUser godoc
//
//	@Summary	Follow a user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		userID	path		int	true	"User ID"
//	@Success	200		{object}	store.User
//	@Failure	404		{object}	error
//	@Router		/v1/users/{userID}/follow [put]
func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)

	followedID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Follow(ctx, followerUser.ID, followedID); err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFound(w, r, err)
			return
		case store.ErrDuplicate:
			app.conflictResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

// UnfollowUser godoc
//
//	@Summary	Unfollow a user
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		userID	path		int	true	"User ID"
//	@Success	200		{object}	store.User
//	@Failure	404		{object}	error
//	@Router		/v1/users/{userID}/unfollow [put]
func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollowerUser := getUserFromContext(r)

	unfollowedID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Followers.Unfollow(ctx, unfollowerUser.ID, unfollowedID); err != nil {
		switch err {
		case store.ErrNotFound:
			app.notFound(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
	}
}

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtxKey).(*store.User)
	return user
}
