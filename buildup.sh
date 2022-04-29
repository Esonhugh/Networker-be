function build_main_web() {
    go build -o output/webserver ./main.go
}

function build_export_tools() {
    go build -o output/exporter ./admin/export.go
}

function build_clean_tool() {
    go build -o output/cleaner ./admin/clean.go
}

function build_tools() {
    build_clean_tool
    build_export_tool
}

function  build() {
    build_main_web
    build_tools
}

mkdir output
build
cp ./application.yaml ./output/
echo "build Success... Look your ./output/ dir"