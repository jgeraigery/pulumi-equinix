// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package equinix

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Metro string

const (
	MetroAmsterdam     = Metro("AM")
	MetroAshburn       = Metro("DC")
	MetroAtlanta       = Metro("AT")
	MetroBarcelona     = Metro("BA")
	MetroBogota        = Metro("BG")
	MetroBordeaux      = Metro("BX")
	MetroBoston        = Metro("BO")
	MetroBrussels      = Metro("BL")
	MetroCalgary       = Metro("CL")
	MetroCanberra      = Metro("CA")
	MetroChicago       = Metro("CH")
	MetroDallas        = Metro("DA")
	MetroDenver        = Metro("DE")
	MetroDubai         = Metro("DX")
	MetroDublin        = Metro("DB")
	MetroFrankfurt     = Metro("FR")
	MetroGeneva        = Metro("GV")
	MetroHamburg       = Metro("HH")
	MetroHelsinki      = Metro("HE")
	MetroHongKong      = Metro("HK")
	MetroIstanbul      = Metro("IL")
	MetroKamloops      = Metro("KA")
	MetroLisbon        = Metro("LS")
	MetroLondon        = Metro("LD")
	MetroLosAngeles    = Metro("LA")
	MetroMadrid        = Metro("MD")
	MetroManchester    = Metro("MA")
	MetroMelbourne     = Metro("ME")
	MetroMexicoCity    = Metro("MX")
	MetroMiami         = Metro("MI")
	MetroMilan         = Metro("ML")
	MetroMontreal      = Metro("MT")
	MetroMumbai        = Metro("MB")
	MetroMunich        = Metro("MU")
	MetroNewYork       = Metro("NY")
	MetroOsaka         = Metro("OS")
	MetroParis         = Metro("PA")
	MetroPerth         = Metro("PE")
	MetroPhiladelphia  = Metro("PH")
	MetroRioDeJaneiro  = Metro("RJ")
	MetroSaoPaulo      = Metro("SP")
	MetroSeattle       = Metro("SE")
	MetroSeoul         = Metro("SL")
	MetroSiliconValley = Metro("SV")
	MetroSingapore     = Metro("SG")
	MetroSofia         = Metro("SO")
	MetroStockholm     = Metro("SK")
	MetroSydney        = Metro("SY")
	MetroTokyo         = Metro("TY")
	MetroToronto       = Metro("TR")
	MetroVancouver     = Metro("VA")
	MetroWarsaw        = Metro("WA")
	MetroWinnipeg      = Metro("WI")
	MetroZurich        = Metro("ZH")
)

func (Metro) ElementType() reflect.Type {
	return reflect.TypeOf((*Metro)(nil)).Elem()
}

func (e Metro) ToMetroOutput() MetroOutput {
	return pulumi.ToOutput(e).(MetroOutput)
}

func (e Metro) ToMetroOutputWithContext(ctx context.Context) MetroOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MetroOutput)
}

func (e Metro) ToMetroPtrOutput() MetroPtrOutput {
	return e.ToMetroPtrOutputWithContext(context.Background())
}

func (e Metro) ToMetroPtrOutputWithContext(ctx context.Context) MetroPtrOutput {
	return Metro(e).ToMetroOutputWithContext(ctx).ToMetroPtrOutputWithContext(ctx)
}

func (e Metro) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Metro) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Metro) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Metro) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MetroOutput struct{ *pulumi.OutputState }

func (MetroOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Metro)(nil)).Elem()
}

func (o MetroOutput) ToMetroOutput() MetroOutput {
	return o
}

func (o MetroOutput) ToMetroOutputWithContext(ctx context.Context) MetroOutput {
	return o
}

func (o MetroOutput) ToMetroPtrOutput() MetroPtrOutput {
	return o.ToMetroPtrOutputWithContext(context.Background())
}

func (o MetroOutput) ToMetroPtrOutputWithContext(ctx context.Context) MetroPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Metro) *Metro {
		return &v
	}).(MetroPtrOutput)
}

func (o MetroOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MetroOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Metro) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MetroOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetroOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Metro) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MetroPtrOutput struct{ *pulumi.OutputState }

func (MetroPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metro)(nil)).Elem()
}

func (o MetroPtrOutput) ToMetroPtrOutput() MetroPtrOutput {
	return o
}

func (o MetroPtrOutput) ToMetroPtrOutputWithContext(ctx context.Context) MetroPtrOutput {
	return o
}

func (o MetroPtrOutput) Elem() MetroOutput {
	return o.ApplyT(func(v *Metro) Metro {
		if v != nil {
			return *v
		}
		var ret Metro
		return ret
	}).(MetroOutput)
}

func (o MetroPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetroPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Metro) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MetroInput is an input type that accepts MetroArgs and MetroOutput values.
// You can construct a concrete instance of `MetroInput` via:
//
//	MetroArgs{...}
type MetroInput interface {
	pulumi.Input

	ToMetroOutput() MetroOutput
	ToMetroOutputWithContext(context.Context) MetroOutput
}

var metroPtrType = reflect.TypeOf((**Metro)(nil)).Elem()

type MetroPtrInput interface {
	pulumi.Input

	ToMetroPtrOutput() MetroPtrOutput
	ToMetroPtrOutputWithContext(context.Context) MetroPtrOutput
}

type metroPtr string

func MetroPtr(v string) MetroPtrInput {
	return (*metroPtr)(&v)
}

func (*metroPtr) ElementType() reflect.Type {
	return metroPtrType
}

func (in *metroPtr) ToMetroPtrOutput() MetroPtrOutput {
	return pulumi.ToOutput(in).(MetroPtrOutput)
}

func (in *metroPtr) ToMetroPtrOutputWithContext(ctx context.Context) MetroPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MetroPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetroInput)(nil)).Elem(), Metro("AM"))
	pulumi.RegisterInputType(reflect.TypeOf((*MetroPtrInput)(nil)).Elem(), Metro("AM"))
	pulumi.RegisterOutputType(MetroOutput{})
	pulumi.RegisterOutputType(MetroPtrOutput{})
}
