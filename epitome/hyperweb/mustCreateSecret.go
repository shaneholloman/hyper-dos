package hyperweb

import (
	"context"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func mustCreateOperatorOAuthSecret(
	clientset kubernetes.Clientset,
	namespace string,
	name string,
	clientId string,
	clientSecret string,
) (err error) {

	// create secret
	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "operator-oauth",
			Namespace: hyperwebNamespace,
		},
		Data: map[string][]byte{
			"clientID":     []byte(clientId),
			"clientSecret": []byte(clientSecret),
		},
	}

	_, err = clientset.CoreV1().Secrets(hyperwebNamespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		logrus.Fatalf("failed to create secret: %v", err)
	}

	logrus.Infof("created operator-oauth secret")

	// TODO delete hyperbolic token?
	return
}
