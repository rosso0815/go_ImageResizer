#------------------------------------------------------------------------------

export IM_PATH="/usr/local/opt/imagemagick@6"
export PATH="${IM_PATH}/bin:$PATH"
export LDFLAGS="-L${IM_PATH}/lib"
export CPPFLAGS="-I${IM_PATH}/include"
export PKG_CONFIG_PATH="/${IM_PATH}/lib/pkgconfig"

pkg-config --cflags --libs MagickWand

export CGO_CFLAGS_ALLOW='-Xpreprocessor'

go get -u gopkg.in/gographics/imagick.v3/imagick

GOFLAGS="-count=1" time go test ./mygraphics/

#------------------------------------------------------------------------------

go test -run TestReadDirectory

# mv conv_*jpg to *jpg

find . -type f -name 'conv_*' | while read FILE ; do
    newfile="$(echo ${FILE} |sed -e 's/conv_//')" ;
    echo mv "${FILE}" "${newfile}" ;
done 


