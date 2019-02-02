package go4th

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/gabriel-vasile/mimetype"
)

// Artifact defines an Alert/Case artifact
type Artifact struct {
	DataType string   `json:"dataType,omitempty"`
	Data     string   `json:"data,omitempty"`
	Message  string   `json:"message,omitempty"`
	TLP      TLP      `json:"tlp,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

// NewArtifact returns a new artifact
func NewArtifact(dataType, data string) (*Artifact, error) {
	var art *Artifact
	art = new(Artifact)
	art.DataType = dataType
	art.TLP = Amber
	art.Tags = []string{}

	if data == "" {
		return nil, fmt.Errorf("missing artifact data")
	}
	if dataType != "" && dataType == "file" {
		if _, err := os.Stat(data); !os.IsNotExist(err) {
			dat, err := ioutil.ReadFile(data)
			if err != nil {
				return nil, fmt.Errorf("unable to open/read artifact %s", data)
			}
			mime, _ := mimetype.Detect(dat)
			filename := path.Base(data)
			encodedData := base64.StdEncoding.EncodeToString(dat)

			art.DataType = dataType
			art.Data = fmt.Sprintf("%s;%s;%s", filename, mime, encodedData)
		} else {
			return nil, fmt.Errorf("unable to read artifact %s", data)
		}
	} else {
		art.Data = data
	}
	return art, nil
}

/*
	Artifact methods
*/

// SetTLP sets the TLP for the asset
func (a *Artifact) SetTLP(tlp TLP) error {
	if tlp != White && tlp != Green && tlp != Amber && tlp != Red {
		return fmt.Errorf("tlp provided is not a valid value")
	}
	a.TLP = tlp
	return nil
}

// SetMessage sets the message for the artifact
func (a *Artifact) SetMessage(msg string) error {
	if msg == "" {
		return fmt.Errorf("msg could not be empty")
	}
	a.Message = msg
	return nil
}

// SetTags sets the tags for the artifact
func (a *Artifact) SetTags(tags []string) error {
	if len(tags) == 0 {
		return fmt.Errorf("tags could not be empty")
	}
	a.Tags = tags
	return nil
}

/*
	API Calls
*/
