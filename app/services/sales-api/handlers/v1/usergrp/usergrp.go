// Package usergrp maintains the group of handlers for user access.
package usergrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ardanlabs/service/business/core/user"
	v1 "github.com/ardanlabs/service/business/web/v1"
	"github.com/ardanlabs/service/business/web/v1/paging"
	"github.com/ardanlabs/service/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	user *user.Core
}

// New constructs a handlers for route access.
func New(user *user.Core) *Handlers {
	return &Handlers{
		user: user,
	}
}

// Create adds a new user to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewUser
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	nc, err := toCoreNewUser(app)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	usr, err := h.user.Create(ctx, nc)
	if err != nil {
		if errors.Is(err, user.ErrUniqueEmail) {
			return v1.NewRequestError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: usr[%+v]: %w", usr, err)
	}

	return web.Respond(ctx, w, toAppUser(usr), http.StatusCreated)
}

// Update updates a user in the system.
// func (h *Handlers) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	var app AppUpdateUser
// 	if err := web.Decode(r, &app); err != nil {
// 		return err
// 	}

// 	userID := auth.GetUserID(ctx)

// 	usr, err := h.user.QueryByID(ctx, userID)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, user.ErrNotFound):
// 			return v1.NewRequestError(err, http.StatusNotFound)
// 		default:
// 			return fmt.Errorf("querybyid: userID[%s]: %w", userID, err)
// 		}
// 	}

// 	uu, err := toCoreUpdateUser(app)
// 	if err != nil {
// 		return v1.NewRequestError(err, http.StatusBadRequest)
// 	}

// 	usr, err = h.user.Update(ctx, usr, uu)
// 	if err != nil {
// 		return fmt.Errorf("update: userID[%s] uu[%+v]: %w", userID, uu, err)
// 	}

// 	return web.Respond(ctx, w, toAppUser(usr), http.StatusOK)
// }

// Delete removes a user from the system.
// func (h *Handlers) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	userID := auth.GetUserID(ctx)

// 	usr, err := h.user.QueryByID(ctx, userID)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, user.ErrNotFound):
// 			return web.Respond(ctx, w, nil, http.StatusNoContent)
// 		default:
// 			return fmt.Errorf("querybyid: userID[%s]: %w", userID, err)
// 		}
// 	}

// 	if err := h.user.Delete(ctx, usr); err != nil {
// 		return fmt.Errorf("delete: userID[%s]: %w", userID, err)
// 	}

// 	return web.Respond(ctx, w, nil, http.StatusNoContent)
// }

// Query returns a list of users with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, err := paging.ParseRequest(r)
	if err != nil {
		return err
	}

	filter, err := parseFilter(r)
	if err != nil {
		return err
	}

	orderBy, err := parseOrder(r)
	if err != nil {
		return err
	}

	users, err := h.user.Query(ctx, filter, orderBy, page.Number, page.RowsPerPage)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	items := make([]AppUser, len(users))
	for i, usr := range users {
		items[i] = toAppUser(usr)
	}

	total, err := h.user.Count(ctx, filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return web.Respond(ctx, w, paging.NewResponse(items, total, page.Number, page.RowsPerPage), http.StatusOK)
}

// QueryByID returns a user by its ID.
// func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 	id := auth.GetUserID(ctx)

// 	usr, err := h.user.QueryByID(ctx, id)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, user.ErrNotFound):
// 			return v1.NewRequestError(err, http.StatusNotFound)
// 		default:
// 			return fmt.Errorf("querybyid: id[%s]: %w", id, err)
// 		}
// 	}

// 	return web.Respond(ctx, w, toAppUser(usr), http.StatusOK)
// }
