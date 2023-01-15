package variable

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "github.com/at0x0ft/cod2e2/yaml_evaluator/node"
)

func visitableFactory(parentPath string, n *yaml.Node) (visitable, error) {
    if node.IsMapping(n) {
        return &mappingNode{*node.CreateMapping(parentPath, n)}, nil
    } else if node.IsSequence(n) {
        return &sequenceNode{*node.CreateSequence(parentPath, n)}, nil
    } else if node.IsScalar(n) {
        return &scalarNode{*node.CreateScalar(parentPath, n)}, nil
    }
    return nil, fmt.Errorf("Undefined Node!\nKind = %v, Tag = %v\n", n.Kind, n.Tag)
}
