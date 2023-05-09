package middlewares

var interceptedRoutes = map[string]string{
	"GET/healthz": "public",

	"POST/user":       "create",
	"GET/user":        "read",
	"GET/user/:id":    "read",
	"PATCH/user/:id":  "update",
	"DELETE/user/:id": "delete",
}
