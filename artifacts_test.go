package go4th

import (
	"testing"
)

func TestAddArtifact(t *testing.T) {
	alert := NewAlert()

	art, err := NewArtifact("file", "./testData/thehive.txt")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	alert.AddArtifact(art)

	if len(alert.Artifacts) != 1 {
		t.Errorf("expecting one artifact, but found %d", len(alert.Artifacts))
	}
}

func TestSetTLPArtifact(t *testing.T) {
	art, err := NewArtifact("file", "./testData/thehive.txt")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if art.TLP != Amber {
		t.Errorf("expected TLP to be %d, but found %d", Amber, art.TLP)
	}

	art.SetTLP(Red)

	if art.TLP != Red {
		t.Errorf("expected TLP to be %d, but found %d", Red, art.TLP)
	}

	err = art.SetTLP(10)
	if err != nil && err.Error() != "tlp provided is not a valid value" {
		t.Errorf("expected tlp provided is not a valid value as error, but none was found")
	}
}

func TestSetMessage(t *testing.T) {
	art, err := NewArtifact("file", "./testData/thehive.txt")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if art.Message != "" {
		t.Errorf("expected Message to be empty, but found %s", art.Message)
	}

	art.SetMessage("Testing")

	if art.Message != "Testing" {
		t.Errorf("expected Message to be Testing, but found %s", art.Message)
	}

	err = art.SetMessage("")

	if err != nil && err.Error() != "msg could not be empty" {
		t.Errorf("expected msg could not be empty as error, but none was found")
	}
}

func TestSetTagsArtifact(t *testing.T) {
	art, err := NewArtifact("file", "./testData/thehive.txt")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
	}

	if len(art.Tags) != 0 {
		t.Errorf("expected Tags to be empty, but found %d", len(art.Tags))
	}

	art.SetTags([]string{"one", "two"})

	if len(art.Tags) != 2 {
		t.Errorf("expected Tags to have two, but found %d", len(art.Tags))
	}

	if art.Tags[0] != "one" || art.Tags[1] != "two" {
		t.Errorf("expected Tags to be [one, two], but found %s", art.Tags)
	}

	err = art.SetTags([]string{})

	if err != nil && err.Error() != "tags could not be empty" {
		t.Errorf("expected tags could not be empty as error, but none was found")
	}
}
