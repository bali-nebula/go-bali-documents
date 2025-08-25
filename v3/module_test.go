/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	doc "github.com/bali-nebula/go-bali-documents/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	var filenames = uti.ReadDirectory(testDirectory)
	for _, filename := range filenames {
		if sts.HasSuffix(filename, ".bali") {
			filename = testDirectory + filename
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var document = doc.ParseSource(source)
			var formatted = doc.FormatDocument(document)
			ass.Equal(t, source, formatted)
		}
	}
}

func TestParameterAccess(t *tes.T) {
	var source = `[ ]`
	var component = doc.ParseSource(source)
	var key = fra.SymbolFromString("$type")
	var parameter = component.GetParameter(key)
	ass.Equal(t, nil, parameter)

	source = `[ ]($type: "foo")`
	component = doc.ParseSource(source)
	parameter = component.GetParameter(key)
	ass.Equal(t, "\"foo\"", doc.FormatComponent(parameter))

	source = `[ ]($type: "foo" $hype: /bar $skype: none)`
	component = doc.ParseSource(source)
	key = fra.SymbolFromString("$type")
	parameter = component.GetParameter(key)
	ass.Equal(t, "\"foo\"", doc.FormatComponent(parameter))
	key = fra.SymbolFromString("$hype")
	parameter = component.GetParameter(key)
	ass.Equal(t, "/bar", doc.FormatComponent(parameter))
	key = fra.SymbolFromString("$skype")
	parameter = component.GetParameter(key)
	ass.Equal(t, "none", doc.FormatComponent(parameter))
}

func TestObjectAccess(t *tes.T) {
	var source = `[ ]`
	var component = doc.ParseSource(source)
	var index uti.Index = 1
	var object = component.GetObject(index)
	ass.Equal(t, nil, object)

	var component2 = doc.ParseSource("$new")
	index = 0 // Append a new object.
	component.SetObject(component2, index)
	index = 1
	object = component.GetObject(index)
	ass.Equal(t, "$new", doc.FormatComponent(object))
	component.RemoveObject(index)
	object = component.GetObject(index)
	ass.Equal(t, nil, object)

	source = `[
    $alpha
    $beta
    $gamma
]`
	component = doc.ParseSource(source)
	index = 1
	object = component.GetObject(index)
	ass.Equal(t, "$alpha", doc.FormatComponent(object))
	index = 2
	object = component.GetObject(index)
	ass.Equal(t, "$beta", doc.FormatComponent(object))
	index = 3
	object = component.GetObject(index)
	ass.Equal(t, "$gamma", doc.FormatComponent(object))
	component2 = doc.ParseSource("$delta")
	component.SetObject(component2, index)
	object = component.GetObject(index)
	ass.Equal(t, "$delta", doc.FormatComponent(object))
	index = 0
	component2 = doc.ParseSource("$epsilon")
	component.SetObject(component2, index)
	index = 4
	object = component.GetObject(index)
	ass.Equal(t, "$epsilon", doc.FormatComponent(object))
	component.RemoveObject(index)
	index = -1
	object = component.GetObject(index)
	ass.Equal(t, "$delta", doc.FormatComponent(object))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	component = doc.ParseSource(source)
	var key = fra.Symbol("alpha")
	object = component.GetObject(key)
	ass.Equal(t, "\"1\"", doc.FormatComponent(object))
	key = fra.Symbol("beta")
	object = component.GetObject(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(object))
	key = fra.Symbol("gamma")
	object = component.GetObject(key)
	ass.Equal(t, "\"3\"", doc.FormatComponent(object))
	component2 = doc.ParseSource("\"5\"")
	component.SetObject(component2, key)
	object = component.GetObject(key)
	ass.Equal(t, "\"5\"", doc.FormatComponent(object))
	component.RemoveObject(key)
	key = fra.Symbol("beta")
	object = component.GetObject(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(object))

	source = `[
    $items: [
        1
        2
        3
    ]
    $attributes: [
        $alpha: "1"
        $beta: "2"
        $gamma: "3"
    ]
]`
	component = doc.ParseSource(source)
	key = fra.Symbol("items")
	index = 2
	object = component.GetObject(key, index)
	ass.Equal(t, "2", doc.FormatComponent(object))
	key = fra.Symbol("attributes")
	var key2 = fra.Symbol("gamma")
	object = component.GetObject(key, key2)
	ass.Equal(t, "\"3\"", doc.FormatComponent(object))
	component2 = doc.ParseSource("\"5\"")
	component.SetObject(component2, key, key2)
	object = component.GetObject(key, key2)
	ass.Equal(t, "\"5\"", doc.FormatComponent(object))
	component.RemoveObject(key, key2)
	key2 = fra.Symbol("beta")
	object = component.GetObject(key, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(object))

	source = `[
    [
        1
        2
        [
            ~pi: 3.14
            ~tau: 6.28
        ]
    ]
    [
        $alpha: [
            'a'
            'b'
            'c'
        ]
        $beta: "2"
        $gamma: "3"
    ]
]`
	component = doc.ParseSource(source)
	index = 1
	var index2 uti.Index = 2
	object = component.GetObject(index, index2)
	ass.Equal(t, "2", doc.FormatComponent(object))
	index = 2
	key2 = fra.Symbol("beta")
	object = component.GetObject(index, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(object))
	index = 1
	index2 = 3
	var key3 = fra.AngleFromString("~tau")
	object = component.GetObject(index, index2, key3)
	ass.Equal(t, "6.28", doc.FormatComponent(object))
	component2 = doc.ParseSource("~τ")
	component.SetObject(component2, index, index2, key3)
	object = component.GetObject(index, index2, key3)
	ass.Equal(t, "~τ", doc.FormatComponent(object))
	index = 2
	key2 = fra.Symbol("alpha")
	var index3 uti.Index = -1
	component.RemoveObject(index, key2, index3)
	object = component.GetObject(index, key2, index3)
	ass.Equal(t, "'b'", doc.FormatComponent(object))
	index = 3
	object = component.GetObject(index)
	ass.Equal(t, nil, object)
	index = 2
	key2 = fra.Symbol("delta")
	object = component.GetObject(index, key2)
	ass.Equal(t, nil, object)
}
