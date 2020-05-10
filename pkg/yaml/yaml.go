package yaml

import (
	"io/ioutil"
	"os"

	jsonpatch "github.com/evanphx/json-patch/v5"
	"github.com/ghodss/yaml"
	jsondiff "github.com/mattbaird/jsonpatch"
)

type ByteStream = []byte

func load(file *os.File) (ByteStream, error) {
	bsYaml, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	bsJson, err := yaml.YAMLToJSON(bsYaml)
	if err != nil {
		return nil, err
	}
	return bsJson, nil
}

func loadFromPath(path string) (ByteStream, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return load(file)
}

func CreatePatchFromFiles(original string, target string) (ByteStream, error) {
	bsOriginal, err := loadFromPath(original)
	if err != nil {
		return nil, err
	}
	bsTarget, err := loadFromPath(target)
	if err != nil {
		return nil, err
	}
	return CreatePatchFromBytes(bsOriginal, bsTarget)
}

func CreatePatchFromBytes(bsOriginal ByteStream, bsTarget ByteStream) (ByteStream, error) {
	patch, err := jsondiff.CreatePatch(bsOriginal, bsTarget)
	if err != nil {
		return nil, err
	}
	bsPatch, err := yaml.Marshal(patch)
	if err != nil {
		return nil, err
	}
	return bsPatch, nil
}

func PatchInplace(original string, sources []string) (ByteStream, error) {
	file, err := os.OpenFile(original, os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}
	bsTarget, err := Patch(file, sources)
	if err != nil {
		return nil, err
	}
	file.Truncate(0)
	_, err = file.Write(bsTarget)
	return bsTarget, err
}

func Patch(file *os.File, sources []string) (ByteStream, error) {

	bsTarget, err := load(file)
	if err != nil {
		return nil, err
	}

	for _, source := range sources {
		bsSource, err := loadFromPath(source)
		if err != nil {
			return bsTarget, err
		}
		patch, err := jsonpatch.DecodePatch(bsSource)
		if err != nil {
			return bsTarget, err
		}
		bsTarget, err = patch.Apply(bsTarget)
		if err != nil {
			return bsTarget, err
		}
	}
	return yaml.JSONToYAML(bsTarget)
}
