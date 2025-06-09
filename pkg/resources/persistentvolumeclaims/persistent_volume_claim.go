package persistentvolumeclaims

import (
	"github.com/arkmq-org/activemq-artemis-operator/api/v1beta1"
	"github.com/arkmq-org/activemq-artemis-operator/pkg/utils/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PersistentVolumeClaim(namespace string, existing *corev1.PersistentVolumeClaim, templateFromCr *v1beta1.VolumeClaimTemplate) *corev1.PersistentVolumeClaim {

	var desired *corev1.PersistentVolumeClaim = existing
	if desired == nil {
		desired = &corev1.PersistentVolumeClaim{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "v1",
				Kind:       "PersistentVolumeClaim",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      templateFromCr.Name,
				Namespace: namespace,
			},
			Spec: corev1.PersistentVolumeClaimSpec{},
		}
	}
	//apply desired
	desired.ObjectMeta.Labels = templateFromCr.Labels
	common.ApplyAnnotations(&desired.ObjectMeta, templateFromCr.Annotations)

	// apply defaults
	defaultVolumeMode := corev1.PersistentVolumeFilesystem
	if templateFromCr.Spec.VolumeMode == nil {
		templateFromCr.Spec.VolumeMode = &defaultVolumeMode
	}

	desired.Spec = templateFromCr.Spec

	return desired
}
