package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sebgott/event-creator/internal/models"

	"gopkg.in/yaml.v3"
)

func GenerateTopicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	namespace := r.FormValue("namespace")
	partitions, _ := strconv.Atoi(r.FormValue("partitions"))
	replicas, _ := strconv.Atoi(r.FormValue("replicas"))

	topic := models.NewKafkaTopic(name, namespace, partitions, replicas)
	y, _ := yaml.Marshal(topic)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<pre>%s</pre>`, y)
}
