package wasihttp

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	monotonicclock "go.wasmcloud.dev/component/gen/wasi/clocks/monotonic-clock"
	outgoinghandler "go.wasmcloud.dev/component/gen/wasi/http/outgoing-handler"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
	"go.wasmcloud.dev/component/log/wasilog"
)

// Transport implements http.RoundTripper
type Transport struct {
	ConnectTimeout time.Duration
}

var Logger = wasilog.DefaultLogger

var _ http.RoundTripper = (*Transport)(nil)

var (
	DefaultTransport = &Transport{
		// NOTE(lxf): Same as stdlib http.Transport
		ConnectTimeout: 30 * time.Second,
	}
	DefaultClient = &http.Client{Transport: DefaultTransport}
)

func (r *Transport) requestOptions() types.RequestOptions {
	options := types.NewRequestOptions()
	options.SetConnectTimeout(cm.Some(monotonicclock.Duration(r.ConnectTimeout)))
	return options
}

func (r *Transport) RoundTrip(incomingRequest *http.Request) (*http.Response, error) {
	var err error

	outHeaders := types.NewFields()
	if err := toWasiHeader(incomingRequest.Header, outHeaders); err != nil {
		return nil, err
	}

	outRequest := types.NewOutgoingRequest(outHeaders)

	outRequest.SetAuthority(cm.Some(incomingRequest.Host))
	outRequest.SetMethod(toWasiMethod(incomingRequest.Method))
	outRequest.SetPathWithQuery(cm.Some(incomingRequest.URL.Path + "?" + incomingRequest.URL.Query().Encode()))

	switch incomingRequest.URL.Scheme {
	case "http":
		outRequest.SetScheme(cm.Some(types.SchemeHTTP()))
	case "https":
		outRequest.SetScheme(cm.Some(types.SchemeHTTPS()))
	default:
		outRequest.SetScheme(cm.Some(types.SchemeOther(incomingRequest.URL.Scheme)))
	}

	var adaptedBody io.WriteCloser
	bodyRes := outRequest.Body()
	if bodyRes.IsErr() {
		return nil, fmt.Errorf("failed to acquire resource handle to request body: %s", bodyRes.Err())
	}
	body := bodyRes.OK()

	if incomingRequest.Body != nil {
		adaptedBody, err = NewOutgoingBody(body)
		if err != nil {
			return nil, fmt.Errorf("failed to adapt body: %s", err)
		}
	}

	handleResp := outgoinghandler.Handle(outRequest, cm.Some(r.requestOptions()))
	if handleResp.Err() != nil {
		return nil, fmt.Errorf("%v", handleResp.Err())
	}

	// NOTE(lxf): If request includes a body, copy it to the adapted wasi body
	if adaptedBody != nil {
		Logger.Info("Proxying body")
		if _, err := io.Copy(adaptedBody, incomingRequest.Body); err != nil {
			return nil, fmt.Errorf("failed to copy body: %v", err)
		}
	}

	outTrailers := types.NewFields()
	if err := toWasiHeader(incomingRequest.Trailer, outTrailers); err != nil {
		return nil, err
	}

	outFinish := types.OutgoingBodyFinish(*body, cm.Some(outTrailers))
	if outFinish.IsErr() {
		return nil, fmt.Errorf("failed to set trailer: %v", outFinish.Err())
	}

	// NOTE(lxf): Request is fully sent. Processing response.
	top := handleResp.OK()
	defer top.ResourceDrop()

	// wait until resp is returned
	subscription := top.Subscribe()
	defer subscription.ResourceDrop()

	subscription.Block()

	pollableOption := top.Get()
	if pollableOption.None() {
		return nil, fmt.Errorf("incoming resp is None")
	}

	pollableResult := pollableOption.Some()
	if pollableResult.IsErr() {
		return nil, fmt.Errorf("error is %v", pollableResult.Err())
	}

	resultOption := pollableResult.OK()
	if resultOption.IsErr() {
		return nil, fmt.Errorf("%v", resultOption.Err())
	}

	incomingResponse := *resultOption.OK()
	incomingBody, incomingTrailers, err := NewIncomingBodyTrailer(incomingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to consume incoming request %s", err)
	}

	incomingHeaders := http.Header{}
	toHttpHeader(incomingResponse.Headers(), &incomingHeaders)

	resp := &http.Response{
		StatusCode: int(incomingResponse.Status()),
		Status:     http.StatusText(int(incomingResponse.Status())),
		Request:    incomingRequest,
		Header:     incomingHeaders,
		Body:       incomingBody,
		Trailer:    incomingTrailers,
	}

	return resp, nil
}
