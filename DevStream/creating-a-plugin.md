# Create a plugin

## 0 Thank you for your contribution!

First, please read our [CONTRIBUTING](https://github.com/devstream-io/devstream/blob/main/CONTRIBUTING.md) documentation.

## 1 Automatically create the code framework for a new plugin

Run `dtm develop create-plugin --name=YOUR-PLUGIN-NAME` , dtm will automatically generate the following file.

> ### /cmd/plugin/YOUR-PLUGIN-NAME/main.go

This is the only main entry point for the plugin code.


You do not need to modify this file. If you feel there is a problem with this automatically generated file, you can create a PR to modify [template] directly (https://github.com/devstream-io/devstream/blob/main/internal/pkg/develop/plugin/template/main .go).

> ### /docs/plugins/YOUR-PLUGIN-NAME.md

This is the documentation for the automatically generated plugins.

Although the purpose of `dtm` is to automate, it doesn't magically generate the documentation. You need to write your own documentation for this plugin that you want to create.

> ###/internal/pkg/plugin/YOUR-PLUGIN-NAME/

Please write the main logic of the plugin here.

You can check our [Standard Go Project Layout](project-layout.md) file for detailed instructions on project layout.


## 2 Interfaces

### 2.1 Definitions

Each plugin needs to implement all the interfaces defined in [pluginengine](https://github.com/devstream-io/devstream/blob/main/internal/pkg/pluginengine/plugin.go#L10).

Currently, there are 4 interfaces, which are subject to change. Currently, the 4 interfaces are.

- [`create`](https://github.com/devstream-io/devstream/blob/main/internal/pkg/pluginengine/plugin.go#L12)
- [`read`](https://github.com/devstream-io/devstream/blob/main/internal/pkg/pluginengine/plugin.go#L13)
- [`update`](https://github.com/devstream-io/devstream/blob/main/internal/pkg/pluginengine/plugin.go#L14)
- [`delete`](https://github.com/devstream-io/devstream/blob/main/internal/pkg/pluginengine/plugin.go#L16)

### 2.2 Return values

The `create`, `read` and `update` methods return two values `(map[string]interface{}, error)`; the first one is the `status`.

The `delete` interface returns two values `(bool, error)`. If there is no error, it returns `(true, nil)`; otherwise it will return `(false, error)`.

If there is no error, it returns `(true, nil)`; otherwise it will return `(false, error)`.

## 3 How do plugins work?

DevStream is using [go plugin](https://pkg.go.dev/plugin) to implement custom plugins.

When you execute a command that calls any of the interfaces (`Create`, `Read`, `Update`, `Delete`), DevStream's `pluginengine` calls the [`plugin.Lookup("DevStreamPlugin")` function](https:// github.com/devstream-io/devstream/blob/38307894bbc08f691b2c5015366d9e45cc87970c/internal/pkg/pluginengine/plugin_helper.go#L28) to load the plugin, get the variable `DevStreamPlugin` that implements the `DevStreamPlugin` interface, and then you can call the corresponding plugin interface to implement the logic. So we do not recommend you to modify `/cmd/plugin/YOUR-PLUGIN-NAME/main.go` file directly, because the file is automatically generated according to the interface definition.

Note: `main()` in the `/cmd/plugin/YOUR-PLUGIN-NAME/main.go` file will not be executed, it is only used to avoid the golangci-lint error.
