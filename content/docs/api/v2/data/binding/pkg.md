---
tags: [api]
title: binding (package)
slug: (package)

aliases:
- /api/v2.0/data/binding
- /api/v2.1/data/binding
- /api/v2.2/data/binding
- /api/v2.3/data/binding
- /api/v2.4/data/binding
- /api/v2.5/data/binding
- /api/v2.6/data/binding
- /api/v2.7/data/binding

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

Package binding provides support for binding data to widgets. All APIs in the binding package are safe to invoke directly from any goroutine.

## Usage

```go
const DataTreeRootID = ""
```
DataTreeRootID const is the value used as ID for the root of any tree binding.

#### types

 * [Bool](bool.html)
 * [BoolList](boollist.html)
 * [BoolTree](booltree.html)
 * [Bytes](bytes.html)
 * [BytesList](byteslist.html)
 * [BytesTree](bytestree.html)
 * [DataItem](dataitem.html)
 * [DataList](datalist.html)
 * [DataListener](datalistener.html)
 * [DataMap](datamap.html)
 * [DataTree](datatree.html)
 * [ExternalBool](externalbool.html)
 * [ExternalBoolList](externalboollist.html)
 * [ExternalBoolTree](externalbooltree.html)
 * [ExternalBytes](externalbytes.html)
 * [ExternalBytesList](externalbyteslist.html)
 * [ExternalBytesTree](externalbytestree.html)
 * [ExternalFloat](externalfloat.html)
 * [ExternalFloatList](externalfloatlist.html)
 * [ExternalFloatTree](externalfloattree.html)
 * [ExternalInt](externalint.html)
 * [ExternalIntList](externalintlist.html)
 * [ExternalIntTree](externalinttree.html)
 * [ExternalItem](externalitem.html)
 * [ExternalList](externallist.html)
 * [ExternalRune](externalrune.html)
 * [ExternalRuneList](externalrunelist.html)
 * [ExternalRuneTree](externalrunetree.html)
 * [ExternalString](externalstring.html)
 * [ExternalStringList](externalstringlist.html)
 * [ExternalStringTree](externalstringtree.html)
 * [ExternalTree](externaltree.html)
 * [ExternalURI](externaluri.html)
 * [ExternalURIList](externalurilist.html)
 * [ExternalURITree](externaluritree.html)
 * [ExternalUntyped](externaluntyped.html)
 * [ExternalUntypedList](externaluntypedlist.html)
 * [ExternalUntypedMap](externaluntypedmap.html)
 * [ExternalUntypedTree](externaluntypedtree.html)
 * [Float](float.html)
 * [FloatList](floatlist.html)
 * [FloatTree](floattree.html)
 * [Int](int.html)
 * [IntList](intlist.html)
 * [IntTree](inttree.html)
 * [Item](item.html)
 * [List](list.html)
 * [Rune](rune.html)
 * [RuneList](runelist.html)
 * [RuneTree](runetree.html)
 * [String](string.html)
 * [StringList](stringlist.html)
 * [StringTree](stringtree.html)
 * [Struct](struct.html)
 * [Tree](tree.html)
 * [URI](uri.html)
 * [URIList](urilist.html)
 * [URITree](uritree.html)
 * [Untyped](untyped.html)
 * [UntypedList](untypedlist.html)
 * [UntypedMap](untypedmap.html)
 * [UntypedTree](untypedtree.html)
