package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var errEmailRequired error = errors.New("email is required")
var errFirstNameRequired error = errors.New("first name is required")
var errLastNameRequired error = errors.New("last name is required")
var errPasswordRequired error = errors.New("password is required")

type UserService struct {
	store Store
}

// Constructor
func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

// Methods
func (us *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", us.handleUserRegister).Methods("POST")
	r.HandleFunc("/users/login", us.handleLoginUser).Methods("POST")
}

func (us *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user *User
	err = json.Unmarshal(body, &user)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	if err := validateUserPayload(user); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	// antes de crear el usuario tenemos que encriptar la contraseña
	hashedPasswd, err := hashPassword(user.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating the user"})
		return
	}
	user.Password = hashedPasswd

	u, err := us.store.CreateUser(user)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating the user"})
		return
	}
	// luego de crear el usuario debemos crear un token para el mismo
	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating session"})
		return
	}

	WriteJSON(w, http.StatusCreated, token)	
}

func (us *UserService) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	// Traerse un usuario por email
	var loginUser LoginUser
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	u, err := us.store.GetUserByEmail(loginUser.Email)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrorResponse{Error: "user not found"})
		return
	}

	// Comparar la contraseña con la hash password
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginUser.Password)); err != nil {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid password"})
		return
	}
	// Crear un JWT y setearlo en una cookie
	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "error creating session"})
		return
	}
	// return la JWT como respuesta
	WriteJSON(w, http.StatusOK, token)
}

func validateUserPayload(u *User) error {
	if u.Email == "" {
		return errEmailRequired
	}

	if u.FirstName == "" {
		return errFirstNameRequired
	}

	if u.LastName == "" {
		return errLastNameRequired
	}

	if u.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", nil
	}

	http.SetCookie(w, &http.Cookie{
		Name: "Authorization",
		Value: token,
	})

	return token, nil
}