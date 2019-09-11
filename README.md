# goify-multipass
Implementation of shopify's multipass in golang

[Shopify](http://shopify.com) provides a mechanism for single sign-on known as Multipass.  Multipass uses an AES encrypted JSON hash and multipassify provides functions for generating tokens

More details on Multipass with Shopify can be found [here](http://docs.shopify.com/api/tutorials/multipass-login).

## Installation
<pre>
    go get github.com/claimh-solais/goify-multipass
</pre>

## Usage

To use Multipass an Enterprise / Plus plan is required. The Multipass secret can be found in your shop Admin (Settings > Checkout > Customer Accounts).
Make sure "Accounts are required" or "Accounts are optional" is selected and Multipass is enabled.

``` go
    import multipass "github.com/claimh-solais/goify-multipass"

    // Construct the Multipassify encoder
    m := multipass.New("SHOPIFY MULTIPASS SECRET");

    // Create your customer info hash
    customerInfo := new(multipass.CustomerInfo{
        Email: "test@example.com",
        RemoteIP: "USERS IP ADDRESS",
        ReturnTo: "http://some.url",
    })

    // Encode a Multipass token
    token := m.Encode(customerInfo)

    // Generate a Shopify multipass URL to your shop
    urlString := m.GenerateURL(customerInfo, "yourstorename.myshopify.com");

  // Generates a URL like:  https://yourstorename.myshopify.com/account/login/multipass/<MULTIPASS-TOKEN>
```
