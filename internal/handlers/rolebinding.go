package handlers

import (
	"fmt"
	"net/http"

	"github.com/sebgott/event-creator/internal/models"

	"gopkg.in/yaml.v3"
)

func GenerateRoleBindingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	namespace := r.FormValue("namespace")
	principalType := r.FormValue("principalType")
	principalName := r.FormValue("principalName")
	role := r.FormValue("roleName")
	resName := r.FormValue("resourceName")

	rb := models.NewRoleBinding(name, namespace, principalType, principalName, role, resName)
	y, _ := yaml.Marshal(rb)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<pre>%s</pre>`, y)
}
