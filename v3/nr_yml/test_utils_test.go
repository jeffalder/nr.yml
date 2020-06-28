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

