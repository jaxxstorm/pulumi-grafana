// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"errors"
	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource manages Grafana SLOs.
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/slo/)
// * [API documentation](https://grafana.com/docs/grafana-cloud/slo/api/)
// * [Additional Information On Alerting Rule Annotations and Labels](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/#templating/)
//
// ## Example Usage
type SLO struct {
	pulumi.CustomResourceState

	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings pulumix.GArrayOutput[SLOAlerting, SLOAlertingOutput] `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumix.Output[string] `pulumi:"description"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels pulumix.GArrayOutput[SLOLabel, SLOLabelOutput] `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumix.Output[string] `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives pulumix.GArrayOutput[SLOObjective, SLOObjectiveOutput] `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries pulumix.GArrayOutput[SLOQuery, SLOQueryOutput] `pulumi:"queries"`
}

// NewSLO registers a new resource with the given unique name, arguments, and options.
func NewSLO(ctx *pulumi.Context,
	name string, args *SLOArgs, opts ...pulumi.ResourceOption) (*SLO, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Objectives == nil {
		return nil, errors.New("invalid value for required argument 'Objectives'")
	}
	if args.Queries == nil {
		return nil, errors.New("invalid value for required argument 'Queries'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SLO
	err := ctx.RegisterResource("grafana:index/sLO:SLO", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSLO gets an existing SLO resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSLO(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SLOState, opts ...pulumi.ResourceOption) (*SLO, error) {
	var resource SLO
	err := ctx.ReadResource("grafana:index/sLO:SLO", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SLO resources.
type sloState struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings []SLOAlerting `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description *string `pulumi:"description"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels []SLOLabel `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name *string `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives []SLOObjective `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries []SLOQuery `pulumi:"queries"`
}

type SLOState struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings pulumix.Input[[]*SLOAlertingArgs]
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumix.Input[*string]
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels pulumix.Input[[]*SLOLabelArgs]
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumix.Input[*string]
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives pulumix.Input[[]*SLOObjectiveArgs]
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries pulumix.Input[[]*SLOQueryArgs]
}

func (SLOState) ElementType() reflect.Type {
	return reflect.TypeOf((*sloState)(nil)).Elem()
}

type sloArgs struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings []SLOAlerting `pulumi:"alertings"`
	// Description is a free-text field that can provide more context to an SLO.
	Description string `pulumi:"description"`
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels []SLOLabel `pulumi:"labels"`
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name *string `pulumi:"name"`
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives []SLOObjective `pulumi:"objectives"`
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries []SLOQuery `pulumi:"queries"`
}

// The set of arguments for constructing a SLO resource.
type SLOArgs struct {
	// Configures the alerting rules that will be generated for each
	// 			time window associated with the SLO. Grafana SLOs can generate
	// 			alerts when the short-term error budget burn is very high, the
	// 			long-term error budget burn rate is high, or when the remaining
	// 			error budget is below a certain threshold. Annotations and Labels support templating.
	Alertings pulumix.Input[[]*SLOAlertingArgs]
	// Description is a free-text field that can provide more context to an SLO.
	Description pulumix.Input[string]
	// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
	Labels pulumix.Input[[]*SLOLabelArgs]
	// Name should be a short description of your indicator. Consider names like "API Availability"
	Name pulumix.Input[*string]
	// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
	Objectives pulumix.Input[[]*SLOObjectiveArgs]
	// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
	Queries pulumix.Input[[]*SLOQueryArgs]
}

func (SLOArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sloArgs)(nil)).Elem()
}

type SLOOutput struct{ *pulumi.OutputState }

func (SLOOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SLO)(nil)).Elem()
}

func (o SLOOutput) ToSLOOutput() SLOOutput {
	return o
}

func (o SLOOutput) ToSLOOutputWithContext(ctx context.Context) SLOOutput {
	return o
}

func (o SLOOutput) ToOutput(ctx context.Context) pulumix.Output[SLO] {
	return pulumix.Output[SLO]{
		OutputState: o.OutputState,
	}
}

// Configures the alerting rules that will be generated for each
//
//	time window associated with the SLO. Grafana SLOs can generate
//	alerts when the short-term error budget burn is very high, the
//	long-term error budget burn rate is high, or when the remaining
//	error budget is below a certain threshold. Annotations and Labels support templating.
func (o SLOOutput) Alertings() pulumix.GArrayOutput[SLOAlerting, SLOAlertingOutput] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.GArrayOutput[SLOAlerting, SLOAlertingOutput] { return v.Alertings })
	unwrapped := pulumix.Flatten[[]SLOAlerting, pulumix.GArrayOutput[SLOAlerting, SLOAlertingOutput]](value)
	return pulumix.GArrayOutput[SLOAlerting, SLOAlertingOutput]{OutputState: unwrapped.OutputState}
}

// Description is a free-text field that can provide more context to an SLO.
func (o SLOOutput) Description() pulumix.Output[string] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.Output[string] { return v.Description })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Additional labels that will be attached to all metrics generated from the query. These labels are useful for grouping SLOs in dashboard views that you create by hand. Labels must adhere to Prometheus label name schema - "^[a-zA-Z*][a-zA-Z0-9*]*$"
func (o SLOOutput) Labels() pulumix.GArrayOutput[SLOLabel, SLOLabelOutput] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.GArrayOutput[SLOLabel, SLOLabelOutput] { return v.Labels })
	unwrapped := pulumix.Flatten[[]SLOLabel, pulumix.GArrayOutput[SLOLabel, SLOLabelOutput]](value)
	return pulumix.GArrayOutput[SLOLabel, SLOLabelOutput]{OutputState: unwrapped.OutputState}
}

// Name should be a short description of your indicator. Consider names like "API Availability"
func (o SLOOutput) Name() pulumix.Output[string] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.Output[string] { return v.Name })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Over each rolling time window, the remaining error budget will be calculated, and separate alerts can be generated for each time window based on the SLO burn rate or remaining error budget.
func (o SLOOutput) Objectives() pulumix.GArrayOutput[SLOObjective, SLOObjectiveOutput] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.GArrayOutput[SLOObjective, SLOObjectiveOutput] { return v.Objectives })
	unwrapped := pulumix.Flatten[[]SLOObjective, pulumix.GArrayOutput[SLOObjective, SLOObjectiveOutput]](value)
	return pulumix.GArrayOutput[SLOObjective, SLOObjectiveOutput]{OutputState: unwrapped.OutputState}
}

// Query describes the indicator that will be measured against the objective. Freeform Query types are currently supported.
func (o SLOOutput) Queries() pulumix.GArrayOutput[SLOQuery, SLOQueryOutput] {
	value := pulumix.Apply[SLO](o, func(v SLO) pulumix.GArrayOutput[SLOQuery, SLOQueryOutput] { return v.Queries })
	unwrapped := pulumix.Flatten[[]SLOQuery, pulumix.GArrayOutput[SLOQuery, SLOQueryOutput]](value)
	return pulumix.GArrayOutput[SLOQuery, SLOQueryOutput]{OutputState: unwrapped.OutputState}
}

func init() {
	pulumi.RegisterOutputType(SLOOutput{})
}
