package v2alpha1

func (h *HorusecPlatform) GetDefaultLabel() map[string]string {
	return map[string]string{"app.kubernetes.io/managed-by": "horusec"}
}