/*
 Copyright 2016 Padduck, LLC

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

package writefile

import (
	"github.com/pufferpanel/pufferd/programs/operations/ops"
	"io/ioutil"

	"github.com/pufferpanel/apufferi/common"
	"github.com/pufferpanel/apufferi/logging"
	"github.com/pufferpanel/pufferd/environments"
)

type WriteFile struct {
	TargetFile string
	Text       string
}

func (c WriteFile) Run(env environments.Environment) error {
	logging.Debugf("Writing data to file: %s", c.TargetFile)
	env.DisplayToConsole("Writing some data to file: %s\n ", c.TargetFile)
	target := common.JoinPath(env.GetRootDirectory(), c.TargetFile)
	return ioutil.WriteFile(target, []byte(c.Text), 0644)
}

type WriteFileOperationFactory struct {
}

func (of WriteFileOperationFactory) Create(op ops.CreateOperation) ops.Operation {
	text := op.OperationArgs["text"].(string)
	target := op.OperationArgs["target"].(string)
	return WriteFile{TargetFile: target, Text: text}
}

func (of WriteFileOperationFactory) Key() string {
	return "writefile"
}


var Factory WriteFileOperationFactory