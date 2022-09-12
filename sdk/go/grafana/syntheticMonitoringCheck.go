// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package grafana

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Synthetic Monitoring checks are tests that run on selected probes at defined
// intervals and report metrics and logs back to your Grafana Cloud account. The
// target for checks can be a domain name, a server, or a website, depending on
// what information you would like to gather about your endpoint. You can define
// multiple checks for a single endpoint to check different capabilities.
//
// * [Official documentation](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/checks/)
//
// ## Example Usage
// ### DNS Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "dns", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("DNS Defaults"),
//				Target:  pulumi.String("grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Atlanta),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Dns: nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### DNS Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "dns", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("DNS Updated"),
//				Target:  pulumi.String("grafana.net"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Dns: &SyntheticMonitoringCheckSettingsDnsArgs{
//						IpVersion:  pulumi.String("Any"),
//						Server:     pulumi.String("8.8.4.4"),
//						Port:       pulumi.Int(8600),
//						RecordType: pulumi.String("CNAME"),
//						Protocol:   pulumi.String("TCP"),
//						ValidRCodes: pulumi.StringArray{
//							pulumi.String("NOERROR"),
//							pulumi.String("NOTAUTH"),
//						},
//						ValidateAnswerRrs: &SyntheticMonitoringCheckSettingsDnsValidateAnswerRrsArgs{
//							FailIfMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-bad-stuff*"),
//							},
//							FailIfNotMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-good-stuff*"),
//							},
//						},
//						ValidateAuthorityRrs: &SyntheticMonitoringCheckSettingsDnsValidateAuthorityRrsArgs{
//							FailIfMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-bad-stuff*"),
//							},
//							FailIfNotMatchesRegexps: pulumi.StringArray{
//								pulumi.String(".+-good-stuff*"),
//							},
//						},
//						ValidateAdditionalRrs: SyntheticMonitoringCheckSettingsDnsValidateAdditionalRrArray{
//							&SyntheticMonitoringCheckSettingsDnsValidateAdditionalRrArgs{
//								FailIfMatchesRegexps: pulumi.StringArray{
//									pulumi.String(".+-bad-stuff*"),
//								},
//								FailIfNotMatchesRegexps: pulumi.StringArray{
//									pulumi.String(".+-good-stuff*"),
//								},
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### HTTP Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "http", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("HTTP Defaults"),
//				Target:  pulumi.String("https://grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Atlanta),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Http: nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### HTTP Complex
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "http", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("HTTP Defaults"),
//				Target:  pulumi.String("https://grafana.org"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Bangalore),
//					pulumi.Int(main.Probes.Mumbai),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Http: &SyntheticMonitoringCheckSettingsHttpArgs{
//						IpVersion:                  pulumi.String("V6"),
//						Method:                     pulumi.String("TRACE"),
//						Body:                       pulumi.String("and spirit"),
//						NoFollowRedirects:          pulumi.Bool(true),
//						BearerToken:                pulumi.String("asdfjkl;"),
//						ProxyUrl:                   pulumi.String("https://almost-there"),
//						FailIfSsl:                  pulumi.Bool(true),
//						FailIfNotSsl:               pulumi.Bool(true),
//						CacheBustingQueryParamName: pulumi.String("pineapple"),
//						TlsConfig: &SyntheticMonitoringCheckSettingsHttpTlsConfigArgs{
//							ServerName: pulumi.String("grafana.org"),
//							ClientCert: pulumi.String(fmt.Sprintf(`-----BEGIN CERTIFICATE-----
//
// MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
// RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
// MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
// 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
// h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
// BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
// iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
// a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
// FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
// qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
// FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
// Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
// 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
// UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
// yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
// e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
// XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
// tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
// QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
// tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
// prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
// 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
// l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
// 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
// vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
// -----END CERTIFICATE-----
// `)),
//
//						},
//						Headers: pulumi.StringArray{
//							pulumi.String("Content-Type: multipart/form-data; boundary=something"),
//						},
//						BasicAuth: &SyntheticMonitoringCheckSettingsHttpBasicAuthArgs{
//							Username: pulumi.String("open"),
//							Password: pulumi.String("sesame"),
//						},
//						ValidStatusCodes: pulumi.IntArray{
//							pulumi.Int(200),
//							pulumi.Int(201),
//						},
//						ValidHttpVersions: pulumi.StringArray{
//							pulumi.String("HTTP/2"),
//						},
//						FailIfBodyMatchesRegexps: pulumi.StringArray{
//							pulumi.String("*bad stuff*"),
//						},
//						FailIfBodyNotMatchesRegexps: pulumi.StringArray{
//							pulumi.String("*good stuff*"),
//						},
//						FailIfHeaderMatchesRegexps: SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArray{
//							&SyntheticMonitoringCheckSettingsHttpFailIfHeaderMatchesRegexpArgs{
//								Header:       pulumi.String("Content-Type"),
//								Regexp:       pulumi.String("application/soap*"),
//								AllowMissing: pulumi.Bool(true),
//							},
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Ping Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "ping", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("Ping Defaults"),
//				Target:  pulumi.String("grafana.com"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Atlanta),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Ping: nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Ping Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "ping", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("Ping Updated"),
//				Target:  pulumi.String("grafana.net"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Ping: &SyntheticMonitoringCheckSettingsPingArgs{
//						IpVersion:    pulumi.String("Any"),
//						PayloadSize:  pulumi.Int(20),
//						DontFragment: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### TCP Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "tcp", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("TCP Defaults"),
//				Target:  pulumi.String("grafana.com:80"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Atlanta),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Tcp: nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### TCP Complex
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "tcp", &grafana.SyntheticMonitoringCheckArgs{
//				Job:     pulumi.String("TCP Defaults"),
//				Target:  pulumi.String("grafana.com:443"),
//				Enabled: pulumi.Bool(false),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Tcp: &SyntheticMonitoringCheckSettingsTcpArgs{
//						IpVersion: pulumi.String("V6"),
//						Tls:       pulumi.Bool(true),
//						QueryResponses: SyntheticMonitoringCheckSettingsTcpQueryResponseArray{
//							&SyntheticMonitoringCheckSettingsTcpQueryResponseArgs{
//								Send:   pulumi.String("howdy"),
//								Expect: pulumi.String("hi"),
//							},
//							&SyntheticMonitoringCheckSettingsTcpQueryResponseArgs{
//								Send:     pulumi.String("like this"),
//								Expect:   pulumi.String("like that"),
//								StartTls: pulumi.Bool(true),
//							},
//						},
//						TlsConfig: &SyntheticMonitoringCheckSettingsTcpTlsConfigArgs{
//							ServerName: pulumi.String("grafana.com"),
//							CaCert: pulumi.String(fmt.Sprintf(`-----BEGIN CERTIFICATE-----
//
// MIIEljCCAn4CCQCKJPUQQxeO0zANBgkqhkiG9w0BAQsFADANMQswCQYDVQQGEwJT
// RTAeFw0yMTA1MjkxOTIyNTdaFw0yNDAzMTgxOTIyNTdaMA0xCzAJBgNVBAYTAlNF
// MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnmbazDNUT0rSI4BpGZK+
// 0AJ+9FDkIYWJUtRLJoxw8CF+AobMFploYA2L2Myt80cTA1w8FrewjC8qlqdnrPWr
// h1ely2zsUljgi1/niH0ndjFzliL7UkinXQiAsTtYOrOQmzyd/o5PNdu7dz0m7stD
// BN/Sz5TlXZnA1/eJbqV/kqMau6b1MaBx8SbRfUG9+cSmUobFJwuktDrPuwJhcEkl
// iDmhEqu1GuZzmKvzPacLTVia1vSlmCTCu89NiHI8iGiiLtqNrapup7f8j5m3a3SL
// a+vXhplFj2piNl7Nc0dfuVgtEliTI+qUL2/+4A7gzRWZpHy21/LxMMXmBhdJW9En
// FWkev97VZLgb5TR3+qpSWmXcodjPy4dibvwsOMpdd+Q4AYulwvlDw5idRPVgGvk7
// qq03+w9ppZ5Fugws9k2CD9F/75JX2mCbRpkuPe8XXZ7bqrMaQgQMLOrs68HuiiCk
// FTklglq4DMKxnf/Y/T/MgIa9Q1o28YSevh6A7FnfPGARj2H2T4rToi+bC1Vf7qNB
// Z18bDpz99tRUTbyiRUSBMWLCGhU6c4HAqUrfrkpperOKFBQ3i38a79838oFdXHBW
// 6rx1t5cC3XwtEoUyeBKAygez8G1LDXbN3607MxVhAjhHKtPkYvuBfysSNU6JrR0z
// UV1IURJANt2UMuKgSEkG/IMCAwEAATANBgkqhkiG9w0BAQsFAAOCAgEAcipMhp/w
// yzfPy61faVAw9SPaMNRlnW9FCDC3N9CGOjo2knjXpObPzyzsJiUURTjrA9eFMpRA
// e2Rgn2j+nvm2XdLAlC4Kh8jqv/wCL0X6BTQMdN5aOhXdSiXtpXOMvXYY/dQ4ebRZ
// XeRCVWQD79JbV6/uyx0nCV3FVcU7L1P4UjxroefVr0soLPMirgxHmOxLnkoVgdcB
// tqufP5kJx9CIeJXPx3QQsk1XfEtxtUvuw4ZaZkQnNUqvGl7V+AZpur5Eqfv3zBi8
// QxxL7qGkARNssNWH2Ju+tqpM/UZRnjlFrDR4SXUgT0coTduBalUY6qHkciHmRpiP
// tf3SgpDeiCSOV2iVFGdaR1mz3muWoAYWFstcWN3a3HjjVugIi23yLN8Gv8CNeoH4
// prulinFCLrFgAh8SLAF8mOAZanT06LH8jOIFYrdUxH+ZeRBR0rLoFjUF+JB7UKD9
// 5TA+B4EBzQ1tMbGFU1DX79MjAejq0IV0Nzq+GMfBvLHxEf4+Oz8nqhDXQcJ6TdtY
// l3Lyw5zBvOL80SBK+Mr0UP7d9U3VXgbGHCYVJU6Ot1TwiGwahtWALRALA3TWeGkq
// 7kyD1H+nm+9lfKhuyBRQnRGBVyze2lAp7oxwshJuhBwEXosXFxq1Cy6QhPN77r6N
// vuhxvtppolNnyOgGxwG4zquqq2V5/+vKjKY=
// -----END CERTIFICATE-----
// `)),
//
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Traceroute Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "traceroute", &grafana.SyntheticMonitoringCheckArgs{
//				Job:       pulumi.String("Traceroute defaults"),
//				Target:    pulumi.String("grafana.com"),
//				Enabled:   pulumi.Bool(false),
//				Frequency: pulumi.Int(120000),
//				Timeout:   pulumi.Int(30000),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Atlanta),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Traceroute: nil,
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Traceroute Complex
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi-grafana/sdk/go/grafana"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			main, err := grafana.GetSyntheticMonitoringProbes(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			_, err = grafana.NewSyntheticMonitoringCheck(ctx, "traceroute", &grafana.SyntheticMonitoringCheckArgs{
//				Job:       pulumi.String("Traceroute complex"),
//				Target:    pulumi.String("grafana.net"),
//				Enabled:   pulumi.Bool(false),
//				Frequency: pulumi.Int(120000),
//				Timeout:   pulumi.Int(30000),
//				Probes: pulumi.IntArray{
//					pulumi.Int(main.Probes.Frankfurt),
//					pulumi.Int(main.Probes.London),
//				},
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("baz"),
//				},
//				Settings: &SyntheticMonitoringCheckSettingsArgs{
//					Traceroute: &SyntheticMonitoringCheckSettingsTracerouteArgs{
//						MaxHops:        pulumi.Int(25),
//						MaxUnknownHops: pulumi.Int(10),
//						PtrLookup:      pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck check {{check-id}}
//
// ```
type SyntheticMonitoringCheck struct {
	pulumi.CustomResourceState

	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrOutput `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrOutput `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrOutput `pulumi:"frequency"`
	// Name used for job label.
	Job pulumi.StringOutput `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayOutput `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings SyntheticMonitoringCheckSettingsOutput `pulumi:"settings"`
	// Hostname to ping.
	Target pulumi.StringOutput `pulumi:"target"`
	// The tenant ID of the check.
	TenantId pulumi.IntOutput `pulumi:"tenantId"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
}

// NewSyntheticMonitoringCheck registers a new resource with the given unique name, arguments, and options.
func NewSyntheticMonitoringCheck(ctx *pulumi.Context,
	name string, args *SyntheticMonitoringCheckArgs, opts ...pulumi.ResourceOption) (*SyntheticMonitoringCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Job == nil {
		return nil, errors.New("invalid value for required argument 'Job'")
	}
	if args.Probes == nil {
		return nil, errors.New("invalid value for required argument 'Probes'")
	}
	if args.Settings == nil {
		return nil, errors.New("invalid value for required argument 'Settings'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SyntheticMonitoringCheck
	err := ctx.RegisterResource("grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyntheticMonitoringCheck gets an existing SyntheticMonitoringCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyntheticMonitoringCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyntheticMonitoringCheckState, opts ...pulumi.ResourceOption) (*SyntheticMonitoringCheck, error) {
	var resource SyntheticMonitoringCheck
	err := ctx.ReadResource("grafana:index/syntheticMonitoringCheck:SyntheticMonitoringCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyntheticMonitoringCheck resources.
type syntheticMonitoringCheckState struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity *string `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly *bool `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
	Frequency *int `pulumi:"frequency"`
	// Name used for job label.
	Job *string `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels map[string]string `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes []int `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings *SyntheticMonitoringCheckSettings `pulumi:"settings"`
	// Hostname to ping.
	Target *string `pulumi:"target"`
	// The tenant ID of the check.
	TenantId *int `pulumi:"tenantId"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout *int `pulumi:"timeout"`
}

type SyntheticMonitoringCheckState struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrInput
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrInput
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrInput
	// Name used for job label.
	Job pulumi.StringPtrInput
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapInput
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayInput
	// Check settings. Should contain exactly one nested block.
	Settings SyntheticMonitoringCheckSettingsPtrInput
	// Hostname to ping.
	Target pulumi.StringPtrInput
	// The tenant ID of the check.
	TenantId pulumi.IntPtrInput
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrInput
}

func (SyntheticMonitoringCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*syntheticMonitoringCheckState)(nil)).Elem()
}

type syntheticMonitoringCheckArgs struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity *string `pulumi:"alertSensitivity"`
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly *bool `pulumi:"basicMetricsOnly"`
	// Whether to enable the check. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
	Frequency *int `pulumi:"frequency"`
	// Name used for job label.
	Job string `pulumi:"job"`
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels map[string]string `pulumi:"labels"`
	// List of probe location IDs where this target will be checked from.
	Probes []int `pulumi:"probes"`
	// Check settings. Should contain exactly one nested block.
	Settings SyntheticMonitoringCheckSettings `pulumi:"settings"`
	// Hostname to ping.
	Target string `pulumi:"target"`
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout *int `pulumi:"timeout"`
}

// The set of arguments for constructing a SyntheticMonitoringCheck resource.
type SyntheticMonitoringCheckArgs struct {
	// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
	AlertSensitivity pulumi.StringPtrInput
	// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
	BasicMetricsOnly pulumi.BoolPtrInput
	// Whether to enable the check. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
	Frequency pulumi.IntPtrInput
	// Name used for job label.
	Job pulumi.StringInput
	// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
	Labels pulumi.StringMapInput
	// List of probe location IDs where this target will be checked from.
	Probes pulumi.IntArrayInput
	// Check settings. Should contain exactly one nested block.
	Settings SyntheticMonitoringCheckSettingsInput
	// Hostname to ping.
	Target pulumi.StringInput
	// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
	Timeout pulumi.IntPtrInput
}

func (SyntheticMonitoringCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syntheticMonitoringCheckArgs)(nil)).Elem()
}

type SyntheticMonitoringCheckInput interface {
	pulumi.Input

	ToSyntheticMonitoringCheckOutput() SyntheticMonitoringCheckOutput
	ToSyntheticMonitoringCheckOutputWithContext(ctx context.Context) SyntheticMonitoringCheckOutput
}

func (*SyntheticMonitoringCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**SyntheticMonitoringCheck)(nil)).Elem()
}

func (i *SyntheticMonitoringCheck) ToSyntheticMonitoringCheckOutput() SyntheticMonitoringCheckOutput {
	return i.ToSyntheticMonitoringCheckOutputWithContext(context.Background())
}

func (i *SyntheticMonitoringCheck) ToSyntheticMonitoringCheckOutputWithContext(ctx context.Context) SyntheticMonitoringCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringCheckOutput)
}

// SyntheticMonitoringCheckArrayInput is an input type that accepts SyntheticMonitoringCheckArray and SyntheticMonitoringCheckArrayOutput values.
// You can construct a concrete instance of `SyntheticMonitoringCheckArrayInput` via:
//
//	SyntheticMonitoringCheckArray{ SyntheticMonitoringCheckArgs{...} }
type SyntheticMonitoringCheckArrayInput interface {
	pulumi.Input

	ToSyntheticMonitoringCheckArrayOutput() SyntheticMonitoringCheckArrayOutput
	ToSyntheticMonitoringCheckArrayOutputWithContext(context.Context) SyntheticMonitoringCheckArrayOutput
}

type SyntheticMonitoringCheckArray []SyntheticMonitoringCheckInput

func (SyntheticMonitoringCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyntheticMonitoringCheck)(nil)).Elem()
}

func (i SyntheticMonitoringCheckArray) ToSyntheticMonitoringCheckArrayOutput() SyntheticMonitoringCheckArrayOutput {
	return i.ToSyntheticMonitoringCheckArrayOutputWithContext(context.Background())
}

func (i SyntheticMonitoringCheckArray) ToSyntheticMonitoringCheckArrayOutputWithContext(ctx context.Context) SyntheticMonitoringCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringCheckArrayOutput)
}

// SyntheticMonitoringCheckMapInput is an input type that accepts SyntheticMonitoringCheckMap and SyntheticMonitoringCheckMapOutput values.
// You can construct a concrete instance of `SyntheticMonitoringCheckMapInput` via:
//
//	SyntheticMonitoringCheckMap{ "key": SyntheticMonitoringCheckArgs{...} }
type SyntheticMonitoringCheckMapInput interface {
	pulumi.Input

	ToSyntheticMonitoringCheckMapOutput() SyntheticMonitoringCheckMapOutput
	ToSyntheticMonitoringCheckMapOutputWithContext(context.Context) SyntheticMonitoringCheckMapOutput
}

type SyntheticMonitoringCheckMap map[string]SyntheticMonitoringCheckInput

func (SyntheticMonitoringCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyntheticMonitoringCheck)(nil)).Elem()
}

func (i SyntheticMonitoringCheckMap) ToSyntheticMonitoringCheckMapOutput() SyntheticMonitoringCheckMapOutput {
	return i.ToSyntheticMonitoringCheckMapOutputWithContext(context.Background())
}

func (i SyntheticMonitoringCheckMap) ToSyntheticMonitoringCheckMapOutputWithContext(ctx context.Context) SyntheticMonitoringCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyntheticMonitoringCheckMapOutput)
}

type SyntheticMonitoringCheckOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyntheticMonitoringCheck)(nil)).Elem()
}

func (o SyntheticMonitoringCheckOutput) ToSyntheticMonitoringCheckOutput() SyntheticMonitoringCheckOutput {
	return o
}

func (o SyntheticMonitoringCheckOutput) ToSyntheticMonitoringCheckOutputWithContext(ctx context.Context) SyntheticMonitoringCheckOutput {
	return o
}

// Can be set to `none`, `low`, `medium`, or `high` to correspond to the check [alert levels](https://grafana.com/docs/grafana-cloud/synthetic-monitoring/synthetic-monitoring-alerting/). Defaults to `none`.
func (o SyntheticMonitoringCheckOutput) AlertSensitivity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.StringPtrOutput { return v.AlertSensitivity }).(pulumi.StringPtrOutput)
}

// Metrics are reduced by default. Set this to `false` if you'd like to publish all metrics. We maintain a [full list of metrics](https://github.com/grafana/synthetic-monitoring-agent/tree/main/internal/scraper/testdata) collected for each. Defaults to `true`.
func (o SyntheticMonitoringCheckOutput) BasicMetricsOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.BoolPtrOutput { return v.BasicMetricsOnly }).(pulumi.BoolPtrOutput)
}

// Whether to enable the check. Defaults to `true`.
func (o SyntheticMonitoringCheckOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// How often the check runs in milliseconds (the value is not truly a "frequency" but a "period"). The minimum acceptable value is 1 second (1000 ms), and the maximum is 120 seconds (120000 ms). Defaults to `60000`.
func (o SyntheticMonitoringCheckOutput) Frequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.IntPtrOutput { return v.Frequency }).(pulumi.IntPtrOutput)
}

// Name used for job label.
func (o SyntheticMonitoringCheckOutput) Job() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.StringOutput { return v.Job }).(pulumi.StringOutput)
}

// Custom labels to be included with collected metrics and logs. The maximum number of labels that can be specified per check is 5. These are applied, along with the probe-specific labels, to the outgoing metrics. The names and values of the labels cannot be empty, and the maximum length is 32 bytes.
func (o SyntheticMonitoringCheckOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// List of probe location IDs where this target will be checked from.
func (o SyntheticMonitoringCheckOutput) Probes() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.IntArrayOutput { return v.Probes }).(pulumi.IntArrayOutput)
}

// Check settings. Should contain exactly one nested block.
func (o SyntheticMonitoringCheckOutput) Settings() SyntheticMonitoringCheckSettingsOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) SyntheticMonitoringCheckSettingsOutput { return v.Settings }).(SyntheticMonitoringCheckSettingsOutput)
}

// Hostname to ping.
func (o SyntheticMonitoringCheckOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// The tenant ID of the check.
func (o SyntheticMonitoringCheckOutput) TenantId() pulumi.IntOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.IntOutput { return v.TenantId }).(pulumi.IntOutput)
}

// Specifies the maximum running time for the check in milliseconds. The minimum acceptable value is 1 second (1000 ms), and the maximum 10 seconds (10000 ms). Defaults to `3000`.
func (o SyntheticMonitoringCheckOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyntheticMonitoringCheck) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

type SyntheticMonitoringCheckArrayOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyntheticMonitoringCheck)(nil)).Elem()
}

func (o SyntheticMonitoringCheckArrayOutput) ToSyntheticMonitoringCheckArrayOutput() SyntheticMonitoringCheckArrayOutput {
	return o
}

func (o SyntheticMonitoringCheckArrayOutput) ToSyntheticMonitoringCheckArrayOutputWithContext(ctx context.Context) SyntheticMonitoringCheckArrayOutput {
	return o
}

func (o SyntheticMonitoringCheckArrayOutput) Index(i pulumi.IntInput) SyntheticMonitoringCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyntheticMonitoringCheck {
		return vs[0].([]*SyntheticMonitoringCheck)[vs[1].(int)]
	}).(SyntheticMonitoringCheckOutput)
}

type SyntheticMonitoringCheckMapOutput struct{ *pulumi.OutputState }

func (SyntheticMonitoringCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyntheticMonitoringCheck)(nil)).Elem()
}

func (o SyntheticMonitoringCheckMapOutput) ToSyntheticMonitoringCheckMapOutput() SyntheticMonitoringCheckMapOutput {
	return o
}

func (o SyntheticMonitoringCheckMapOutput) ToSyntheticMonitoringCheckMapOutputWithContext(ctx context.Context) SyntheticMonitoringCheckMapOutput {
	return o
}

func (o SyntheticMonitoringCheckMapOutput) MapIndex(k pulumi.StringInput) SyntheticMonitoringCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyntheticMonitoringCheck {
		return vs[0].(map[string]*SyntheticMonitoringCheck)[vs[1].(string)]
	}).(SyntheticMonitoringCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringCheckInput)(nil)).Elem(), &SyntheticMonitoringCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringCheckArrayInput)(nil)).Elem(), SyntheticMonitoringCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyntheticMonitoringCheckMapInput)(nil)).Elem(), SyntheticMonitoringCheckMap{})
	pulumi.RegisterOutputType(SyntheticMonitoringCheckOutput{})
	pulumi.RegisterOutputType(SyntheticMonitoringCheckArrayOutput{})
	pulumi.RegisterOutputType(SyntheticMonitoringCheckMapOutput{})
}
