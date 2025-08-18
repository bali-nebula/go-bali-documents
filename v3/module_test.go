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
	var document = doc.ParseSource(source)
	var key = fra.SymbolFromString("$type")
	var parameter = document.GetParameter(key)
	ass.Equal(t, nil, parameter)

	source = `[ ]($type: "foo")`
	document = doc.ParseSource(source)
	parameter = document.GetParameter(key)
	ass.Equal(t, "\"foo\"\n", doc.FormatDocument(parameter))

	source = `[ ]($type: "foo" $hype: /bar)`
	document = doc.ParseSource(source)
	key = fra.SymbolFromString("$hype")
	parameter = document.GetParameter(key)
	ass.Equal(t, "/bar\n", doc.FormatDocument(parameter))
	parameter = doc.ParseSource("'A'")
	document.SetParameter(key, parameter)
	parameter = document.GetParameter(key)
	ass.Equal(t, "'A'\n", doc.FormatDocument(parameter))
	key = fra.SymbolFromString("$new")
	parameter = doc.ParseSource("none")
	document.SetParameter(key, parameter)
	parameter = document.GetParameter(key)
	ass.Equal(t, "none\n", doc.FormatDocument(parameter))
}

func TestAttributeAccess(t *tes.T) {
	var source = `[ ]`
	var document = doc.ParseSource(source)
	var index uti.Index = 1
	var attribute = document.GetAttribute(index)
	ass.Equal(t, nil, attribute)

	source = `[ ]`
	document = doc.ParseSource(source)
	attribute = document.GetAttribute(index)
	ass.Equal(t, nil, attribute)
	attribute = doc.ParseSource("$new")
	index = 0 // Append a new attribute.
	document.SetAttribute(attribute, index)
	index = 1
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$new\n", doc.FormatDocument(attribute))
	document.RemoveAttribute(index)
	attribute = document.GetAttribute(index)
	ass.Equal(t, nil, attribute)

	source = `[
    $alpha
    $beta
    $gamma
]`
	document = doc.ParseSource(source)
	index = 1
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$alpha\n", doc.FormatDocument(attribute))
	index = 2
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$beta\n", doc.FormatDocument(attribute))
	index = 3
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$gamma\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("$delta")
	document.SetAttribute(attribute, index)
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$delta\n", doc.FormatDocument(attribute))
	index = 0
	attribute = doc.ParseSource("$epsilon")
	document.SetAttribute(attribute, index)
	index = 4
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$epsilon\n", doc.FormatDocument(attribute))
	document.RemoveAttribute(index)
	index = -1
	attribute = document.GetAttribute(index)
	ass.Equal(t, "$delta\n", doc.FormatDocument(attribute))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	document = doc.ParseSource(source)
	var key = fra.Symbol("alpha")
	attribute = document.GetAttribute(key)
	ass.Equal(t, "\"1\"\n", doc.FormatDocument(attribute))
	key = fra.Symbol("beta")
	attribute = document.GetAttribute(key)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))
	key = fra.Symbol("gamma")
	attribute = document.GetAttribute(key)
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("\"5\"")
	document.SetAttribute(attribute, key)
	attribute = document.GetAttribute(key)
	ass.Equal(t, "\"5\"\n", doc.FormatDocument(attribute))
	document.RemoveAttribute(key)
	key = fra.Symbol("beta")
	attribute = document.GetAttribute(key)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))

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
	document = doc.ParseSource(source)
	key = fra.Symbol("items")
	index = 2
	attribute = document.GetAttribute(key, index)
	ass.Equal(t, "2\n", doc.FormatDocument(attribute))
	key = fra.Symbol("attributes")
	var key2 = fra.Symbol("gamma")
	attribute = document.GetAttribute(key, key2)
	ass.Equal(t, "\"3\"\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("\"5\"")
	document.SetAttribute(attribute, key, key2)
	attribute = document.GetAttribute(key, key2)
	ass.Equal(t, "\"5\"\n", doc.FormatDocument(attribute))
	document.RemoveAttribute(key, key2)
	key2 = fra.Symbol("beta")
	attribute = document.GetAttribute(key, key2)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))

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
	document = doc.ParseSource(source)
	index = 1
	var index2 uti.Index = 2
	attribute = document.GetAttribute(index, index2)
	ass.Equal(t, "2\n", doc.FormatDocument(attribute))
	index = 2
	key2 = fra.Symbol("beta")
	attribute = document.GetAttribute(index, key2)
	ass.Equal(t, "\"2\"\n", doc.FormatDocument(attribute))
	index = 1
	index2 = 3
	var key3 = fra.AngleFromString("~tau")
	attribute = document.GetAttribute(index, index2, key3)
	ass.Equal(t, "6.28\n", doc.FormatDocument(attribute))
	attribute = doc.ParseSource("~τ")
	document.SetAttribute(attribute, index, index2, key3)
	attribute = document.GetAttribute(index, index2, key3)
	ass.Equal(t, "~τ\n", doc.FormatDocument(attribute))
	index = 2
	key2 = fra.Symbol("alpha")
	var index3 uti.Index = -1
	document.RemoveAttribute(index, key2, index3)
	attribute = document.GetAttribute(index, key2, index3)
	ass.Equal(t, "'b'\n", doc.FormatDocument(attribute))
	index = 3
	attribute = document.GetAttribute(index)
	ass.Equal(t, nil, attribute)
	index = 2
	key2 = fra.Symbol("delta")
	attribute = document.GetAttribute(index, key2)
	ass.Equal(t, nil, attribute)
}
