where can I found this .so file
cp ../openstream-plugin-githubactions/githubactions_0.0.1.so plugins/

how does 'githubactions_0.0.1.so' implement the interface 'OpenStreamPlugin'

p.Install(&tool.Options)

when does the program call the function 'installCMDFunc'

go test -v ./...

```go
type RepositoryContentFileOptions struct {
	Message   *string       `json:"message,omitempty"` // ?? *string, omitempty
	Content   []byte        `json:"content,omitempty"` // unencoded
	SHA       *string       `json:"sha,omitempty"`
	Branch    *string       `json:"branch,omitempty"`
	Author    *CommitAuthor `json:"author,omitempty"`
	Committer *CommitAuthor `json:"committer,omitempty"`
}
```

```go
import (
    . "github.com/devstream-io/devstream/internal/pkg/plugin/common/helm"
    _ "github.com/devstream-io/devstream/internal/pkg/plugin/common/helm"
)
```