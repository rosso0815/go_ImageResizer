#------------------------------------------------------------------------------

export IM_PATH="/usr/local/im"
export PATH="${IM_PATH}/bin:$PATH"
export LDFLAGS="-L${IM_PATH}/lib"
export CPPFLAGS="-I${IM_PATH}/include"
export PKG_CONFIG_PATH="/${IM_PATH}/lib/pkgconfig"

pkg-config --cflags --libs MagickWand

export CGO_CFLAGS_ALLOW='-Xpreprocessor'

go clean --modcache
go clean -i -r -cache -testcache -modcache

GODEBUG=gocacheverify=1 go get -u gopkg.in/gographics/imagick.v3/imagick

GOFLAGS="-count=1" time go test ./mygraphics/

#------------------------------------------------------------------------------

go test -run TestReadDirectory

# mv conv_*jpg to *jpg

find . -type f -name 'conv_*' | while read FILE ; do
    newfile="$(echo ${FILE} |sed -e 's/conv_//')" ;
    echo mv "${FILE}" "${newfile}" ;
done 


CPU=1 TIME=15sec

# Docker

## Build

```
# lets build
docker build .

```