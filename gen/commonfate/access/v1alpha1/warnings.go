package accessv1alpha1

// OK returns true if there are no warnings.
func (w *Warnings) OK() bool {
	return len(w.GrantsInvalidStatus) == 0 && len(w.GrantsPermissionDenied) == 0
}
