# Get accounts from Facebook Graph API using super user access token.

## Code put together using this package:
* https://github.com/huandu/facebook (you can find more examples and information about it)

You will need to get the huandu/facebook package to make it work:
* `go get -u github.com/huandu/facebook`

## You can easily extend this functionality to get other things such as:
### Campaigns by Account ID:
```
    // Account ID = 123456789
	res, _ := fb.Get("123456789/campaigns?limit=1000", fb.Params{
		"fields":       "id,name",
		"access_token": FB_ACCESS_TOKEN,
	})
```

### Adsets by Campaign ID:
```
    // Campaign ID = 123456789
	res, _ := fb.Get("123456789/adsets?limit=1000", fb.Params{
		"fields":       "id,name",
		"access_token": FB_ACCESS_TOKEN,
	})
```

### Adinsights by Adset ID:
```
    // Adset ID = 123456789
    // Fields: ad_id & spend
	res, _ := fb.Get("/123456789/insights?&fields=ad_id%2Cspend", fb.Params{
		"access_token": FB_ACCESS_TOKEN,
	})
```
