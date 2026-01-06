package controller

import (
	"KageNoEn/lib"
	"KageNoEn/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var input model.UserInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, genErr := lib.GenerateId(input.Username)
	if genErr != nil {
		http.Error(w, genErr.Error(), http.StatusInternalServerError)
		return
	}

	eloRank, _ := c.R.GetbyElo(0)
	role, _ := c.R.GetRoleByLabel("player")
	status, _ := c.R.GetUserStatusByLabel("active")

	pssw, err := lib.HashPass(input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &model.User{
		Id:        res.Id,
		Username:  input.Username,
		Password:  pssw,
		Email:     input.Email,
		RankId:    eloRank.Id,
		RoleId:    role.Id,
		StatusId:  status.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := c.R.CreateUser(*user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := lib.GenerateToken(res.Id, role.Label, input.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   30 * 24 * 60 * 60,
	})

	secure := model.SecureUserRes{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Rank:      user.Rank,
		Status:    user.Status,
		Role:      user.Role,
		Elo:       user.Elo,
		IsOnline:  user.IsOnline,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	data := &model.RegisterResponse{
		Data: secure,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var input model.UserLoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.R.GetUserByUsername(input.Username)
	if err != nil {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}

	if !lib.CheckPass(input.Password, user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	token, err := lib.GenerateToken(user.Id, user.Role.Label, user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   30 * 24 * 60 * 60,
	})	

	user.IsOnline = true
	user.LastLogin = time.Now()
	_ = c.R.UpdateUser(user)

	secure := model.SecureUserRes{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Rank:      user.Rank,
		Status:    user.Status,
		Role:      user.Role,
		Elo:       user.Elo,
		IsOnline:  user.IsOnline,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.RegisterResponse{
		Data: secure,
	})
}

func (c *Controller) ChangePass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := c.R.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input model.ChangePassInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !lib.CheckPass(input.OldPass, user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	pssw, err := lib.HashPass(input.NewPass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Password = pssw
	if err := c.R.UpdateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := &model.UpdateMessage{
		Message: "Password changed successfully",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := c.R.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	user.IsOnline = false
	user.LastLogin = time.Now()
	_ = c.R.UpdateUser(user)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged out successfully",
	})
}

func (c *Controller) Session(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"connection": false})
		return
	}

	claims, err := lib.ValidateToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"connection": false})
		return
	}

	if claims.ExpiresAt < time.Now().Unix() {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"connection": false})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"connection": true})
}
