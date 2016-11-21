package tools

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValueStoreSimple(t *testing.T) {
	tmpf, err := ioutil.TempFile("", "lfskeyvaluetest1")
	assert.Nil(t, err)
	filename := tmpf.Name()
	defer os.Remove(filename)
	tmpf.Close()

	kvs, err := NewKeyValueStore(filename)
	assert.Nil(t, err)

	// We'll include storing custom structs
	type customData struct {
		Val1 string
		Val2 int
	}
	// Needed to store custom struct
	RegisterTypeForKeyValueStorage(&customData{})

	kvs.Set("stringVal", "This is a string value")
	kvs.Set("intVal", 3)
	kvs.Set("floatVal", 3.142)
	kvs.Set("structVal", &customData{"structTest", 20})

	s := kvs.Get("stringVal")
	assert.Equal(t, "This is a string value", s)
	i := kvs.Get("intVal")
	assert.Equal(t, 3, i)
	f := kvs.Get("floatVal")
	assert.Equal(t, 3.142, f)
	c := kvs.Get("structVal")
	assert.Equal(t, c, &customData{"structTest", 20})
	n := kvs.Get("noValue")
	assert.Nil(t, n)

	kvs.Remove("stringVal")
	s = kvs.Get("stringVal")
	assert.Nil(t, s)
	// Set the string value again before saving
	kvs.Set("stringVal", "This is a string value")

	err = kvs.Save()
	assert.Nil(t, err)
	kvs = nil

	// Now confirm that we can read it all back
	kvs2, err := NewKeyValueStore(filename)
	assert.Nil(t, err)
	s = kvs2.Get("stringVal")
	assert.Equal(t, "This is a string value", s)
	i = kvs2.Get("intVal")
	assert.Equal(t, 3, i)
	f = kvs2.Get("floatVal")
	assert.Equal(t, 3.142, f)
	c = kvs2.Get("structVal")
	assert.Equal(t, c, &customData{"structTest", 20})
	n = kvs2.Get("noValue")
	assert.Nil(t, n)

}
