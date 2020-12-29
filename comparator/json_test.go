package comparator

import (
	"bytes"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
)

/* TODO: Implement big files tests */

/* Example files */
const (
	InvalidPath     = "----------------"
	InvalidJSONFile = "test_examples/invalid.json"
	Empty           = "test_examples/empty.json"

	Simple              = "test_examples/simple.json"
	SimilarUnordered    = "test_examples/simple_unordered.json"
	DifferentValue      = "test_examples/simple_different_value.json"
	DifferentExtraKey   = "test_examples/simple_different_extra_key.json"
	DifferentMissingKey = "test_examples/simple_different_missing_key.json"

	MultipleTypes            = "test_examples/multiple_types.json"
	MultipleSimilarUnordered = "test_examples/multiple_types_unordered.json"
	MultipleDifferentValue   = "test_examples/multiple_types_different_value.json"
	MultipleMixedTypes       = "test_examples/multiple_types_mixed.json"

	BigFileObjects = 5000
	BigDefault     = "test_examples/big.json"
	BigDifferent   = "test_examples/big_different.json"
)

const charSet = "abcdedfghijklmnopqrst"

func generateObject() string {
	objectData := bytes.Buffer{}
	objectData.WriteString("{")

	object := map[string]interface{}{}
	for i := 0; i < rand.Intn(3)+1; i++ {
		switch rand.Intn(3) {
		case 0:
			object["id"] = rand.Intn(100000)
		case 1:
			object["letter"] = string(charSet[rand.Intn(len(charSet))])
		case 2:
			object["ratio"] = rand.Float64() * 100000
		}
	}

	first := true
	for key, val := range object {
		if !first {
			objectData.WriteString(",")
		}
		objectData.WriteString("\"")
		objectData.WriteString(key)
		objectData.WriteString("\":")
		switch reflect.TypeOf(val).Kind() {
		case reflect.String:
			objectData.WriteString("\"")
			objectData.WriteString(val.(string))
			objectData.WriteString("\"")
		case reflect.Float64:
			objectData.WriteString(strconv.FormatFloat(val.(float64), 'f', -1, 64))
		case reflect.Int:
			objectData.WriteString(strconv.Itoa(val.(int)))
		}

		first = false
	}

	objectData.WriteString("}")
	return objectData.String()
}

func generateJSONFile(objectsCount int, file string) error {
	fileData := bytes.Buffer{}
	fileData.WriteString("[")
	for i := 0; i < objectsCount; i++ {
		if i != 0 {
			fileData.WriteString(",")
		}
		fileData.WriteString(generateObject())
	}
	fileData.WriteString("]")

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	_, err = f.Write(fileData.Bytes())
	return err
}

/* Tests to package */

func TestAreSimilar(t *testing.T) {
	areSimilar, err := AreSimilar(Simple, Simple)
	if err != nil || !areSimilar {
		t.Error("Files should be similar")
	}

	areSimilar, err = AreSimilar(Simple, SimilarUnordered)
	if err != nil || !areSimilar {
		t.Error("Files should be similar even when not ordered")
	}

	areSimilar, err = AreSimilar(Simple, DifferentValue)
	if err != nil || areSimilar {
		t.Error("Files should be different when any value is different")
	}

	areSimilar, err = AreSimilar(Simple, DifferentExtraKey)
	if err != nil || areSimilar {
		t.Error("Files should be different when the second one has more keys")
	}

	areSimilar, err = AreSimilar(Simple, DifferentMissingKey)
	if err != nil || areSimilar {
		t.Error("Files should be different when a key is missing")
	}
}

func TestAreSimilarMultipleTypes(t *testing.T) {
	areSimilar, err := AreSimilar(MultipleTypes, MultipleTypes)
	if err != nil || !areSimilar {
		t.Error("Multiple types files should be equal")
	}

	areSimilar, err = AreSimilar(MultipleTypes, MultipleSimilarUnordered)
	if err != nil || !areSimilar {
		t.Error("Multiple types files in different order should be similar")
	}

	areSimilar, err = AreSimilar(MultipleTypes, MultipleDifferentValue)
	if err != nil || areSimilar {
		t.Error("Multiple types with different values should be different")
	}

	areSimilar, err = AreSimilar(MultipleTypes, MultipleMixedTypes)
	if err != nil || areSimilar {
		t.Error("Multiple types with different types should be different")
	}
}

func TestAreSimilarBig(t *testing.T) {
	/* Big files are generated during the test so we don't upload them to git */
	generateJSONFile(BigFileObjects, BigDefault)
	generateJSONFile(BigFileObjects, BigDifferent)
	defer os.Remove(BigDefault)
	defer os.Remove(BigDifferent)

	areSimilar, err := AreSimilar(BigDefault, BigDefault)
	if err != nil || !areSimilar {
		t.Error("Big files should be equal")
	}

	areSimilar, err = AreSimilar(BigDefault, BigDifferent)
	if err != nil || areSimilar {
		t.Error("Big files should be different")
	}
}

func TestAreSimilarEmpty(t *testing.T) {
	areSimilar, err := AreSimilar(Empty, Empty)
	if err != nil || !areSimilar {
		t.Error("Empty File should be equal to another empty file")
	}

	areSimilar, err = AreSimilar(Simple, Empty)
	if err != nil || areSimilar {
		t.Error("File should be different from the empty file")
	}

	areSimilar, err = AreSimilar(Simple, Empty)
	if err != nil || areSimilar {
		t.Error("Empty File should be different from the default example")
	}
}

func TestAreSimilarInvalid(t *testing.T) {
	_, err := AreSimilar(InvalidPath, Simple)
	if err == nil {
		t.Error("File on the first parameter should be an non existing path")
	}

	_, err = AreSimilar(Simple, InvalidPath)
	if err == nil {
		t.Error("File on the second parameter should be an non existing path")
	}

	_, err = AreSimilar(InvalidJSONFile, Simple)
	if err == nil {
		t.Error("File on the first parameter should have invalid json")
	}

	_, err = AreSimilar(Simple, InvalidJSONFile)
	if err == nil {
		t.Error("File on the second parameter should have invalid json")
	}
}
