// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cluster "k8s.io/kube-deploy/ext-apiserver/pkg/apis/cluster"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_Cluster_To_cluster_Cluster,
		Convert_cluster_Cluster_To_v1alpha1_Cluster,
		Convert_v1alpha1_ClusterList_To_cluster_ClusterList,
		Convert_cluster_ClusterList_To_v1alpha1_ClusterList,
		Convert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec,
		Convert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec,
		Convert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus,
		Convert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus,
		Convert_v1alpha1_ClusterStatusStrategy_To_cluster_ClusterStatusStrategy,
		Convert_cluster_ClusterStatusStrategy_To_v1alpha1_ClusterStatusStrategy,
		Convert_v1alpha1_ClusterStrategy_To_cluster_ClusterStrategy,
		Convert_cluster_ClusterStrategy_To_v1alpha1_ClusterStrategy,
		Convert_v1alpha1_Machine_To_cluster_Machine,
		Convert_cluster_Machine_To_v1alpha1_Machine,
		Convert_v1alpha1_MachineList_To_cluster_MachineList,
		Convert_cluster_MachineList_To_v1alpha1_MachineList,
		Convert_v1alpha1_MachineSpec_To_cluster_MachineSpec,
		Convert_cluster_MachineSpec_To_v1alpha1_MachineSpec,
		Convert_v1alpha1_MachineStatus_To_cluster_MachineStatus,
		Convert_cluster_MachineStatus_To_v1alpha1_MachineStatus,
		Convert_v1alpha1_MachineStatusStrategy_To_cluster_MachineStatusStrategy,
		Convert_cluster_MachineStatusStrategy_To_v1alpha1_MachineStatusStrategy,
		Convert_v1alpha1_MachineStrategy_To_cluster_MachineStrategy,
		Convert_cluster_MachineStrategy_To_v1alpha1_MachineStrategy,
	)
}

func autoConvert_v1alpha1_Cluster_To_cluster_Cluster(in *Cluster, out *cluster.Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Cluster_To_cluster_Cluster is an autogenerated conversion function.
func Convert_v1alpha1_Cluster_To_cluster_Cluster(in *Cluster, out *cluster.Cluster, s conversion.Scope) error {
	return autoConvert_v1alpha1_Cluster_To_cluster_Cluster(in, out, s)
}

func autoConvert_cluster_Cluster_To_v1alpha1_Cluster(in *cluster.Cluster, out *Cluster, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_Cluster_To_v1alpha1_Cluster is an autogenerated conversion function.
func Convert_cluster_Cluster_To_v1alpha1_Cluster(in *cluster.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_cluster_Cluster_To_v1alpha1_Cluster(in, out, s)
}

func autoConvert_v1alpha1_ClusterList_To_cluster_ClusterList(in *ClusterList, out *cluster.ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.Cluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_ClusterList_To_cluster_ClusterList is an autogenerated conversion function.
func Convert_v1alpha1_ClusterList_To_cluster_ClusterList(in *ClusterList, out *cluster.ClusterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterList_To_cluster_ClusterList(in, out, s)
}

func autoConvert_cluster_ClusterList_To_v1alpha1_ClusterList(in *cluster.ClusterList, out *ClusterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Cluster)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_ClusterList_To_v1alpha1_ClusterList is an autogenerated conversion function.
func Convert_cluster_ClusterList_To_v1alpha1_ClusterList(in *cluster.ClusterList, out *ClusterList, s conversion.Scope) error {
	return autoConvert_cluster_ClusterList_To_v1alpha1_ClusterList(in, out, s)
}

func autoConvert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec(in *ClusterSpec, out *cluster.ClusterSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec is an autogenerated conversion function.
func Convert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec(in *ClusterSpec, out *cluster.ClusterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterSpec_To_cluster_ClusterSpec(in, out, s)
}

func autoConvert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec(in *cluster.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	return nil
}

// Convert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec is an autogenerated conversion function.
func Convert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec(in *cluster.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	return autoConvert_cluster_ClusterSpec_To_v1alpha1_ClusterSpec(in, out, s)
}

func autoConvert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus(in *ClusterStatus, out *cluster.ClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus is an autogenerated conversion function.
func Convert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus(in *ClusterStatus, out *cluster.ClusterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterStatus_To_cluster_ClusterStatus(in, out, s)
}

func autoConvert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus(in *cluster.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	return nil
}

// Convert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus is an autogenerated conversion function.
func Convert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus(in *cluster.ClusterStatus, out *ClusterStatus, s conversion.Scope) error {
	return autoConvert_cluster_ClusterStatus_To_v1alpha1_ClusterStatus(in, out, s)
}

func autoConvert_v1alpha1_ClusterStatusStrategy_To_cluster_ClusterStatusStrategy(in *ClusterStatusStrategy, out *cluster.ClusterStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_ClusterStatusStrategy_To_cluster_ClusterStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_ClusterStatusStrategy_To_cluster_ClusterStatusStrategy(in *ClusterStatusStrategy, out *cluster.ClusterStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterStatusStrategy_To_cluster_ClusterStatusStrategy(in, out, s)
}

func autoConvert_cluster_ClusterStatusStrategy_To_v1alpha1_ClusterStatusStrategy(in *cluster.ClusterStatusStrategy, out *ClusterStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_cluster_ClusterStatusStrategy_To_v1alpha1_ClusterStatusStrategy is an autogenerated conversion function.
func Convert_cluster_ClusterStatusStrategy_To_v1alpha1_ClusterStatusStrategy(in *cluster.ClusterStatusStrategy, out *ClusterStatusStrategy, s conversion.Scope) error {
	return autoConvert_cluster_ClusterStatusStrategy_To_v1alpha1_ClusterStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_ClusterStrategy_To_cluster_ClusterStrategy(in *ClusterStrategy, out *cluster.ClusterStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_ClusterStrategy_To_cluster_ClusterStrategy is an autogenerated conversion function.
func Convert_v1alpha1_ClusterStrategy_To_cluster_ClusterStrategy(in *ClusterStrategy, out *cluster.ClusterStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClusterStrategy_To_cluster_ClusterStrategy(in, out, s)
}

func autoConvert_cluster_ClusterStrategy_To_v1alpha1_ClusterStrategy(in *cluster.ClusterStrategy, out *ClusterStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_cluster_ClusterStrategy_To_v1alpha1_ClusterStrategy is an autogenerated conversion function.
func Convert_cluster_ClusterStrategy_To_v1alpha1_ClusterStrategy(in *cluster.ClusterStrategy, out *ClusterStrategy, s conversion.Scope) error {
	return autoConvert_cluster_ClusterStrategy_To_v1alpha1_ClusterStrategy(in, out, s)
}

func autoConvert_v1alpha1_Machine_To_cluster_Machine(in *Machine, out *cluster.Machine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_MachineSpec_To_cluster_MachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_MachineStatus_To_cluster_MachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Machine_To_cluster_Machine is an autogenerated conversion function.
func Convert_v1alpha1_Machine_To_cluster_Machine(in *Machine, out *cluster.Machine, s conversion.Scope) error {
	return autoConvert_v1alpha1_Machine_To_cluster_Machine(in, out, s)
}

func autoConvert_cluster_Machine_To_v1alpha1_Machine(in *cluster.Machine, out *Machine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_cluster_MachineSpec_To_v1alpha1_MachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_cluster_MachineStatus_To_v1alpha1_MachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_cluster_Machine_To_v1alpha1_Machine is an autogenerated conversion function.
func Convert_cluster_Machine_To_v1alpha1_Machine(in *cluster.Machine, out *Machine, s conversion.Scope) error {
	return autoConvert_cluster_Machine_To_v1alpha1_Machine(in, out, s)
}

func autoConvert_v1alpha1_MachineList_To_cluster_MachineList(in *MachineList, out *cluster.MachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]cluster.Machine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_MachineList_To_cluster_MachineList is an autogenerated conversion function.
func Convert_v1alpha1_MachineList_To_cluster_MachineList(in *MachineList, out *cluster.MachineList, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineList_To_cluster_MachineList(in, out, s)
}

func autoConvert_cluster_MachineList_To_v1alpha1_MachineList(in *cluster.MachineList, out *MachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Machine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_cluster_MachineList_To_v1alpha1_MachineList is an autogenerated conversion function.
func Convert_cluster_MachineList_To_v1alpha1_MachineList(in *cluster.MachineList, out *MachineList, s conversion.Scope) error {
	return autoConvert_cluster_MachineList_To_v1alpha1_MachineList(in, out, s)
}

func autoConvert_v1alpha1_MachineSpec_To_cluster_MachineSpec(in *MachineSpec, out *cluster.MachineSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_MachineSpec_To_cluster_MachineSpec is an autogenerated conversion function.
func Convert_v1alpha1_MachineSpec_To_cluster_MachineSpec(in *MachineSpec, out *cluster.MachineSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineSpec_To_cluster_MachineSpec(in, out, s)
}

func autoConvert_cluster_MachineSpec_To_v1alpha1_MachineSpec(in *cluster.MachineSpec, out *MachineSpec, s conversion.Scope) error {
	return nil
}

// Convert_cluster_MachineSpec_To_v1alpha1_MachineSpec is an autogenerated conversion function.
func Convert_cluster_MachineSpec_To_v1alpha1_MachineSpec(in *cluster.MachineSpec, out *MachineSpec, s conversion.Scope) error {
	return autoConvert_cluster_MachineSpec_To_v1alpha1_MachineSpec(in, out, s)
}

func autoConvert_v1alpha1_MachineStatus_To_cluster_MachineStatus(in *MachineStatus, out *cluster.MachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_MachineStatus_To_cluster_MachineStatus is an autogenerated conversion function.
func Convert_v1alpha1_MachineStatus_To_cluster_MachineStatus(in *MachineStatus, out *cluster.MachineStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineStatus_To_cluster_MachineStatus(in, out, s)
}

func autoConvert_cluster_MachineStatus_To_v1alpha1_MachineStatus(in *cluster.MachineStatus, out *MachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_cluster_MachineStatus_To_v1alpha1_MachineStatus is an autogenerated conversion function.
func Convert_cluster_MachineStatus_To_v1alpha1_MachineStatus(in *cluster.MachineStatus, out *MachineStatus, s conversion.Scope) error {
	return autoConvert_cluster_MachineStatus_To_v1alpha1_MachineStatus(in, out, s)
}

func autoConvert_v1alpha1_MachineStatusStrategy_To_cluster_MachineStatusStrategy(in *MachineStatusStrategy, out *cluster.MachineStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_MachineStatusStrategy_To_cluster_MachineStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_MachineStatusStrategy_To_cluster_MachineStatusStrategy(in *MachineStatusStrategy, out *cluster.MachineStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineStatusStrategy_To_cluster_MachineStatusStrategy(in, out, s)
}

func autoConvert_cluster_MachineStatusStrategy_To_v1alpha1_MachineStatusStrategy(in *cluster.MachineStatusStrategy, out *MachineStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_cluster_MachineStatusStrategy_To_v1alpha1_MachineStatusStrategy is an autogenerated conversion function.
func Convert_cluster_MachineStatusStrategy_To_v1alpha1_MachineStatusStrategy(in *cluster.MachineStatusStrategy, out *MachineStatusStrategy, s conversion.Scope) error {
	return autoConvert_cluster_MachineStatusStrategy_To_v1alpha1_MachineStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_MachineStrategy_To_cluster_MachineStrategy(in *MachineStrategy, out *cluster.MachineStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_MachineStrategy_To_cluster_MachineStrategy is an autogenerated conversion function.
func Convert_v1alpha1_MachineStrategy_To_cluster_MachineStrategy(in *MachineStrategy, out *cluster.MachineStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineStrategy_To_cluster_MachineStrategy(in, out, s)
}

func autoConvert_cluster_MachineStrategy_To_v1alpha1_MachineStrategy(in *cluster.MachineStrategy, out *MachineStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_cluster_MachineStrategy_To_v1alpha1_MachineStrategy is an autogenerated conversion function.
func Convert_cluster_MachineStrategy_To_v1alpha1_MachineStrategy(in *cluster.MachineStrategy, out *MachineStrategy, s conversion.Scope) error {
	return autoConvert_cluster_MachineStrategy_To_v1alpha1_MachineStrategy(in, out, s)
}
