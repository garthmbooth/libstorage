package tasks

import (
	"github.com/akutz/gofig"
	"github.com/codedellemc/libstorage/api/registry"
	"github.com/codedellemc/libstorage/api/server/httputils"
	"github.com/codedellemc/libstorage/api/types"
)

func init() {
	registry.RegisterRouter(&router{})
}

type router struct {
	routes []types.Route
}

func (r *router) Name() string {
	return "tasks-router"
}

func (r *router) Init(config gofig.Config) {
	r.initRoutes()
}

// Routes returns the available routes.
func (r *router) Routes() []types.Route {
	return r.routes
}

func (r *router) initRoutes() {

	r.routes = []types.Route{

		// GET
		httputils.NewGetRoute(
			"tasks",
			"/tasks",
			r.tasks),

		// GET
		httputils.NewGetRoute(
			"taskInspect",
			"/tasks/{taskID}",
			r.taskInspect),
	}
}
