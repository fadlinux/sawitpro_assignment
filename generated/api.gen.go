// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// CreateEstateRequest defines model for CreateEstateRequest.
type CreateEstateRequest struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

// CreateEstateResponse defines model for CreateEstateResponse.
type CreateEstateResponse struct {
	ID *string `json:"ID,omitempty"`
}

// CreateTreeRequest defines model for CreateTreeRequest.
type CreateTreeRequest struct {
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

// CreateTreeResponse defines model for CreateTreeResponse.
type CreateTreeResponse struct {
	ID *string `json:"ID,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// CreateEstateJSONRequestBody defines body for CreateEstate for application/json ContentType.
type CreateEstateJSONRequestBody = CreateEstateRequest

// CreateTreeJSONRequestBody defines body for CreateTree for application/json ContentType.
type CreateTreeJSONRequestBody = CreateTreeRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// This Endpoint For Create Estate
	// (POST /estate)
	CreateEstate(ctx echo.Context) error
	// APIs for stores tree data in a given estate with the ID, It receives three positive integers x, y, and height
	// (POST /estate/{Id}/tree)
	CreateTree(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// CreateEstate converts echo context to params.
func (w *ServerInterfaceWrapper) CreateEstate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateEstate(ctx)
	return err
}

// CreateTree converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTree(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "Id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "Id", runtime.ParamLocationPath, ctx.Param("Id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter Id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateTree(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/estate", wrapper.CreateEstate)
	router.POST(baseURL+"/estate/:Id/tree", wrapper.CreateTree)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWTXPbNhD9Kztoj7TJpHac8tZUboeHzHQa99TxASJWIjIkgCxWX6PRf+8sRFqW9WFn",
	"pk4uuVEQdt97i/dArlXtu+AdOo6qXKtYN9jp9Pg7oWa8jawZ/8YvM4wsy4F8QGKLaVOLbsqNPPEqoCqV",
	"dYxTJLXJ1MKa439tMkX4ZWYJjSr/7fdlQ6/7bCjw489Ys7Ta5xKDdxEPyVSjR2iRybrpAZg1ZwDuCE9L",
	"bdBOGz4udXl8efUC9Usl+7Kh/XPkXkH7LZGn0507jFFP8fn2w8ZDDNlp3cQnx9gaexynO9n1sboTGmy5",
	"lZ//RCT4hDS3NapMzZGi9U6V6s1lcVnITh/Q6WBVqX5JS5kKmptENsfkkaTCb09RtGi23lVGlXtOUlv+",
	"GPmDN+msau8YXSrTIbS2ToX55ygEhnTI08+EE1Wqn/JdfPI+O/mx4Gz2h8U0w7SwHXqi/rZ480oU+pNN",
	"HAzGmmzg7UjvMDIQ8oycDPaqKP43CvuuOoL9QRt4GE+mrr8lduUYyek2+QwJUkEydJx1naaVzKaxEW6d",
	"Cd46hj88wXaq8GAe1tMoxt+uR9DOQGRPGMHhAnor3kvb3pf5ujKbnAmfdaikPRmbdIeMJEBrhUvdhZSS",
	"G3yH1+P3eFEUN/riamyuLsbj2lxc34zfY/FOF29/NUpSp8qUDpUNcauMemrF7NFcJ546zbuUZwepv3/N",
	"2Dy+gr9LaPau2R+R+arI/PZXFWHiaUiBGB2MZg3WgYapnaPrYwELyw1wg1CNMqhkpDXauRQ1UhV8tGzn",
	"CP17M8Iyg1WWMta/Kg8DmGB3qBI94ZcIbwM0o1aVqmEOZZ63vtZtIyEUT/fN1k90P2iqz6VcBIoYgR3r",
	"iJeQro+FbdtBGfDC79Skj57Ua/vZkwFhIIzo2LopaAi67cDbFkKrHacj7bEud1k+f/NsspNivtEBPaH5",
	"9HzOEEyTi7YL7apPW+Ii1CL4SfqRGvWDf0y6Gu2g/0T+qsIzlE4wmnVgbGTtahz6G/IOofPOspdLE5j0",
	"HNuXcP308nbfnetHvYTR0G4kfdTmfvNfAAAA//+S+KIoUwwAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}