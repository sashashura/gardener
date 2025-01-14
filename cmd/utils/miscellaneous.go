// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	coreinstall "github.com/gardener/gardener/pkg/apis/core/install"
	seedmanagementinstall "github.com/gardener/gardener/pkg/apis/seedmanagement/install"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	k8s "k8s.io/client-go/kubernetes"
	kubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"
)

// CreateRecorder creates a record.EventRecorder that is not limited to a namespace having a specific eventSourceName.
func CreateRecorder(kubeClient k8s.Interface, eventSourceName string) record.EventRecorder {
	scheme := runtime.NewScheme()

	utilruntime.Must(kubernetesscheme.AddToScheme(scheme))
	coreinstall.Install(scheme)
	seedmanagementinstall.Install(scheme)

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(1)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: typedcorev1.New(kubeClient.CoreV1().RESTClient()).Events("")})
	return eventBroadcaster.NewRecorder(scheme, corev1.EventSource{Component: eventSourceName})
}
