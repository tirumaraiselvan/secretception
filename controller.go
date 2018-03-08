/*
Copyright 2017 The Kubernetes Authors.

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

package main

import (
	"crypto/md5"
	"encoding/hex"
	l "github.com/sirupsen/logrus"
	// appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	// kerr "k8s.io/apimachinery/pkg/api/errors"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// appsv1beta2 "k8s.io/api/apps/v1beta2"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
)

type SecretController struct {
	kubeclientset *kubernetes.Clientset
}

func NewController(
	kubeclientset *kubernetes.Clientset,
	kubeInformerFactory kubeinformers.SharedInformerFactory,
) *SecretController {

	secretInformer := kubeInformerFactory.Core().V1().Secrets()

	controller := &SecretController{
		kubeclientset: kubeclientset,
	}

	l.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	secretInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			object := obj.(*corev1.Secret)
			if object.Namespace == "kube-system" {
				l.Info("kube-system secret")
				return

			}

			l.Info("added secret")
			l.Infof("%#v", object.ObjectMeta)

			value := object.Data["password"]
			object.Data["password"] = []byte(GetMD5Hash(string(value)))
			_, err := controller.kubeclientset.CoreV1().Secrets(object.Namespace).Update(object)
			if err != nil {
				l.Error("error updating")
			}

			// controller.enqueue(obj)
		},
		UpdateFunc: func(old, new interface{}) {
			oldSecret := old.(*corev1.Secret)
			newSecret := new.(*corev1.Secret)
			if oldSecret.ResourceVersion == newSecret.ResourceVersion {
				return
			}

			if newSecret.Namespace == "kube-system" {
				l.Info("kube-system secret")
				return

			}

			l.Info("updated secret")
			l.Infof("%#v", newSecret.ObjectMeta)

		},
	})
	return controller
}

func (c *SecretController) Run(stop <-chan struct{}) {
	l.Print("waiting for cache sync")
	if !cache.WaitForCacheSync(stop) {
		l.Print("timed out waiting for cache sync")
		return
	}
	l.Print("caches are synced")

	// wait until we're told to stop
	l.Print("waiting for stop signal")
	<-stop
	l.Print("received stop signal")
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
