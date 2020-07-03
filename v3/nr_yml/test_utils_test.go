//   Copyright 2020, Jeff Alder
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
package nr_yml

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func withContents(contents string, t *testing.T, test func(filename string, t *testing.T)) {
	tempFile, err := ioutil.TempFile(os.TempDir(), t.Name()+"*")
	assert.NoError(t, err)

	defer os.Remove(tempFile.Name())

	_, err = tempFile.WriteString(contents)
	assert.NoError(t, err)

	test(tempFile.Name(), t)
}

