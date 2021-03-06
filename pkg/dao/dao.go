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

package dao

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	logging "github.com/op/go-logging"
	"github.com/openshift/ansible-service-broker/pkg/apb"
	"github.com/openshift/ansible-service-broker/pkg/clients"
	"github.com/pborman/uuid"
)

// Config - contains dao configuration
type Config struct {
	EtcdHost       string `yaml:"etcd_host"`
	EtcdPort       string `yaml:"etcd_port"`
	EtcdCaFile     string `yaml:"etcd_ca_file"`
	EtcdClientCert string `yaml:"etcd_client_cert"`
	EtcdClientKey  string `yaml:"etcd_client_key"`
}

// GetEtcdConfig - Simple EtcdConfig getter
func (c Config) GetEtcdConfig() clients.EtcdConfig {
	return clients.EtcdConfig{EtcdHost: c.EtcdHost,
		EtcdPort:       c.EtcdPort,
		EtcdClientCert: c.EtcdClientCert,
		EtcdCaFile:     c.EtcdCaFile,
		EtcdClientKey:  c.EtcdClientKey}
}

// Dao - object to interface with the data store.
type Dao struct {
	config Config
	log    *logging.Logger
	client client.Client
	kapi   client.KeysAPI // Used to interact with kvp API over HTTP
}

// NewDao - Create a new Dao object
func NewDao(config Config, log *logging.Logger) (*Dao, error) {
	dao := Dao{
		config: config,
		log:    log,
	}

	etcdClient, err := clients.Etcd(config.GetEtcdConfig(), log)
	if err != nil {
		return nil, err
	}
	dao.client = etcdClient
	dao.kapi = client.NewKeysAPI(dao.client)
	return &dao, nil
}

// SetRaw - Allows the setting of the value json string to the key in the kvp API.
func (d *Dao) SetRaw(key string, val string) error {
	_, err := d.kapi.Set(context.Background(), key, val /*opts*/, nil)
	return err
}

// GetRaw - gets a specific json string for a key from the kvp API.
func (d *Dao) GetRaw(key string) (string, error) {
	res, err := d.kapi.Get(context.Background(), key /*opts*/, nil)
	if err != nil {
		return "", err
	}

	val := res.Node.Value
	return val, nil
}

// BatchGetRaw - Get multiple  types as individual json strings
// TODO: Streaming interface? Going to need to optimize all this for
// a full-load catalog response of 10k
// This is more likely to be paged given current proposal
// In which case, we need paged Batch gets
// 2 steps?
// GET /spec/manifest [/*ordered ids*/]
// BatchGet(offset, count)?
func (d *Dao) BatchGetRaw(dir string) (*[]string, error) {
	d.log.Debug("Dao::BatchGetRaw")

	var res *client.Response
	var err error

	opts := &client.GetOptions{Recursive: true}
	if res, err = d.kapi.Get(context.Background(), dir, opts); err != nil {
		return nil, err
	}

	specNodes := res.Node.Nodes
	specCount := len(specNodes)

	d.log.Debug("Successfully loaded [ %d ] objects from etcd dir [ %s ]", specCount, dir)

	payloads := make([]string, specCount)
	for i, node := range specNodes {
		payloads[i] = node.Value
	}

	return &payloads, nil
}

// GetSpec - Retrieve the spec for the kvp API.
func (d *Dao) GetSpec(id string) (*apb.Spec, error) {
	spec := &apb.Spec{}
	if err := d.getObject(specKey(id), spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// SetSpec - set spec for an id in the kvp API.
func (d *Dao) SetSpec(id string, spec *apb.Spec) error {
	return d.setObject(specKey(id), spec)
}

// DeleteSpec - Delete the spec for a given spec id.
func (d *Dao) DeleteSpec(specID string) error {
	d.log.Debug(fmt.Sprintf("Dao::DeleteSpec-> [ %s ]", specID))
	_, err := d.kapi.Delete(context.Background(), specKey(specID), nil)
	return err
}

// BatchSetSpecs - set specs based on SpecManifest in the kvp API.
func (d *Dao) BatchSetSpecs(specs apb.SpecManifest) error {
	// TODO: Is there no batch insert in the etcd api?
	for id, spec := range specs {
		err := d.SetSpec(id, spec)
		if err != nil {
			return err
		}
	}

	return nil
}

// BatchGetSpecs - Retrieve all the specs for dir.
func (d *Dao) BatchGetSpecs(dir string) ([]*apb.Spec, error) {
	payloads, err := d.BatchGetRaw(dir)

	if client.IsKeyNotFound(err) {
		return []*apb.Spec{}, nil
	} else if err != nil {
		return []*apb.Spec{}, err
	}

	specs := make([]*apb.Spec, len(*payloads))
	for i, payload := range *payloads {
		spec := &apb.Spec{}
		apb.LoadJSON(payload, spec)
		specs[i] = spec
		d.log.Debug("Batch idx [ %d ] -> [ %s ]", i, spec.ID)
	}

	return specs, nil
}

// BatchDeleteSpecs - set specs based on SpecManifest in the kvp API.
func (d *Dao) BatchDeleteSpecs(specs []*apb.Spec) error {
	// TODO: Is there no batch insert in the etcd api?
	for _, spec := range specs {
		err := d.DeleteSpec(spec.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// FindJobStateByState - Retrieve all the jobs that match the specified state
func (d *Dao) FindJobStateByState(state apb.State) ([]apb.RecoverStatus, error) {
	d.log.Debug("Dao::FindByState")

	var res *client.Response
	var err error

	opts := &client.GetOptions{Recursive: true}
	if res, err = d.kapi.Get(context.Background(), "/state/", opts); err != nil {
		return nil, err
	}

	stateNodes := res.Node.Nodes
	stateCount := len(stateNodes)

	d.log.Debug("Successfully loaded [ %d ] jobstate objects from etcd dir [ /state/ ]", stateCount)

	var recoverstatus []apb.RecoverStatus
	for _, node := range stateNodes {
		k := fmt.Sprintf("%s/job", node.Key)

		status := apb.RecoverStatus{InstanceID: uuid.Parse(stateKeyID(node.Key))}
		jobstate := apb.JobState{}
		nodes, e := d.kapi.Get(context.Background(), k, opts)
		if e != nil {
			// if key is invalid do we keep it?
			d.log.Warning(
				fmt.Sprintf("Error processing jobstate record, moving on to next. %v", e.Error()))
			continue
		}

		for _, n := range nodes.Node.Nodes {
			apb.LoadJSON(n.Value, &jobstate)
			if jobstate.State == state {
				d.log.Debug(fmt.Sprintf(
					"Found! jobstate [%v] matched given state: [%v].", jobstate, state))
				status.State = jobstate
				recoverstatus = append(recoverstatus, status)
			} else {
				// we could probably remove this once we're happy with how this
				// works.
				d.log.Debug(fmt.Sprintf(
					"Skipping, jobstate [%v] did not match given state: [%v].", jobstate, state))
			}
		}
	}

	return recoverstatus, nil
}

// GetSvcInstJobsByState - Lookup all jobs of a given state for a specific instance
func (d *Dao) GetSvcInstJobsByState(
	instanceID string, reqState apb.State,
) ([]apb.JobState, error) {
	d.log.Debug("Dao::GetSvcInstJobsByState")
	allStates, err := d.getJobsForSvcInst(instanceID)

	if err != nil {
		return nil, err
	} else if len(allStates) == 0 {
		return allStates, nil
	}

	filtStates := []apb.JobState{}
	for _, state := range allStates {
		if state.State == reqState {
			filtStates = append(filtStates, state)
		}
	}

	d.log.Debugf("Filtered on state: [ %v ], returning %d jobs", reqState, len(filtStates))

	return filtStates, nil
}

func (d *Dao) getJobsForSvcInst(instanceID string) ([]apb.JobState, error) {
	d.log.Debug("Dao::getJobsForSvcInst")

	var res *client.Response
	var err error

	lookupKey := fmt.Sprintf("/state/%s/job", instanceID)
	opts := &client.GetOptions{Recursive: true}
	if res, err = d.kapi.Get(context.Background(), lookupKey, opts); err != nil {
		return nil, err
	}

	jobNodes := res.Node.Nodes
	jobsCount := len(jobNodes)
	if jobsCount == 0 {
		return []apb.JobState{}, nil
	}

	d.log.Debug("Successfully loaded [ %d ] jobs objects from [ %s ]",
		jobsCount, lookupKey)

	retJobs := []apb.JobState{}
	for _, node := range jobNodes {
		js := apb.JobState{}
		err := apb.LoadJSON(node.Value, &js)
		if err != nil {
			return nil, fmt.Errorf("An error occurred trying to parse job state of [ %s ]\n%s", node.Key, err.Error())
		}
		retJobs = append(retJobs, js)
	}
	return retJobs, nil
}

// GetServiceInstance - Retrieve specific service instance from the kvp API.
func (d *Dao) GetServiceInstance(id string) (*apb.ServiceInstance, error) {
	spec := &apb.ServiceInstance{}
	if err := d.getObject(serviceInstanceKey(id), spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// SetServiceInstance - Set service instance for an id in the kvp API.
func (d *Dao) SetServiceInstance(id string, serviceInstance *apb.ServiceInstance) error {
	return d.setObject(serviceInstanceKey(id), serviceInstance)
}

// DeleteServiceInstance - Delete the service instance for an service instance id.
func (d *Dao) DeleteServiceInstance(id string) error {
	d.log.Debug(fmt.Sprintf("Dao::DeleteServiceInstance -> [ %s ]", id))
	_, err := d.kapi.Delete(context.Background(), serviceInstanceKey(id), nil)
	return err
}

// GetBindInstance - Retrieve a specific bind instance from the kvp API
func (d *Dao) GetBindInstance(id string) (*apb.BindInstance, error) {
	spec := &apb.BindInstance{}
	if err := d.getObject(bindInstanceKey(id), spec); err != nil {
		return nil, err
	}
	return spec, nil
}

// SetBindInstance - Set the bind instance for id in the kvp API.
func (d *Dao) SetBindInstance(id string, bindInstance *apb.BindInstance) error {
	return d.setObject(bindInstanceKey(id), bindInstance)
}

// DeleteBindInstance - Delete the binding instance for an id in the kvp API.
func (d *Dao) DeleteBindInstance(id string) error {
	d.log.Debug(fmt.Sprintf("Dao::DeleteBindInstance -> [ %s ]", id))
	_, err := d.kapi.Delete(context.Background(), bindInstanceKey(id), nil)
	return err
}

// GetExtractedCredentials - Get the extracted credentials for an id in the kvp API.
func (d *Dao) GetExtractedCredentials(id string) (*apb.ExtractedCredentials, error) {
	extractedCredentials := &apb.ExtractedCredentials{}
	if err := d.getObject(extractedCredentialsKey(id), extractedCredentials); err != nil {
		return nil, err
	}
	return extractedCredentials, nil
}

// SetExtractedCredentials - Set the extracted credentials for an id in the kvp API.
func (d *Dao) SetExtractedCredentials(id string, extractedCredentials *apb.ExtractedCredentials) error {
	return d.setObject(extractedCredentialsKey(id), extractedCredentials)
}

// DeleteExtractedCredentials - delete the extracted credentials for an id in the kvp API.
func (d *Dao) DeleteExtractedCredentials(id string) error {
	_, err := d.kapi.Delete(context.Background(), extractedCredentialsKey(id), nil)
	return err
}

// SetState - Set the Job State in the kvp API for id.
func (d *Dao) SetState(id string, state apb.JobState) error {
	return d.setObject(stateKey(id, state.Token), state)
}

// GetState - Retrieve a job state from the kvp API for an ID and Token.
func (d *Dao) GetState(id string, token string) (apb.JobState, error) {
	state := apb.JobState{}
	if err := d.getObject(stateKey(id, token), &state); err != nil {
		return apb.JobState{State: apb.StateFailed}, err
	}
	return state, nil
}

// BatchSetPlanNames - set plannames based on PlanNameManifest in the kvp API.
func (d *Dao) BatchSetPlanNames(planNames map[string]string) error {
	for id, planName := range planNames {
		err := d.SetPlanName(id, planName)
		if err != nil {
			return err
		}
	}

	return nil
}

// SetPlanName - Set the Plan name in the kvp API for the given id.
func (d *Dao) SetPlanName(id string, name string) error {
	// it's just a string so no other work is needed
	return d.SetRaw(planNameKey(id), name)
}

// GetPlanName - Retrieve the plan name associated with the ID
func (d *Dao) GetPlanName(id string) (string, error) {
	// it's just a string so no other work is needed
	return d.GetRaw(planNameKey(id))
}

func (d *Dao) getObject(key string, data interface{}) error {
	raw, err := d.GetRaw(key)
	if err != nil {
		return err
	}
	apb.LoadJSON(raw, data)
	return nil
}

func (d *Dao) setObject(key string, data interface{}) error {
	payload, err := apb.DumpJSON(data)
	if err != nil {
		return err
	}
	return d.SetRaw(key, payload)
}

////////////////////////////////////////////////////////////
// Key generators
////////////////////////////////////////////////////////////

func stateKey(id string, jobid string) string {
	return fmt.Sprintf("/state/%s/job/%s", id, jobid)
}

func stateKeyID(key string) string {
	s := strings.TrimPrefix(key, "/state/")
	s = strings.TrimSuffix(s, "/job")
	return s
}

func extractedCredentialsKey(id string) string {
	return fmt.Sprintf("/extracted_credentials/%s", id)
}

func specKey(id string) string {
	return fmt.Sprintf("/spec/%s", id)
}

func serviceInstanceKey(id string) string {
	return fmt.Sprintf("/service_instance/%s", id)
}

func bindInstanceKey(id string) string {
	return fmt.Sprintf("/bind_instance/%s", id)
}

func planNameKey(id string) string {
	return fmt.Sprintf("/plan_name/%s", id)
}
