//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcegraph

import "encoding/json"

func unmarshalFacetClassification(rawMsg json.RawMessage) (FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b FacetClassification
	switch m["resultType"] {
	case "FacetError":
		b = &FacetError{}
	case "FacetResult":
		b = &FacetResult{}
	default:
		b = &Facet{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalFacetClassificationArray(rawMsg json.RawMessage) ([]FacetClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]FacetClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalFacetClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}
