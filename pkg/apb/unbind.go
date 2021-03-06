//
// Copyright (c) 2017 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package apb

import (
	"fmt"

	logging "github.com/op/go-logging"
	"github.com/openshift/ansible-service-broker/pkg/runtime"
)

// Unbind - runs the abp with the unbind action.
// TODO: Figure out the right way to allow apb to log
// It's passed in here, but that's a hard coupling point to
// github.com/op/go-logging, which is used all over the broker
// Maybe apb defines its own interface and accepts that optionally
// Little looser, but still not great
func Unbind(
	instance *ServiceInstance,
	parameters *Parameters,
	clusterConfig ClusterConfig,
	log *logging.Logger,
) error {
	log.Notice("============================================================")
	log.Notice("                       UNBINDING                            ")
	log.Notice("============================================================")
	log.Notice(fmt.Sprintf("ServiceInstance.ID: %s", instance.Spec.ID))
	log.Notice(fmt.Sprintf("ServiceInstance.Name: %v", instance.Spec.FQName))
	log.Notice(fmt.Sprintf("ServiceInstance.Image: %s", instance.Spec.Image))
	log.Notice(fmt.Sprintf("ServiceInstance.Description: %s", instance.Spec.Description))
	log.Notice("============================================================")

	executionContext, err := ExecuteApb(
		"unbind",
		clusterConfig,
		instance.Spec,
		instance.Context,
		parameters,
		log,
	)
	defer runtime.Provider.DestroySandbox(
		executionContext.PodName,
		executionContext.Namespace,
		executionContext.Targets,
		clusterConfig.Namespace,
		clusterConfig.KeepNamespace,
		clusterConfig.KeepNamespaceOnError,
	)
	if err != nil {
		log.Errorf("Problem executing apb [%s] unbind", executionContext.PodName)
		return err
	}

	err = watchPod(executionContext.PodName, executionContext.Namespace, log)
	if err != nil {
		log.Errorf("APB Execution failed - %v", err)
		return err
	}

	return nil
}
