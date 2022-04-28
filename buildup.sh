function build_web() {
    go build -o output/webserver ./main.go
}

function build_tools() {
    go build -o output/exporter ./admin/export.go
}

mkdir output
build_web
build_tools
echo "build Success... Look your ./output/ dir"