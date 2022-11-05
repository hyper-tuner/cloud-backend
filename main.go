package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/custom/tunes/byTuneId/:tuneId",
			Handler: func(c echo.Context) error {
				record, _err := app.Dao().FindFirstRecordByData("tunes", "tuneId", c.PathParam("tuneId"))

				if _err != nil {
					return apis.NewNotFoundError("Tune not found", nil)
				}

				_errors := app.Dao().ExpandRecord(record, []string{"author"}, func(relCollection *models.Collection, relIds []string) ([]*models.Record, error) {
					record, _err := app.Dao().FindFirstRecordByData(relCollection.Name, "id", relIds[0])

					return []*models.Record{record}, _err
				})

				if len(_errors) > 0 {
					return apis.NewNotFoundError("Author not found", nil)
				}

				return c.JSON(http.StatusOK, record)
			},
		})

		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/custom/iniFiles/bySignature/:signature",
			Handler: func(c echo.Context) error {
				record, _err := app.Dao().FindFirstRecordByData("iniFiles", "signature", c.PathParam("signature"))

				if _err != nil {
					return c.JSON(http.StatusNotFound, _err)
				}

				return c.JSON(http.StatusOK, record)
			},
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
