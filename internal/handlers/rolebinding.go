package handlers

import (
	"fmt"
	"net/http"

	"github.com/sebgott/event-creator/internal/models"

	"sigs.k8s.io/yaml"
)

func GenerateRoleBindingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	principal := r.FormValue("principal")
	role := r.FormValue("roleName")
	resType := r.FormValue("resourceType")
	resName := r.FormValue("resourceName")

	rb := models.NewRoleBinding(name, principal, role, resType, resName)
	y, _ := yaml.Marshal(rb)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<pre>%s</pre>`, y)
}
