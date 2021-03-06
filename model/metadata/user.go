// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package metadata

import (
	"net"

	"github.com/elastic/beats/v7/libbeat/common"
)

type User struct {
	ID        string
	Email     string
	Name      string
	IP        net.IP
	UserAgent string
}

func (u *User) Fields() common.MapStr {
	if u == nil {
		return nil
	}
	var user mapStr
	user.maybeSetString("id", u.ID)
	user.maybeSetString("email", u.Email)
	user.maybeSetString("name", u.Name)
	return common.MapStr(user)
}

func (u *User) ClientFields() common.MapStr {
	if u == nil || u.IP == nil {
		return nil
	}
	return common.MapStr{"ip": u.IP.String()}
}

func (u *User) UserAgentFields() common.MapStr {
	if u == nil || u.UserAgent == "" {
		return nil
	}
	return common.MapStr{"original": u.UserAgent}
}
