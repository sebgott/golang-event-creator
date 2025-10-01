package models

type ConfluentRoleBinding struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		Principal    string `yaml:"principal"`
		RoleName     string `yaml:"roleName"`
		ResourceType string `yaml:"resourceType"`
		ResourceName string `yaml:"resourceName"`
	} `yaml:"spec"`
}

func NewRoleBinding(name, principal, roleName, resType, resName string) ConfluentRoleBinding {
	rb := ConfluentRoleBinding{
		APIVersion: "platform.confluent.io/v1beta1",
		Kind:       "ConfluentRoleBinding",
	}
	rb.Metadata.Name = name
	rb.Spec.Principal = principal
	rb.Spec.RoleName = roleName
	rb.Spec.ResourceType = resType
	rb.Spec.ResourceName = resName
	return rb
}
