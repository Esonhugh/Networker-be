mkdir output
mkdir output/admintool

function build_main_web () {
    go build -o output/webserver ./main.go
}

function build_export_tool () {
    go build -o output/admintool/exporter ./admin/export.go
}

function build_clean_tool () {
    go build -o output/admintool/cleaner ./admin/clean.go
}

function build_password_tool () {
    go build -o output/admintool/passchanger ./admin/passRecovery.go
}

function build_tools () {
    build_clean_tool
    build_export_tool
    build_password_tool
}

function  build() {
    build_main_web
    build_tools
}

build
cp ./application.yaml ./output/
echo "build Success... Look your ./output/ dir"