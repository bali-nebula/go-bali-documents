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

func TestMemberAccess(t *tes.T) {
	var source = `[ ]`
	var component = doc.ParseSource(source)
	var index uti.Index = 1
	var member = component.GetMember(index)
	ass.Equal(t, nil, member)

	var component2 = doc.ParseSource("$new")
	index = 0 // Append a new member.
	component.SetMember(component2, index)
	index = 1
	member = component.GetMember(index)
	ass.Equal(t, "$new", doc.FormatComponent(member))
	component.RemoveMember(index)
	member = component.GetMember(index)
	ass.Equal(t, nil, member)

	source = `[
    $alpha
    $beta
    $gamma
]`
	component = doc.ParseSource(source)
	index = 1
	member = component.GetMember(index)
	ass.Equal(t, "$alpha", doc.FormatComponent(member))
	index = 2
	member = component.GetMember(index)
	ass.Equal(t, "$beta", doc.FormatComponent(member))
	index = 3
	member = component.GetMember(index)
	ass.Equal(t, "$gamma", doc.FormatComponent(member))
	component2 = doc.ParseSource("$delta")
	component.SetMember(component2, index)
	member = component.GetMember(index)
	ass.Equal(t, "$delta", doc.FormatComponent(member))
	index = 0
	component2 = doc.ParseSource("$epsilon")
	component.SetMember(component2, index)
	index = 4
	member = component.GetMember(index)
	ass.Equal(t, "$epsilon", doc.FormatComponent(member))
	component.RemoveMember(index)
	index = -1
	member = component.GetMember(index)
	ass.Equal(t, "$delta", doc.FormatComponent(member))

	source = `[
    $alpha: "1"
    $beta: "2"
    $gamma: "3"
]`
	component = doc.ParseSource(source)
	var key = fra.Symbol("alpha")
	member = component.GetMember(key)
	ass.Equal(t, "\"1\"", doc.FormatComponent(member))
	key = fra.Symbol("beta")
	member = component.GetMember(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(member))
	key = fra.Symbol("gamma")
	member = component.GetMember(key)
	ass.Equal(t, "\"3\"", doc.FormatComponent(member))
	component2 = doc.ParseSource("\"5\"")
	component.SetMember(component2, key)
	member = component.GetMember(key)
	ass.Equal(t, "\"5\"", doc.FormatComponent(member))
	component.RemoveMember(key)
	key = fra.Symbol("beta")
	member = component.GetMember(key)
	ass.Equal(t, "\"2\"", doc.FormatComponent(member))

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
	member = component.GetMember(key, index)
	ass.Equal(t, "2", doc.FormatComponent(member))
	key = fra.Symbol("attributes")
	var key2 = fra.Symbol("gamma")
	member = component.GetMember(key, key2)
	ass.Equal(t, "\"3\"", doc.FormatComponent(member))
	component2 = doc.ParseSource("\"5\"")
	component.SetMember(component2, key, key2)
	member = component.GetMember(key, key2)
	ass.Equal(t, "\"5\"", doc.FormatComponent(member))
	component.RemoveMember(key, key2)
	key2 = fra.Symbol("beta")
	member = component.GetMember(key, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(member))

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
	member = component.GetMember(index, index2)
	ass.Equal(t, "2", doc.FormatComponent(member))
	index = 2
	key2 = fra.Symbol("beta")
	member = component.GetMember(index, key2)
	ass.Equal(t, "\"2\"", doc.FormatComponent(member))
	index = 1
	index2 = 3
	var key3 = fra.AngleFromString("~tau")
	member = component.GetMember(index, index2, key3)
	ass.Equal(t, "6.28", doc.FormatComponent(member))
	component2 = doc.ParseSource("~τ")
	component.SetMember(component2, index, index2, key3)
	member = component.GetMember(index, index2, key3)
	ass.Equal(t, "~τ", doc.FormatComponent(member))
	index = 2
	key2 = fra.Symbol("alpha")
	var index3 uti.Index = -1
	component.RemoveMember(index, key2, index3)
	member = component.GetMember(index, key2, index3)
	ass.Equal(t, "'b'", doc.FormatComponent(member))
	index = 3
	member = component.GetMember(index)
	ass.Equal(t, nil, member)
	index = 2
	key2 = fra.Symbol("delta")
	member = component.GetMember(index, key2)
	ass.Equal(t, nil, member)
}
