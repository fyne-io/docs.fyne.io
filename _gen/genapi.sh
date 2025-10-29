#!/bin/bash

ROOT="`go env GOPATH`/src/fyne.io/fyne/v2"
bash -c "cd $ROOT; git checkout master; git pull"
VERSION="v2.7"

cd "$(dirname "$0")"
cd ..
go install github.com/andydotxyz/godocdown/godocdown@latest

# generate API docs

DIRS=`find $ROOT -type d | grep -v .git | grep -v vendor | grep -v internal | grep -v testdata | grep -v cmd | grep -v tools`
PREFIX="content/docs/api/v2"
mkdir $PREFIX 2>&1 > /dev/null

godocdown -template="_gen/api.md" -outputDir "$PREFIX/" $ROOT 2>&1 | grep -v "Could not find package"
for DIR in $DIRS; do
  PKG=`echo $DIR | cut -c$((${#ROOT}+2))-`

  if [[ ! -z "$PKG" ]]; then
    mkdir -p `dirname "api/$PKG"` 2>&1 > /dev/null
  fi
  mkdir -p "$PREFIX/$PKG" 2>&1 > /dev/null
 
  godocdown -template="_gen/api.md" -outputDir "$PREFIX/$PKG/" $DIR 2>&1 | grep -v "Could not find package"

  if [[ "x$PKG" = "x" ]]; then 
    mkdir $PREFIX/fyne
	echo -e "---\ntitle: Fyne API\nslug: fyne\n---\n" > $PREFIX/fyne/_index.md
  fi

  for FILE in $PREFIX/$PKG/*.md; do
    NAME=`basename $FILE | sed 's/\.md//'`
	PKGPATH="$PKG\/$NAME"
	OUT=$NAME".md"
	
	if [[ "$NAME" = "index" ]]; then 
		mv $PREFIX/$PKG/index.md $PREFIX/$PKG/pkg.md
		OUT="pkg.md"
		if [[ "x$PKG" = "x" ]]; then
			PARENT="fyne"
		else
			PARENT=`basename "$PKG"`
		fi
		NAME="(package)"
		PKGPATH="$PKG"
		
		echo -e "---\ntitle: $PKG\n---\n" > $PREFIX/$PKG/_index.md
	    sed -i.bak "s|title: $PARENT|title: $PARENT (package)|" "$PREFIX/$PKG/$OUT"
	fi
	
	PKGNAME=$PKG
	if [[ "x$PKG" = "x" ]]; then
		PKGNAME="fyne"
		mv $PREFIX/$OUT $PREFIX/fyne/$OUT
	fi

    declare -a vers=("v2.0" "v2.1" "v2.2" "v2.3" "v2.4" "v2.5" "v2.6" "v2.7")
	
    ALIASLIST="\n- \/api\/$PKGPATH\n- \/api\/$PKGPATH.html"
    for i in "${vers[@]}"
    do
      ALIASLIST=$ALIASLIST"\n- \/api\/$i\/$PKGPATH\n- \/api\/$i\/$PKGPATH.html"
    done
	
    if [[ "$OUT" != '*.md' ]]; then
      sed -i.bak "s|slug:|slug: $NAME\n\naliases:$ALIASLIST|" "$PREFIX/$PKGNAME/$OUT"
      sed -i.bak "s|^# .*||" "$PREFIX/$PKGNAME/$OUT"
    fi
  done
done

rm content/docs/api/v2/_index.md
find content/docs/api/ -name \*.bak -exec rm {} \;
