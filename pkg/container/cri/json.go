/*
Copyright 2019 The Kubernetes Authors.

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

package cri

import (
	"encoding/json"
	"fmt"
)

/*
Custom JSON / yaml (by way of json) serialization for these types
*/

func (m *Mount) MarshalJSON() ([]byte, error) {
	type Alias Mount
	name, ok := MountPropagation_name[int32(m.Propagation)]
	if !ok {
		return nil, fmt.Errorf("unknown propogation value: %v", m.Propagation)
	}
	return json.Marshal(&struct {
		Propagation string `json:"propagation"`
		*Alias
	}{
		Propagation: name,
		Alias:       (*Alias)(m),
	})
}

func (m *Mount) UnmarshalJSON(data []byte) error {
	type Alias Mount
	aux := &struct {
		Propagation string `json:"propagation"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	// if unset, will fallback to the default (0)
	if aux.Propagation != "" {
		val, ok := MountPropagation_value[aux.Propagation]
		if !ok {
			return fmt.Errorf("unknown propogation value: %s", aux.Propagation)
		}
		m.Propagation = MountPropagation(val)
	}
	return nil
}
