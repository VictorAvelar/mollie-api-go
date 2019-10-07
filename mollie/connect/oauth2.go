package connect

import (
	"golang.org/x/oauth2"
)

const (
	AuthURL   = "https://www.mollie.com/oauth2/authorize"
	TokensURL = "https://api.mollie.com/oauth2/tokens"
)

//  Endpoint is Mollies's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   AuthURL,
	TokenURL:  TokensURL,
	AuthStyle: 0,
}

// func (c *Config) TokenSource(ctx context.Context) oauth2.TokenSource {
// 	return oauth2.ReuseTokenSource(nil, jwtSource{
// 		ctx:  ctx,
// 		conf: *c,
// 	})
// }
//
// func (c *Config) Client(ctx context.Context) *http.Client {
// 	return oauth2.NewClient(ctx, c.TokenSource(ctx))
// }
//
// type jwtSource struct {
// 	ctx context.Context
// 	conf Config
// }
//
// func (js jwtSource) Token() (*oauth2.Token, error) {
// 	mc := oauth2.NewClient(js.ctx, nil)
//
// 	res, err := mc.PostForm(js.conf.Endpoint.TokenURL, url.Values{})
// 	if err != nil {
// 		return nil, fmt.Errorf("oauth2: cannot fetch token: %v", err)
// 	}
//
// 	err = mollie.CheckResponse(res)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 1<<20))
//
// 	var token oauth2.Token
// 	if err = json.Unmarshal(body, &token); err != nil {
// 		return nil, err
// 	}
//
// 	return &token, nil
// }
