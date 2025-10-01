package models

type ConfluentRoleBinding struct {
	APIVersion string          `yaml:"apiVersion"`
	Kind       string          `yaml:"kind"`
	Metadata   RoleBindingMeta `yaml:"metadata"`
	Spec       RoleBindingSpec `yaml:"spec"`
}

type RoleBindingMeta struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type RoleBindingSpec struct {
	KafkaRestClassRef Ref               `yaml:"kafkaRestClassRef"`
	Principal         Principal         `yaml:"principal"`
	Role              string            `yaml:"role"`
	ResourcePatterns  []ResourcePattern `yaml:"resourcePatterns"`
}

type Principal struct {
	Type string `yaml:"type"`
	Name string `yaml:"name"`
}

type ResourcePattern struct {
	Name         string `yaml:"name"`
	PatternType  string `yaml:"patternType"`
	ResourceType string `yaml:"resourceType"`
}

func NewRoleBinding(name, namespace, principalType, principalName, role, resourceName string) ConfluentRoleBinding {
	return ConfluentRoleBinding{
		APIVersion: "platform.confluent.io/v1beta1",
		Kind:       "ConfluentRoleBinding",
		Metadata: RoleBindingMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: RoleBindingSpec{
			KafkaRestClassRef: Ref{Name: "kafka-rest-class"},
			Principal: Principal{
				Type: principalType,
				Name: principalName,
			},
			Role: role,
			ResourcePatterns: []ResourcePattern{
				{
					Name:         resourceName,
					PatternType:  "PREFIXED",
					ResourceType: "Topic",
				},
			},
		},
	}
}
