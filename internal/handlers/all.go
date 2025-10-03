package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sebgott/event-creator/internal/models"

	"gopkg.in/yaml.v3"
)

func GenerateAllHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	tmpl := r.FormValue("template")
	var topic models.KafkaTopic
	if preset, ok := models.TopicPresets[tmpl]; ok {
		topic = preset
	}

	topic.Metadata.Name = r.FormValue("topic_name")
	topic.Metadata.Namespace = r.FormValue("topic_namespace")
	topic.Spec.Name = topic.Metadata.Name

	partitions, _ := strconv.Atoi(r.FormValue("partitions"))
	replicas, _ := strconv.Atoi(r.FormValue("replicas"))
	if partitions > 0 {
		topic.Spec.PartitionCount = partitions
	}
	if replicas > 0 {
		topic.Spec.Replicas = replicas
	}

	keys := r.Form["config_key[]"]
	vals := r.Form["config_val[]"]
	if topic.Spec.Configs == nil {
		topic.Spec.Configs = make(map[string]string)
	}
	for i := range keys {
		if keys[i] != "" && vals[i] != "" {
			topic.Spec.Configs[keys[i]] = vals[i]
		}
	}

	topicYaml, _ := yaml.Marshal(topic)

	// RoleBinding
	resourceName := r.FormValue("resourceName")
	if resourceName == "" {
		resourceName = topic.Metadata.Name
	}

	rb := models.NewRoleBinding(
		"rb-"+topic.Metadata.Name,
		topic.Metadata.Namespace,
		r.FormValue("principalType"),
		r.FormValue("principalName"),
		r.FormValue("roleName"),
		resourceName,
	)
	rbYaml, _ := yaml.Marshal(rb)

	var out bytes.Buffer
	out.Write(topicYaml)
	out.WriteString("---\n")
	out.Write(rbYaml)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<pre>%s</pre>", out.String())
}
