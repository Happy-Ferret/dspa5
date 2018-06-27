#!/bin/sh

#SCRIPT_DIR=$(dirname $0)
SCRIPT_DIR=~/EnglishHTSVoices/

#   http://homepages.inf.ed.ac.uk/jyamagis/software/page54/page54.html

$SCRIPT_DIR/build/bin/flite_hts_engine \
     -td $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/tree-dur.inf -tf $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/tree-lf0.inf -tm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/tree-mgc.inf \
     -md $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/dur.pdf         -mf $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/lf0.pdf       -mm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/mgc.pdf \
     -df $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/lf0.win1        -df $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/lf0.win2      -df $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/lf0.win3 \
     -dm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/mgc.win1        -dm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/mgc.win2      -dm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/mgc.win3 \
     -cf $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/gv-lf0.pdf      -cm $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/gv-mgc.pdf    -ef $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/tree-gv-lf0.inf \
     -em $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/tree-gv-mgc.inf -k  $SCRIPT_DIR/build/hts_voice_cmu_us_arctic_slt-1.03/gv-switch.inf -o  $1 \
     /dev/stdin

#$SCRIPT_DIR/build/bin/flite_hts_engine \
#     -td $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/tree-dur.inf -tf $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/tree-lf0.inf -tm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/tree-mgc.inf \
#     -md $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/dur.pdf         -mf $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/lf0.pdf       -mm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/mgc.pdf \
#     -df $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/lf0.win1        -df $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/lf0.win2      -df $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/lf0.win3 \
#     -dm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/mgc.win1        -dm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/mgc.win2      -dm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/mgc.win3 \
#     -cf $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/gv-lf0.pdf      -cm $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/gv-mgc.pdf    -ef $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/tree-gv-lf0.inf \
#     -em $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/tree-gv-mgc.inf -k  $SCRIPT_DIR/build/hts_voice_cstr_uk_female-1.0/gv-switch.inf -o $1 \
#     -s  48000.0\
#     -p  240.0\
#     -a  0.55\
#     -g  0.0\
#     -b  0.2\
#     -u  0.5\
#     -jm 0.8 \
#     /dev/stdin

