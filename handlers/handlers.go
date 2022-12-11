package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"muzz/db"
	"net/http"
	"time"
)

type API struct {
	*echo.Echo
	*db.DB
}

func NewAPI(db *db.DB) *API {
	a := &API{
		echo.New(),
		db,
	}

	a.POST("/user/create", func(c echo.Context) error {
		if c.QueryParam("random") == "false" {
			return a.createProfile(c)
		}
		return a.generateRandomProfile(c)

	})
	a.GET("/profiles", a.getProfiles)
	a.POST("/swipe", a.getProfiles)

	return a
}

type CreateProfileParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}

func (a *API) createProfile(c echo.Context) error {
	p := &db.Profile{}
	err := a.CreateProfile(p)
	if err != nil {
		return err
	}
	return nil
}

func (a *API) getProfiles(c echo.Context) error {
	userID := c.QueryParam("user-id")
	p, err := a.GetProfiles(userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("could not get profiles: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, p)
}

func (a *API) swipe(c echo.Context) error {
	return nil
}

func (a *API) generateRandomProfile(c echo.Context) error {
	p := &db.Profile{
		Age:      rand.Intn(100),
		Gender:   randomString(1),
		Password: randomString(rand.Intn(32)),
		Name:     randomString(rand.Intn(40)),
		Email: fmt.Sprintf("%s@%s.com",
			randomString(rand.Intn(20)),
			randomString(rand.Intn(20))),
	}

	// store in db
	err := a.CreateProfile(p)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("error creating profile: %s", err.Error()))
	}

	return c.JSON(http.StatusCreated, p)
}

func randomString(n int) string {
	return "toto"
}

func generateID() int {
	return int(time.Now().UnixNano())
}
