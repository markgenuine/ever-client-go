package goton

import (
	"testing"
)

func TestClientMethods(t *testing.T) {
	config := NewConfig(0)
	client, err := InitClient(config)
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

	idVersion := client.Version()
	idAPIReference := client.GetAPIReference()
	idBuildInfo := client.GetBuildInfo()

	valueVersion, err := client.GetResp(idVersion)
	if err != nil {
		t.Errorf("test Failed - Error get version, err: %s", err)
	}

	if valueVersion.(ResultOfVersion).Version != VersionLibSDK {
		t.Errorf("Version lib %s different version go-ton-sdk %s", valueVersion.(ResultOfVersion).Version, VersionLibSDK)
	}

	valueAPIReference, err := client.GetResp(idAPIReference)
	if err != nil {
		t.Errorf("test Failed - Error get api reference, err: %s", err)
	}

	if valueAPIReference.(ResultOfGetAPIReference).API.Version != VersionLibSDK {
		t.Errorf("API Version %s different version go-ton-sdk %s", valueAPIReference.(ResultOfGetAPIReference).API.Version, VersionLibSDK)
	}

	_, err = client.GetResp(idBuildInfo)
	if err != nil {
		t.Errorf("test Failed - Error get build info, err: %s", err)
	}
}
