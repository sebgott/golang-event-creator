package models

type KafkaTopic struct {
	APIVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   TopicMetadata `yaml:"metadata"`
	Spec       TopicSpec     `yaml:"spec"`
}

type TopicMetadata struct {
	Name        string            `yaml:"name"`
	Namespace   string            `yaml:"namespace"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}

type TopicSpec struct {
	Name              string            `yaml:"name"`
	Replicas          int               `yaml:"replicas"`
	PartitionCount    int               `yaml:"partitionCount"`
	KafkaRestClassRef Ref               `yaml:"kafkaRestClassRef"`
	Configs           map[string]string `yaml:"configs,omitempty"`
}

type Ref struct {
	Name string `yaml:"name"`
}

func NewKafkaTopic(name, namespace string, partitions, replicas int) KafkaTopic {
	return KafkaTopic{
		APIVersion: "platform.confluent.io/v1beta1",
		Kind:       "KafkaTopic",
		Metadata: TopicMetadata{
			Name:      name,
			Namespace: namespace,
		},
		Spec: TopicSpec{
			Name:              name,
			Replicas:          replicas,
			PartitionCount:    partitions,
			KafkaRestClassRef: Ref{Name: "kafka-rest-class"},
		},
	}
}

var TopicPresets = map[string]KafkaTopic{
	"baseAnnotations": {
		APIVersion: "platform.confluent.io/v1beta1",
		Kind:       "KafkaTopic",
		Metadata: TopicMetadata{
			Annotations: map[string]string{
				"owner":       "owner",
				"environment": "environment",
				"comment":     "event description",
			},
		},
		Spec: TopicSpec{
			Replicas:          3,
			PartitionCount:    1,
			KafkaRestClassRef: Ref{Name: "kafka-rest-class"},
			Configs: map[string]string{
				"min.insync.replicas": "2",
				"max.message.bytes":   "1048588",
				"cleanup.policy":      "delete",
				"compression.type":    "gzip",
			},
		},
	},
}
