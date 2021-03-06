// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package authorization

import (
	reflect "reflect"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	sets "k8s.io/apimachinery/pkg/util/sets"
	api "k8s.io/kubernetes/pkg/api"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationAction, InType: reflect.TypeOf(&Action{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterPolicy, InType: reflect.TypeOf(&ClusterPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterPolicyBinding, InType: reflect.TypeOf(&ClusterPolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterPolicyBindingList, InType: reflect.TypeOf(&ClusterPolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterPolicyList, InType: reflect.TypeOf(&ClusterPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterRole, InType: reflect.TypeOf(&ClusterRole{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterRoleBinding, InType: reflect.TypeOf(&ClusterRoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterRoleBindingList, InType: reflect.TypeOf(&ClusterRoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationClusterRoleList, InType: reflect.TypeOf(&ClusterRoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationGroupRestriction, InType: reflect.TypeOf(&GroupRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationIsPersonalSubjectAccessReview, InType: reflect.TypeOf(&IsPersonalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationLocalResourceAccessReview, InType: reflect.TypeOf(&LocalResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationLocalSubjectAccessReview, InType: reflect.TypeOf(&LocalSubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicy, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicyBinding, InType: reflect.TypeOf(&PolicyBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicyBindingList, InType: reflect.TypeOf(&PolicyBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicyList, InType: reflect.TypeOf(&PolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicyRule, InType: reflect.TypeOf(&PolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationPolicyRuleBuilder, InType: reflect.TypeOf(&PolicyRuleBuilder{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationResourceAccessReview, InType: reflect.TypeOf(&ResourceAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationResourceAccessReviewResponse, InType: reflect.TypeOf(&ResourceAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRole, InType: reflect.TypeOf(&Role{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleBinding, InType: reflect.TypeOf(&RoleBinding{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleBindingList, InType: reflect.TypeOf(&RoleBindingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleBindingRestriction, InType: reflect.TypeOf(&RoleBindingRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleBindingRestrictionList, InType: reflect.TypeOf(&RoleBindingRestrictionList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleBindingRestrictionSpec, InType: reflect.TypeOf(&RoleBindingRestrictionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationRoleList, InType: reflect.TypeOf(&RoleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSelfSubjectRulesReview, InType: reflect.TypeOf(&SelfSubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSelfSubjectRulesReviewSpec, InType: reflect.TypeOf(&SelfSubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationServiceAccountReference, InType: reflect.TypeOf(&ServiceAccountReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationServiceAccountRestriction, InType: reflect.TypeOf(&ServiceAccountRestriction{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSubjectAccessReview, InType: reflect.TypeOf(&SubjectAccessReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSubjectAccessReviewResponse, InType: reflect.TypeOf(&SubjectAccessReviewResponse{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSubjectRulesReview, InType: reflect.TypeOf(&SubjectRulesReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSubjectRulesReviewSpec, InType: reflect.TypeOf(&SubjectRulesReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationSubjectRulesReviewStatus, InType: reflect.TypeOf(&SubjectRulesReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopyauthorizationUserRestriction, InType: reflect.TypeOf(&UserRestriction{})},
	)
}

// DeepCopyauthorizationAction is an autogenerated deepcopy function.
func DeepCopyauthorizationAction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Action)
		out := out.(*Action)
		*out = *in
		// in.Content is kind 'Interface'
		if in.Content != nil {
			newVal, err := c.DeepCopy(&in.Content)
			if err != nil {
				return err
			}
			out.Content = *newVal.(*runtime.Object)
		}
		return nil
	}
}

// DeepCopyauthorizationClusterPolicy is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicy)
		out := out.(*ClusterPolicy)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(ClusterRolesByName)
			for key, val := range *in {
				newVal, err := c.DeepCopy(&val)
				if err != nil {
					return err
				}
				(*out)[key] = *newVal.(**ClusterRole)
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterPolicyBinding is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterPolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBinding)
		out := out.(*ClusterPolicyBinding)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		out.LastModified = in.LastModified.DeepCopy()
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(ClusterRoleBindingsByName)
			for key, val := range *in {
				newVal, err := c.DeepCopy(&val)
				if err != nil {
					return err
				}
				(*out)[key] = *newVal.(**ClusterRoleBinding)
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterPolicyBindingList is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterPolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyBindingList)
		out := out.(*ClusterPolicyBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationClusterPolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterPolicyList is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterPolicyList)
		out := out.(*ClusterPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterPolicy, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationClusterPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterRole is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRole)
		out := out.(*ClusterRole)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationPolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterRoleBinding is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBinding)
		out := out.(*ClusterRoleBinding)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]api.ObjectReference, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationClusterRoleBindingList is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterRoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleBindingList)
		out := out.(*ClusterRoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationClusterRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationClusterRoleList is an autogenerated deepcopy function.
func DeepCopyauthorizationClusterRoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterRoleList)
		out := out.(*ClusterRoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterRole, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationClusterRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationGroupRestriction is an autogenerated deepcopy function.
func DeepCopyauthorizationGroupRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupRestriction)
		out := out.(*GroupRestriction)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]v1.LabelSelector, len(*in))
			for i := range *in {
				newVal, err := c.DeepCopy(&(*in)[i])
				if err != nil {
					return err
				}
				(*out)[i] = *newVal.(*v1.LabelSelector)
			}
		}
		return nil
	}
}

// DeepCopyauthorizationIsPersonalSubjectAccessReview is an autogenerated deepcopy function.
func DeepCopyauthorizationIsPersonalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*IsPersonalSubjectAccessReview)
		out := out.(*IsPersonalSubjectAccessReview)
		*out = *in
		return nil
	}
}

// DeepCopyauthorizationLocalResourceAccessReview is an autogenerated deepcopy function.
func DeepCopyauthorizationLocalResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalResourceAccessReview)
		out := out.(*LocalResourceAccessReview)
		*out = *in
		return DeepCopyauthorizationAction(&in.Action, &out.Action, c)
	}
}

// DeepCopyauthorizationLocalSubjectAccessReview is an autogenerated deepcopy function.
func DeepCopyauthorizationLocalSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*LocalSubjectAccessReview)
		out := out.(*LocalSubjectAccessReview)
		*out = *in
		if err := DeepCopyauthorizationAction(&in.Action, &out.Action, c); err != nil {
			return err
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationPolicy is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Policy)
		out := out.(*Policy)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		out.LastModified = in.LastModified.DeepCopy()
		if in.Roles != nil {
			in, out := &in.Roles, &out.Roles
			*out = make(RolesByName)
			for key, val := range *in {
				newVal, err := c.DeepCopy(&val)
				if err != nil {
					return err
				}
				(*out)[key] = *newVal.(**Role)
			}
		}
		return nil
	}
}

// DeepCopyauthorizationPolicyBinding is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicyBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBinding)
		out := out.(*PolicyBinding)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		out.LastModified = in.LastModified.DeepCopy()
		if in.RoleBindings != nil {
			in, out := &in.RoleBindings, &out.RoleBindings
			*out = make(RoleBindingsByName)
			for key, val := range *in {
				newVal, err := c.DeepCopy(&val)
				if err != nil {
					return err
				}
				(*out)[key] = *newVal.(**RoleBinding)
			}
		}
		return nil
	}
}

// DeepCopyauthorizationPolicyBindingList is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicyBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyBindingList)
		out := out.(*PolicyBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PolicyBinding, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationPolicyBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationPolicyList is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyList)
		out := out.(*PolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Policy, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationPolicyRule is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRule)
		out := out.(*PolicyRule)
		*out = *in
		if in.Verbs != nil {
			in, out := &in.Verbs, &out.Verbs
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		// in.AttributeRestrictions is kind 'Interface'
		if in.AttributeRestrictions != nil {
			newVal, err := c.DeepCopy(&in.AttributeRestrictions)
			if err != nil {
				return err
			}
			out.AttributeRestrictions = *newVal.(*runtime.Object)
		}
		if in.APIGroups != nil {
			in, out := &in.APIGroups, &out.APIGroups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.ResourceNames != nil {
			in, out := &in.ResourceNames, &out.ResourceNames
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.NonResourceURLs != nil {
			in, out := &in.NonResourceURLs, &out.NonResourceURLs
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopyauthorizationPolicyRuleBuilder is an autogenerated deepcopy function.
func DeepCopyauthorizationPolicyRuleBuilder(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PolicyRuleBuilder)
		out := out.(*PolicyRuleBuilder)
		*out = *in
		return DeepCopyauthorizationPolicyRule(&in.PolicyRule, &out.PolicyRule, c)
	}
}

// DeepCopyauthorizationResourceAccessReview is an autogenerated deepcopy function.
func DeepCopyauthorizationResourceAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReview)
		out := out.(*ResourceAccessReview)
		*out = *in
		return DeepCopyauthorizationAction(&in.Action, &out.Action, c)
	}
}

// DeepCopyauthorizationResourceAccessReviewResponse is an autogenerated deepcopy function.
func DeepCopyauthorizationResourceAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ResourceAccessReviewResponse)
		out := out.(*ResourceAccessReviewResponse)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopyauthorizationRole is an autogenerated deepcopy function.
func DeepCopyauthorizationRole(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Role)
		out := out.(*Role)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationPolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationRoleBinding is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleBinding(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBinding)
		out := out.(*RoleBinding)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		if in.Subjects != nil {
			in, out := &in.Subjects, &out.Subjects
			*out = make([]api.ObjectReference, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationRoleBindingList is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleBindingList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingList)
		out := out.(*RoleBindingList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBinding, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationRoleBinding(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationRoleBindingRestriction is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleBindingRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestriction)
		out := out.(*RoleBindingRestriction)
		*out = *in
		newVal, err := c.DeepCopy(&in.ObjectMeta)
		if err != nil {
			return err
		}
		out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		return DeepCopyauthorizationRoleBindingRestrictionSpec(&in.Spec, &out.Spec, c)
	}
}

// DeepCopyauthorizationRoleBindingRestrictionList is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleBindingRestrictionList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionList)
		out := out.(*RoleBindingRestrictionList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]RoleBindingRestriction, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationRoleBindingRestriction(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationRoleBindingRestrictionSpec is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleBindingRestrictionSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleBindingRestrictionSpec)
		out := out.(*RoleBindingRestrictionSpec)
		*out = *in
		if in.UserRestriction != nil {
			in, out := &in.UserRestriction, &out.UserRestriction
			*out = new(UserRestriction)
			if err := DeepCopyauthorizationUserRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		if in.GroupRestriction != nil {
			in, out := &in.GroupRestriction, &out.GroupRestriction
			*out = new(GroupRestriction)
			if err := DeepCopyauthorizationGroupRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		if in.ServiceAccountRestriction != nil {
			in, out := &in.ServiceAccountRestriction, &out.ServiceAccountRestriction
			*out = new(ServiceAccountRestriction)
			if err := DeepCopyauthorizationServiceAccountRestriction(*in, *out, c); err != nil {
				return err
			}
		}
		return nil
	}
}

// DeepCopyauthorizationRoleList is an autogenerated deepcopy function.
func DeepCopyauthorizationRoleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoleList)
		out := out.(*RoleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Role, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationRole(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationSelfSubjectRulesReview is an autogenerated deepcopy function.
func DeepCopyauthorizationSelfSubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReview)
		out := out.(*SelfSubjectRulesReview)
		*out = *in
		if err := DeepCopyauthorizationSelfSubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return DeepCopyauthorizationSubjectRulesReviewStatus(&in.Status, &out.Status, c)
	}
}

// DeepCopyauthorizationSelfSubjectRulesReviewSpec is an autogenerated deepcopy function.
func DeepCopyauthorizationSelfSubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SelfSubjectRulesReviewSpec)
		out := out.(*SelfSubjectRulesReviewSpec)
		*out = *in
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationServiceAccountReference is an autogenerated deepcopy function.
func DeepCopyauthorizationServiceAccountReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountReference)
		out := out.(*ServiceAccountReference)
		*out = *in
		return nil
	}
}

// DeepCopyauthorizationServiceAccountRestriction is an autogenerated deepcopy function.
func DeepCopyauthorizationServiceAccountRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ServiceAccountRestriction)
		out := out.(*ServiceAccountRestriction)
		*out = *in
		if in.ServiceAccounts != nil {
			in, out := &in.ServiceAccounts, &out.ServiceAccounts
			*out = make([]ServiceAccountReference, len(*in))
			copy(*out, *in)
		}
		if in.Namespaces != nil {
			in, out := &in.Namespaces, &out.Namespaces
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationSubjectAccessReview is an autogenerated deepcopy function.
func DeepCopyauthorizationSubjectAccessReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReview)
		out := out.(*SubjectAccessReview)
		*out = *in
		if err := DeepCopyauthorizationAction(&in.Action, &out.Action, c); err != nil {
			return err
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make(sets.String)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationSubjectAccessReviewResponse is an autogenerated deepcopy function.
func DeepCopyauthorizationSubjectAccessReviewResponse(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectAccessReviewResponse)
		out := out.(*SubjectAccessReviewResponse)
		*out = *in
		return nil
	}
}

// DeepCopyauthorizationSubjectRulesReview is an autogenerated deepcopy function.
func DeepCopyauthorizationSubjectRulesReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReview)
		out := out.(*SubjectRulesReview)
		*out = *in
		if err := DeepCopyauthorizationSubjectRulesReviewSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return DeepCopyauthorizationSubjectRulesReviewStatus(&in.Status, &out.Status, c)
	}
}

// DeepCopyauthorizationSubjectRulesReviewSpec is an autogenerated deepcopy function.
func DeepCopyauthorizationSubjectRulesReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewSpec)
		out := out.(*SubjectRulesReviewSpec)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Scopes != nil {
			in, out := &in.Scopes, &out.Scopes
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopyauthorizationSubjectRulesReviewStatus is an autogenerated deepcopy function.
func DeepCopyauthorizationSubjectRulesReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*SubjectRulesReviewStatus)
		out := out.(*SubjectRulesReviewStatus)
		*out = *in
		if in.Rules != nil {
			in, out := &in.Rules, &out.Rules
			*out = make([]PolicyRule, len(*in))
			for i := range *in {
				if err := DeepCopyauthorizationPolicyRule(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopyauthorizationUserRestriction is an autogenerated deepcopy function.
func DeepCopyauthorizationUserRestriction(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserRestriction)
		out := out.(*UserRestriction)
		*out = *in
		if in.Users != nil {
			in, out := &in.Users, &out.Users
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Selectors != nil {
			in, out := &in.Selectors, &out.Selectors
			*out = make([]v1.LabelSelector, len(*in))
			for i := range *in {
				newVal, err := c.DeepCopy(&(*in)[i])
				if err != nil {
					return err
				}
				(*out)[i] = *newVal.(*v1.LabelSelector)
			}
		}
		return nil
	}
}
