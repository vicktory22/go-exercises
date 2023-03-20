package tree

import (
	"errors"
	"fmt"
    "sort"
)

type ChildIdMap map[int]map[int]struct{}

type EmptyVal = struct{}

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func validateRecords(records []Record) error {
    invalidParentError := errors.New("invalid parent record")
    invalidRecordError := errors.New("invalid record")
    duplicateRecordError := errors.New("duplicate record")
    notContinousError := errors.New("not continous")
    noRootNodeError := errors.New("no root node found")

    duplicateMap := map[string]EmptyVal{}
    nestingDepthMap := map[int]EmptyVal{}
    hasRootNode := false
    nestingDepth := 0

    for _, record := range records {
        if record.ID == 0 && record.Parent != 0 {
            return invalidParentError
        }

        if record.ID != 0 && (record.ID == record.Parent) {
            return invalidRecordError
        }

        if record.ID == 0 && record.Parent == 0 {
            hasRootNode = true
        }

        duplicateKey := fmt.Sprintf("%d|%d", record.ID, record.Parent)

        if _, hasDuplicate := duplicateMap[duplicateKey]; hasDuplicate {
            return duplicateRecordError
        }

        duplicateMap[duplicateKey] = EmptyVal{}

        nestingDepthMap[record.Parent] = EmptyVal{}

        if record.Parent > nestingDepth {
            nestingDepth = record.Parent
        }
    }

    if !hasRootNode {
        return noRootNodeError
    }

    if nestingDepth >= len(nestingDepthMap) {
        return notContinousError
    }

    return nil
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

    err := validateRecords(records)

    if err != nil {
        return nil, err
    }

    childMap := make(ChildIdMap)

	for _, record := range records {
        if _, ok := childMap[record.Parent]; !ok {
            childMap[record.Parent] = map[int]EmptyVal{}
        }

        childMap[record.Parent][record.ID] = EmptyVal{}
	}

    delete(childMap[0], 0)

    return &Node{
        ID: 0,
        Children: getChildren(childMap, 0),
    }, nil
}

func getChildren(childIdMap ChildIdMap, parentId int) []*Node {
    childIds, ok := childIdMap[parentId]
    
    if !ok {
        return nil
    }

    children := []*Node{} 

    for childId := range childIds {
        children = append(children, &Node{
            ID: childId,
            Children: getChildren(childIdMap, childId),
        })
    }

    sort.SliceStable(children, func(a, b int) bool {
        return children[a].ID < children[b].ID
    })

    return children
}
