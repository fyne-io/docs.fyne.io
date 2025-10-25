#!/bin/bash

cd "$(dirname "$0")"
cd ..

PREFIX="content/docs"

legacy_files(){
  for FILE in $1/*; do
    FULL=`echo $FILE | sed "s|$PREFIX||"`
    NAME=`basename $FILE | sed 's/\.md//'`
    OUT=$NAME".md"

    if [ -d "$FILE" ]; then
      legacy_files "$FILE"
    else
      if [[ $NAME != "_index" ]]; then
        REDIR=`echo $FULL | sed "s/\.md/\.html/"`
        sed -i.bak "s|aliases:|aliases:\n- $REDIR|" "$FILE"
      fi
    fi
  done
}

legacy_files "$PREFIX"

find content/docs/ -name \*.bak -exec rm {} \;
