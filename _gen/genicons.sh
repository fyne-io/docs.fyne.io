#!/bin/bash

ROOT="`go env GOPATH`/src/fyne.io/fyne/v2"

cd "$(dirname "$0")"
cd ..

function srcFromRes() {
  echo $1 | sed -e "s/IconRes/.svg/g" | sed --expression 's/\([A-Z]\)/-\L\1/g' --expression 's/^-//' \
    | sed -e "s/radiobutton/radio-button-/g" | sed -e "s/search/search-/g" | sed -e "s/color/color-/g" | sed -e "s/content/content-/g" \
    | sed -e "s/checkbox/check-box-/g" | sed -e "s/document/document-/g" | sed -e "s/more/more-/g" | sed -e "s/menu/menu-/g" \
    | sed -e "s/mail/mail-/g" | sed -e "s/media/media-/g" | sed -e "s/fast/fast-/g" | sed -e "s/skip/skip-/g" \
    | sed -e "s/arrow/arrow-/g" | sed -e "s/drop/drop-/g" | sed -e "s/file/file-/" | sed -e "s/folder/folder-/g" \
    | sed -e "s/view/view-/g" | sed -e "s/zoom/zoom-/g" | sed -e "s/visibility/visibility-/g" | sed -e "s/volume/volume-/g" \
    | sed -e "s/broken/broken-/g" | sed -e "s/reply/reply_/g" | sed -e "s/-\./\./g" | sed -e "s/_\./\./g"
}

function stripDTD() {
  cat $1 | sed -e "/dtd\">$/d" | awk -F "svg11.dtd\">" '{ if ($2) { print $2 } else { print $1 }}'
}

# generate icon list

LINES=`grep NewThemedResource\( $ROOT/theme/icons.go | grep -v func | awk -F"IconName" '{print $2}'| awk -F":" '{gsub(/^[ \t]+/, "", $1); gsub(/[ \t]+$/, "", $1); gsub(/^[ \t]+/, "", $2); print $1 "|" $2}' | sed -e "s/NewThemedResource(//g" | sed -e "s/),//g"`
SORTED=`echo $LINES | sed -e "s/ /\n/g" | sort`

OUT="layout/_shortcodes/iconlist.html"
echo '<ul class="theme-icon-list">' > $OUT
for LINE in $SORTED; do
  IFS='|'; parts=($LINE); unset IFS;
  ICON=${parts[0]}
  RES=${parts[1]}
  if [[ -z $RES ]]; then
	  continue 
  fi

  SRC=$(srcFromRes $RES)
  echo "<li class=\"icon-item\" data-icon-theme-method=\"${ICON}Icon\" data-icon-safeName=\"IconName${ICON}\" id=\"IconName${ICON}\">" >> $OUT
  echo "<figure>" >> $OUT
  stripDTD "$ROOT/theme/icons/$SRC" >> $OUT
  echo "<figcaption>${ICON}Icon</figcaption></figure></li>" >> $OUT
done
echo '</ul>' >> $OUT

sed -i.bak 's/<?xml.*?>//' $OUT
