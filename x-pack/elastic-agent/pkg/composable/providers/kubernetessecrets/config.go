// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package kubernetessecrets

// Config for kubernetes provider
type Config struct {
	KubeConfig string `config:"kube_config"`
}
