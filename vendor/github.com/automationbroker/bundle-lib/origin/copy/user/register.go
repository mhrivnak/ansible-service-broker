package user

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// GroupName -
	GroupName = "user.openshift.io"
	// LegacyGroupName -
	LegacyGroupName = ""
)

// SchemeGroupVersion is group version used to register these objects
var (
	// SchemeGroupVersion -
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	// LegacySchemeGroupVersion -
	LegacySchemeGroupVersion = schema.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}

	// LegacySchemeBuilder -
	LegacySchemeBuilder = runtime.NewSchemeBuilder(addLegacyKnownTypes)
	// AddToShemeInCoreGroup -
	AddToSchemeInCoreGroup = LegacySchemeBuilder.AddToScheme

	// SchemeBuilder -
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme -
	AddToScheme = SchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// LegacyKind -
func LegacyKind(kind string) schema.GroupKind {
	return LegacySchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// LegacyResource takes an unqualified resource and returns back a Group qualified GroupResource
func LegacyResource(resource string) schema.GroupResource {
	return LegacySchemeGroupVersion.WithResource(resource).GroupResource()
}

// IsKindOrLegacy checks if the provided GroupKind matches with the given kind by looking
// up the API group and also the legacy API.
func IsKindOrLegacy(kind string, gk schema.GroupKind) bool {
	return gk == Kind(kind) || gk == LegacyKind(kind)
}

// IsResourceOrLegacy checks if the provided GroupResources matches with the given
// resource by looking up the API group and also the legacy API.
func IsResourceOrLegacy(resource string, gr schema.GroupResource) bool {
	return gr == Resource(resource) || gr == LegacyResource(resource)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&User{},
		&List{},
		&Identity{},
		&IdentityList{},
		&IdentityMapping{},
		&Group{},
		&GroupList{},
	)
	return nil
}

// Adds the list of known types to api.Scheme.
func addLegacyKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&User{},
		&List{},
		&Identity{},
		&IdentityList{},
		&IdentityMapping{},
		&Group{},
		&GroupList{},
	}
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	return nil
}