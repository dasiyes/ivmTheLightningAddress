# The giudes to follow when designing the project
- [The Lightning address offical website](https://lightningaddress.com/)

- [self-hosted solution DIY](https://github.com/andrerfneves/lightning-address/blob/master/DIY.md)

- [LNURL-pay](https://github.com/lnurl/luds/blob/legacy/lnurl-pay.md)

# How to compose the `metadata` response:

`metadata` json array must contain one `text/plain` entry, all other types of entries are optional.
`metadata` json array must contain either one `text/email` entry or one `text/identifier` entry or nethier.
`metadata` json array must contain either one `image/png;base64` entry or one `image/jpeg;base64` entry or nethier.

```json
[
    [
        "text/plain", // must always be present
        content // actual metadata content
    ],
    [
        "image/png;base64", // optional 512x512px PNG thumbnail which will represent this lnurl in a list or grid
        content // base64 string, up to 136536 characters (100Kb of image data in base-64 encoding)
    ],
    [
        "image/jpeg;base64", // optional 512x512px JPG thumbnail which will represent this lnurl in a list or grid
        content // base64 string, up to 136536 characters (100Kb of image data in base-64 encoding)
    ],
    [
        "text/email", // optional indication that this payment link is associated with an email address
        content // an email string in standard user@site.com format
    ],
    [
        "text/identifier", // optional indication that this payment link is associated with an internet identifier string
        content // an internet identifier string in standard user@site.com format
    ]
    ... // more objects for future types
]
```

and be sent as a string:

```
"[[\"text/plain\", \"lorem ipsum blah blah\"]]"
```
