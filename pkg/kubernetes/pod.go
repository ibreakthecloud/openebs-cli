package kubernetes

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// Uncomment to load all auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// Constants are defined here
const (
	// OpenEBSNamespace Default OpenEBS namespace
	OpenEBSNamespace string = "openebs"

	// MayaApiServerLabelSelector is mapi
	MayaAPIServerLabelSelector string = "name=maya-apiserver"

	OpenEBSVersionLabelKey string = "openebs.io/version"
)

//PrintOpenEBSVersion returns pods
func PrintOpenEBSVersion() {
	pods, err := ClientSet.CoreV1().Pods(OpenEBSNamespace).List(metav1.ListOptions{
		LabelSelector: MayaAPIServerLabelSelector,
	})
	if err != nil {
		panic(err.Error())
	}

	for _, c := range pods.Items {
		if c.GetName() != "" {

			labels := c.GetLabels()
			openebsVersion := labels[OpenEBSVersionLabelKey]
			if openebsVersion == "" {
				fmt.Printf("\nOpenEBS is not yet installed :(")
				return
			}
			fmt.Printf("\nOpenEBS Version: %s", openebsVersion)
			return
		}
	}
	fmt.Printf("\nOpenEBS is not yet installed :(")
	return
}
