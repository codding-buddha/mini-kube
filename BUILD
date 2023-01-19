load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "mini-kube_lib",
    srcs = ["main.go"],
    importpath = "github.com/codding-buddha/mini-kube",
    visibility = ["//visibility:private"],
    deps = [
        "//manager",
        "//task",
        "//worker",
        "@com_github_golang_collections_collections//queue:go_default_library",
        "@com_github_google_uuid//:go_default_library",
    ],
)

go_binary(
    name = "mini-kube",
    embed = [":mini-kube_lib"],
    visibility = ["//visibility:public"],
)
