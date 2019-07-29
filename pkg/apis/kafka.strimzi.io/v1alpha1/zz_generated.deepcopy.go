// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kafka) DeepCopyInto(out *Kafka) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kafka.
func (in *Kafka) DeepCopy() *Kafka {
	if in == nil {
		return nil
	}
	out := new(Kafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kafka) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaList) DeepCopyInto(out *KafkaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kafka, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaList.
func (in *KafkaList) DeepCopy() *KafkaList {
	if in == nil {
		return nil
	}
	out := new(KafkaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KafkaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaListener) DeepCopyInto(out *KafkaListener) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaListener.
func (in *KafkaListener) DeepCopy() *KafkaListener {
	if in == nil {
		return nil
	}
	out := new(KafkaListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpec) DeepCopyInto(out *KafkaSpec) {
	*out = *in
	in.Kafka.DeepCopyInto(&out.Kafka)
	out.Zookeeper = in.Zookeeper
	out.EntityOperator = in.EntityOperator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpec.
func (in *KafkaSpec) DeepCopy() *KafkaSpec {
	if in == nil {
		return nil
	}
	out := new(KafkaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpecEntityOperator) DeepCopyInto(out *KafkaSpecEntityOperator) {
	*out = *in
	out.TopicOperator = in.TopicOperator
	out.UserOperator = in.UserOperator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpecEntityOperator.
func (in *KafkaSpecEntityOperator) DeepCopy() *KafkaSpecEntityOperator {
	if in == nil {
		return nil
	}
	out := new(KafkaSpecEntityOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpecKafka) DeepCopyInto(out *KafkaSpecKafka) {
	*out = *in
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make(map[string]KafkaListener, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Config = in.Config
	out.Storage = in.Storage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpecKafka.
func (in *KafkaSpecKafka) DeepCopy() *KafkaSpecKafka {
	if in == nil {
		return nil
	}
	out := new(KafkaSpecKafka)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpecKafkaConfig) DeepCopyInto(out *KafkaSpecKafkaConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpecKafkaConfig.
func (in *KafkaSpecKafkaConfig) DeepCopy() *KafkaSpecKafkaConfig {
	if in == nil {
		return nil
	}
	out := new(KafkaSpecKafkaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaSpecZookeeper) DeepCopyInto(out *KafkaSpecZookeeper) {
	*out = *in
	out.Storage = in.Storage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaSpecZookeeper.
func (in *KafkaSpecZookeeper) DeepCopy() *KafkaSpecZookeeper {
	if in == nil {
		return nil
	}
	out := new(KafkaSpecZookeeper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaStatus) DeepCopyInto(out *KafkaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaStatus.
func (in *KafkaStatus) DeepCopy() *KafkaStatus {
	if in == nil {
		return nil
	}
	out := new(KafkaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaStorage) DeepCopyInto(out *KafkaStorage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaStorage.
func (in *KafkaStorage) DeepCopy() *KafkaStorage {
	if in == nil {
		return nil
	}
	out := new(KafkaStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaTopicOperator) DeepCopyInto(out *KafkaTopicOperator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaTopicOperator.
func (in *KafkaTopicOperator) DeepCopy() *KafkaTopicOperator {
	if in == nil {
		return nil
	}
	out := new(KafkaTopicOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KafkaUserOperator) DeepCopyInto(out *KafkaUserOperator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KafkaUserOperator.
func (in *KafkaUserOperator) DeepCopy() *KafkaUserOperator {
	if in == nil {
		return nil
	}
	out := new(KafkaUserOperator)
	in.DeepCopyInto(out)
	return out
}
