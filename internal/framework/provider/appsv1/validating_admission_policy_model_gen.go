package appsv1

import (
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ValidatingAdmissionPolicyModel struct {
	Timeouts timeouts.Value `tfsdk:"timeouts"`

	ID types.String `tfsdk:"id" manifest:""`

	Metadata struct {
		Annotations     map[string]types.String `tfsdk:"annotations" manifest:"annotations"`
		GenerateName    types.String            `tfsdk:"generate_name" manifest:"generateName"`
		Generation      types.Int64             `tfsdk:"generation" manifest:"generation"`
		Labels          map[string]types.String `tfsdk:"labels" manifest:"labels"`
		Name            types.String            `tfsdk:"name" manifest:"name"`
		Namespace       types.String            `tfsdk:"namespace" manifest:"namespace"`
		ResourceVersion types.String            `tfsdk:"resource_version" manifest:"resourceVersion"`
		UID             types.String            `tfsdk:"uid" manifest:"uid"`
	} `tfsdk:"metadata" manifest:"metadata"`

	Spec struct {
		AuditAnnotations []struct {
			Key              types.String `tfsdk:"key" manifest:"key"`
			ValueExpressions types.String `tfsdk:"value_expressions" manifest:"valueExpressions"`
		} `tfsdk:"audit_annotations" manifest:"auditAnnotations"`
		FailurePolicy   types.String `tfsdk:"failure_policy" manifest:"failurePolicy"`
		MatchConditions []struct {
			Expression types.String `tfsdk:"expression" manifest:"expression"`
			Name       types.String `tfsdk:"Name" manifest:"name"`
		} `tfsdk:"match_conditions" manifest:"matchConditions"`
		MatchConstraints struct {
			ExcludeResourceRules []struct {
				APIGroups     []types.String `tfsdk:"api_groups" manifest:"apiGroups"`
				APIVersions   []types.String `tfsdk:"api_versions" manifest:"apiVersions"`
				Operations    []types.String `tfsdk:"operations" manifest:"operations"`
				ResourceNames []types.String `tfsdk:"resource_names" manifest:"resourceNames"`
				Resources     []types.String `tfsdk:"resources" manifest:"resources"`
				Scope         types.String   `tfsdk:"scope" manifest:"scope"`
			} `tfsdk:"exclude_resource_rules" manifest:"excludeResourceRules"`
			MatchPolicy       types.String `tfsdk:"match_policy" manifest:"matchPolicy"`
			NamespaceSelector struct {
				APIGroups     []types.String `tfsdk:"api_groups" manifest:"apiGroups"`
				APIVersions   []types.String `tfsdk:"api_versions" manifest:"apiVersions"`
				Operations    []types.String `tfsdk:"operations" manifest:"operations"`
				ResourceNames []types.String `tfsdk:"resource_names" manifest:"resourceNames"`
				Resources     []types.String `tfsdk:"resources" manifest:"resources"`
				Scope         types.String   `tfsdk:"scope" manifest:"scope"`
			} `tfsdk:"namespace_selector" manifest:"NamespaceSelector"`
			ObjectSelector struct {
				LabelSelector struct {
					MatchExpressions []struct {
						Key      types.String   `tfsdk:"key" manifest:"key"`
						Operator types.String   `tfsdk:"operator" manifest:"operator"`
						Values   []types.String `tfsdk:"values" manifest:"values"`
					} `tfsdk:"match_expressions" manifest:"matchExpressions"`
					MatchLabels map[string]types.String `tfsdk:"match_labels" manifest:"matchLabels"`
				} `tfsdk:"label_selector" manifest:"labelSelector"`
			} `tfsdk:"object_selector" manifest:"objectSelector"`
			ResourceRules []struct {
				APIGroups     []types.String `tfsdk:"api_groups" manifest:"apiGroups"`
				APIVersions   []types.String `tfsdk:"api_versions" manifest:"apiVersions"`
				Operations    []types.String `tfsdk:"operations" manifest:"operations"`
				ResourceNames []types.String `tfsdk:"resource_names" manifest:"resourceNames"`
				Resources     []types.String `tfsdk:"resources" manifest:"resources"`
				Scope         types.String   `tfsdk:"scope" manifest:"scope"`
			} `tfsdk:"resource_rules" manifest:"resourceRules"`
		} `tfsdk:"match_constraints" manifest:"matchConstraints"`
		ParamKind struct {
			APIVersion types.String `tfsdk:"api_version" manifest:"apiVersion"`
			Kind       types.String `tfsdk:"kind" manifest:"kind"`
		} `tfsdk:"param_kind" manifest:"paramKind"`
		Validations []struct {
			Expression        types.String `tfsdk:"expression" manifest:"expression"`
			Message           types.String `tfsdk:"message" manifest:"message"`
			MessageExpression types.String `tfsdk:"message_expression" manifest:"messageExpression"`
			Reason            types.String `tfsdk:"reason" manifest:"reason"`
		}
		Variable []struct {
			Expression types.String `tfsdk:"expression" manifest:"expression"`
			Name       types.String `tfsdk:"name" manifest:"name"`
		}
	}
}