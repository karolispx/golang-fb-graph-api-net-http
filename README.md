# Get accounts from Facebook Graph API using super user access token and net/http package.

## Code put together using net/http package:
* https://golang.org/pkg/net/http/

## Usage:
* You will need to initialize your own struct (GraphAPIResponse) for the reponse based on what you're expecting to get back.

1. Use Graph API explorer (https://developers.facebook.com/tools/explorer/) to test Graph API to see what you need to retrieve, in my case it was accounts (/me/adaccounts?limit=500&fields=id,name)
2. Click on `</> Get Code` and select cURL
3. Run cURL command in the terminal
4. Assuming data came back, copy everything starting with `{"data"` and ending with `}}}`
5. Paste it into https://mholt.github.io/json-to-go/ to get the struct

## You can easily extend this functionality to get other things such as campaigns, adsets etc.

The code will work for campaigns and adsets already, assuming you only need campaign/adset IDs and names. For other calls you will need to generate your own logic based on the sample above.