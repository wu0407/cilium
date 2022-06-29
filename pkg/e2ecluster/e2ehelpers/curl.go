// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package e2ehelpers

import (
	"fmt"
)

const (
	defaultConnectTimeout = 5  // seconds
	defaultMaxTime        = 20 // seconds
)

type CurlResultFormat string

const (
	CurlResultFormatStats    CurlResultFormat = `time-> DNS: '%{time_namelookup}(%{remote_ip})', Connect: '%{time_connect}', Transfer '%{time_starttransfer}', total '%{time_total}'`
	CurlResultFormatHTTPCode CurlResultFormat = "%%{http_code}"
)

type CurlOpts struct {
	Fail         bool
	ResultFormat CurlResultFormat
	Output       string
	Retries      int
	// ConnectTimeout is the timeout in seconds for the connect() syscall that curl invokes.
	ConnectTimeout int
	// MaxTime is the hard timeout. It starts when curl is invoked and interrupts curl
	// regardless of whether curl is currently connecting or transferring data. CurlMaxTimeout
	// should be at least 5 seconds longer than ConnectTimeout to provide some time to actually
	// transfer data.
	MaxTime        int
	AdditionalOpts []string
}

type CurlOption func(*CurlOpts)

func WithFail(fail bool) CurlOption {
	return func(o *CurlOpts) { o.Fail = fail }
}

func WithResultFormat(outputFormat CurlResultFormat) CurlOption {
	return func(o *CurlOpts) { o.ResultFormat = outputFormat }
}

func WithOutput(output string) CurlOption {
	return func(o *CurlOpts) { o.Output = output }
}

func WithRetries(retries int) CurlOption {
	return func(o *CurlOpts) { o.Retries = retries }
}

func WithConnectTimeout(connectTimeout int) CurlOption {
	return func(o *CurlOpts) { o.ConnectTimeout = connectTimeout }
}

func WithMaxTime(maxTime int) CurlOption {
	return func(o *CurlOpts) { o.MaxTime = maxTime }
}

func WithAdditionalOpts(opts []string) CurlOption {
	return func(o *CurlOpts) { o.AdditionalOpts = opts }
}

func processCurlOpts(opts ...CurlOption) *CurlOpts {
	o := &CurlOpts{
		ConnectTimeout: defaultConnectTimeout,
		MaxTime:        defaultMaxTime,
	}
	for _, op := range opts {
		op(o)
	}
	return o
}

func Curl(url string, opts ...CurlOption) []string {
	o := processCurlOpts(opts...)

	var cmd []string
	cmd = append(cmd, "curl", "--path-as-is", "-s", "-D /dev/stderr")
	if o.Fail {
		cmd = append(cmd, "--fail")
	}
	if o.ResultFormat != "" {
		cmd = append(cmd, fmt.Sprintf("-w %q", o.ResultFormat))
	}
	if o.Output != "" {
		cmd = append(cmd, fmt.Sprintf("--output %s", o.Output))
	}
	if o.Retries > 0 {
		cmd = append(cmd, fmt.Sprintf("--retry %d", o.Retries))
	}
	if o.ConnectTimeout > 0 {
		cmd = append(cmd, fmt.Sprintf("--connect-timeout %d", o.ConnectTimeout))
	}
	if o.MaxTime > 0 {
		cmd = append(cmd, fmt.Sprintf("--max-time %d", o.MaxTime))
	}
	if len(o.AdditionalOpts) > 0 {
		cmd = append(cmd, o.AdditionalOpts...)
	}
	cmd = append(cmd, url)
	return cmd
}
