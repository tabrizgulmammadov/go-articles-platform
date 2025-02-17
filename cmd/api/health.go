package main

import (
	"net/http"
)

// HealthCheck godoc
//
//	@Summary		Health check endpoint
//	@Description	Returns the status of the application along with environment and version details.
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Application is running"
//	@Router			/v1/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
