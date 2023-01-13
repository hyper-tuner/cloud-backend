package main

import (
	"log"
	"net/http"

	_ "main/migrations"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	tunesCollectionName := "tunes"
	iniFilesCollectionName := "iniFiles"
	stargazersCollectionName := "stargazers"

	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/custom/tunes/byTuneId/:tuneId",
			Handler: func(c echo.Context) error {
				record, _err := app.Dao().FindFirstRecordByData(tunesCollectionName, "tuneId", c.PathParam("tuneId"))

				if _err != nil {
					return apis.NewNotFoundError("Tune not found", nil)
				}

				apis.EnrichRecord(c, app.Dao(), record, "author")

				return c.JSON(http.StatusOK, record)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/custom/iniFiles/bySignature/:signature",
			Handler: func(c echo.Context) error {
				record, _err := app.Dao().FindFirstRecordByData(iniFilesCollectionName, "signature", c.PathParam("signature"))

				if _err != nil {
					return c.JSON(http.StatusNotFound, _err)
				}

				return c.JSON(http.StatusOK, record)
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodPost,
			Path:   "/api/custom/stargazers/toggleStar",
			Handler: func(c echo.Context) (err error) {
				auth := c.Get(apis.ContextAuthRecordKey).(*models.Record)
				isStarred := false

				type Stargazer struct {
					User string
					Tune string `json:"tune" form:"tune" query:"tune"`
				}

				stargazer := new(Stargazer)

				if err = c.Bind(stargazer); err != nil {
					return c.String(http.StatusBadRequest, "Bad request")
				}

				stargazer.User = auth.Id

				_err := app.Dao().RunInTransaction(func(txDao *daos.Dao) error {
					collection, err := app.Dao().FindCollectionByNameOrId(stargazersCollectionName)
					if err != nil {
						return err
					}

					stargazerRecord := models.NewRecord(collection)
					stargazerRecord.Set("user", stargazer.User)
					stargazerRecord.Set("tune", stargazer.Tune)

					tune, _err := app.Dao().FindFirstRecordByData(tunesCollectionName, "id", stargazer.Tune)
					if _err != nil {
						return apis.NewNotFoundError("Tune not found", nil)
					}

					_err = txDao.SaveRecord(stargazerRecord)

					// UNIQUE constraint failed most likely, try to unstar
					if _err != nil {
						currentStargazerRecords, _err := app.Dao().FindRecordsByExpr(
							stargazersCollectionName,
							dbx.HashExp{"user": stargazer.User},
							dbx.HashExp{"tune": stargazer.Tune},
						)

						if _err != nil || len(currentStargazerRecords) == 0 {
							return _err
						}

						_err = txDao.DeleteRecord(currentStargazerRecords[0])
						if _err != nil {
							return _err
						}

						isStarred = false
						tune.Set("stars", tune.Get("stars").(float64)-1)
					} else {
						isStarred = true
						tune.Set("stars", tune.Get("stars").(float64)+1)
					}

					_err = txDao.SaveRecord(tune)
					if _err != nil {
						return _err
					}

					return nil
				})

				if _err != nil {
					return apis.NewNotFoundError("Tune not found or already starred", nil)
				}

				// fetch again and return current state
				tune, _err := app.Dao().FindFirstRecordByData(tunesCollectionName, "id", stargazer.Tune)

				return c.JSON(http.StatusOK, map[string]any{
					"stars":     tune.Get("stars").(float64),
					"isStarred": isStarred,
				})
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				apis.LoadAuthContext(app.App),
				apis.RequireAdminOrRecordAuth("users"),
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/custom/stargazers/starredByMe/:tune",
			Handler: func(c echo.Context) (err error) {
				auth := c.Get(apis.ContextAuthRecordKey).(*models.Record)
				isStarred := true

				type Stargazer struct {
					User string
					Tune string
				}

				stargazer := new(Stargazer)

				stargazer.User = auth.Id
				stargazer.Tune = c.PathParam("tune")

				stargazerRecords, _err := app.Dao().FindRecordsByExpr(
					stargazersCollectionName,
					dbx.HashExp{"user": stargazer.User},
					dbx.HashExp{"tune": stargazer.Tune},
				)

				if _err != nil || len(stargazerRecords) == 0 {
					isStarred = false
				}

				return c.JSON(http.StatusOK, map[string]any{
					"isStarred": isStarred,
				})
			},
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(app),
				apis.LoadAuthContext(app.App),
				apis.RequireAdminOrRecordAuth("users"),
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
