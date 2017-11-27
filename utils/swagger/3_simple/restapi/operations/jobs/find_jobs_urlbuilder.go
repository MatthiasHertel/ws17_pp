// Code generated by go-swagger; DO NOT EDIT.

package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// FindJobsURL generates an URL for the find jobs operation
type FindJobsURL struct {
	Page     *int64
	Pagesize *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindJobsURL) WithBasePath(bp string) *FindJobsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *FindJobsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *FindJobsURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/"

	_basePath := o._basePath
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var page string
	if o.Page != nil {
		page = swag.FormatInt64(*o.Page)
	}
	if page != "" {
		qs.Set("page", page)
	}

	var pagesize string
	if o.Pagesize != nil {
		pagesize = swag.FormatInt64(*o.Pagesize)
	}
	if pagesize != "" {
		qs.Set("pagesize", pagesize)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *FindJobsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *FindJobsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *FindJobsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on FindJobsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on FindJobsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *FindJobsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
