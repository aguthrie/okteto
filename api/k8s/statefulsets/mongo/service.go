package mongo

import (
	"github.com/okteto/app/api/model"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

//TranslateService returns the service for mongo
func TranslateService(s *model.Space) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      model.MONGO,
			Namespace: s.Name,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{"app": model.MONGO},
			Type:     apiv1.ServiceTypeClusterIP,
			Ports: []apiv1.ServicePort{
				apiv1.ServicePort{
					Name:       "p27017",
					Port:       27017,
					TargetPort: intstr.IntOrString{StrVal: "27017"},
				},
			},
		},
	}
}
