package yandex

import (
	"fmt"
	"regexp"

	"k8s.io/apimachinery/pkg/types"
)

var (
	regExpProviderID = regexp.MustCompile(`^` + providerName + `://([^/]+)/([^/]+)/([^/]+)$`)
)

// MapNodeNameToInstanceName maps a k8s Node Name to a Yandex.Cloud Instance Name
// Currently - this is a simple string cast.
func MapNodeNameToInstanceName(nodeName types.NodeName) string {
	return string(nodeName)
}

// ParseProviderID splits a providerID into Folder ID, Zone and Instance Name.
func ParseProviderID(providerID string) (folderID string, zone string, instanceName string, err error) {
	// providerID is in the following form "${providerName}://${folderID}/${zone}/${instanceName}"
	// So for input "yandex://b1g4c2a3g6vkffp3qacq/ru-central1-a/e2e-test-node0" output will be  "b1g4c2a3g6vkffp3qacq", "ru-central1-a", "e2e-test-node0".
	matches := regExpProviderID.FindStringSubmatch(providerID)
	if len(matches) != 4 {
		return "", "", "", fmt.Errorf("unexpected input: %s", providerID)
	}

	return matches[1], matches[2], matches[3], nil
}
