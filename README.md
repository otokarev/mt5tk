# MetaTrader5 Web API command line client

It can:
* List all groups available
* Create group from template
* Delete groups
* Duplicate groups
* Modify groups' symbols (add, replace, remove)
* List all symbols

Also it has nice CLI based on https://github.com/spf13/cobra and cute JSONpath output formatter taken from https://github.com/kubernetes/client-go.

## Basic usage
### Configuration
Copy the config sample into your `$HOME` and put in it correct settings 
```shell
cp configs/.mt5tk.yml.example ~/.mt5tk.yml
```
### Get group
```shell
mt5tk group get -g 'test\test-BTC'
```
### Get symbol
```shell
mt5tk symbol get --symbol EURUSD
```
### List all groups
```
mt5tk group list
```
### List all symbols
Listing with output customized with `jsonpath` option (usage hints: https://kubernetes.io/docs/reference/kubectl/jsonpath/)
```
mt5tk symbol list -o jsonpath='{range [*]} "{.Symbol}"`s description: "{.Description}"{"\n"}{end}'
```
### Create new group
```shell
mt5tk group create -g 'test\test-test' -t assets/group-template.json
```

### Duplicate group
Make a copy with name `test\test-BTC` for group `test\test-USD`
```shell
mt5tk group duplicate -g 'test\test-USD' -n 'test\test-BTC'
```

### Add symbol to the group
```shell
mt5tk  group symbol add --group 'test\test-EUR' --symbol EURUSD --template-path assets/group-symbol-template.json
```

### Add all available symbols into the group
```shell
mt5tk group fill-with-symbols  -g 'test\test-test' -t assets/group-symbol-template.json
```
---

More details on flags and commands you can find playing around with help messages of the app.