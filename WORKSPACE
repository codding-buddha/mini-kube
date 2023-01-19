load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "56d8c5a5c91e1af73eca71a6fab2ced959b67c86d12ba37feedb0a2dfea441a6",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.37.0/rules_go-v0.37.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "448e37e0dbf61d6fa8f00aaa12d191745e14f07c31cabfa731f0c8e8a4f41b97",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_protobuf",
    commit = "b71cae3d2f698d0990f7c96cc63e6f5fceed1a81",
    remote = "https://github.com/protocolbuffers/protobuf",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
    version = "v0.0.0-20210617225240-d185dfc1b5a1",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:n56/Zwd5o6whRC5PMGretI4IdRLlmBXYNjScPaBgsbY=",
    version = "v1.1.18",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:Q50tZOPR6T/hjNsyc9g8/syEs6bk8XXApsHjKukMl68=",
    version = "v2.8.1+incompatible",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:UTLdBmHk3bEY+w8qeO5KttOhy6OmXWsl/FEet9Uswog=",
    version = "v20.10.21+incompatible",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_go_chi_chi_v5",
    importpath = "github.com/go-chi/chi/v5",
    sum = "h1:lD+NLqFcAi1ovnVZpsnObHGW4xb4J8lNmoYVfECH1Y0=",
    version = "v5.0.8",
)

go_repository(
    name = "com_github_go_ole_go_ole",
    importpath = "github.com/go-ole/go-ole",
    sum = "h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=",
    version = "v1.2.6",
)

# go_repository(
#     name = "com_github_gogo_protobuf",
#     importpath = "github.com/gogo/protobuf",
#     build_file_proto_mode = "disable_global",
#     sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
#     patches = [
#         "@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch",
#     ],
#     patch_args = ["-p1"],
#     version = "v1.3.2",
# )

go_repository(
    name = "com_github_golang_collections_collections",
    importpath = "github.com/golang-collections/collections",
    sum = "h1:zN2lZNZRflqFyxVaTIU61KNKQ9C0055u9CAfpmqUvo4=",
    version = "v0.0.0-20130729185459-604e922904d3",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
    version = "v0.5.9",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:t6JiXgmwXMjEs8VusXIJk2BXHsn+wx8BZdTaoZ5fu7I=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_lufia_plan9stats",
    importpath = "github.com/lufia/plan9stats",
    sum = "h1:6E+4a0GO5zZEnZ81pIr0yLvtUWk2if982qA3F3QD6H4=",
    version = "v0.0.0-20211012122336-39d0f177ccd0",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:slsWYD/zyx7lCXoZVlvQrj0hPTM1HI4+v1sIda2yDvg=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:HfkjXDfhgVaN5rmueG8cL8KKeFNecRCXFhaJ2qZ5SKA=",
    version = "v0.0.0-20221205130635-1aeaba878587",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:9yCKha/T5XdGtO0q9Q9a6T5NUCsTn/DrBg0D7ufOcFM=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_power_devops_perfstat",
    importpath = "github.com/power-devops/perfstat",
    sum = "h1:ncq/mPwQF4JjgDlrVEn3C11VoGHZN7m8qihwgMEtzYw=",
    version = "v0.0.0-20210106213030-5aafc221ea8c",
)

go_repository(
    name = "com_github_shirou_gopsutil_v3",
    importpath = "github.com/shirou/gopsutil/v3",
    sum = "h1:kxsPKS+Eeo+VnEQ2XCaGJepeP6KY53QoRTETx3+1ndM=",
    version = "v3.22.11",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:w7B6lhMri9wdJUVmEZPGGhZzrYTPvgJArz7wNPgYKsk=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_tklauser_go_sysconf",
    importpath = "github.com/tklauser/go-sysconf",
    sum = "h1:89WgdJhk5SNwJfu+GKyYveZ4IaJ7xAkecBo+KdJV0CM=",
    version = "v0.3.11",
)

go_repository(
    name = "com_github_tklauser_numcpus",
    importpath = "github.com/tklauser/numcpus",
    sum = "h1:kebhY2Qt+3U6RNK7UqpYNA+tJ23IBEGKkB7JQBfDYms=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "com_github_yusufpapurcu_wmi",
    importpath = "github.com/yusufpapurcu/wmi",
    sum = "h1:KBNDSne4vP5mbSWnJbO+51IMOXJB67QiYCSBrubbPRg=",
    version = "v1.2.2",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:7I4JAnoQBe7ZtJcBaYHi5UtiO8tQHbUSXxL+pnGRANg=",
    version = "v0.0.0-20210921155107-089bfa567519",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:6zppjxzCulZykYSLyVDYbneBfbaBIQPYMevg0bEwv2s=",
    version = "v0.6.0-dev.0.20220419223038-86c51ed26bb4",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:Q5QPcMlvfxFTAPV0+07Xz/MpK9NTXu2VDUuy0FeMfaU=",
    version = "v0.4.0",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:uVc8UZUe6tr40fFVnUP5Oj+veunVezqYl9z7DYw9xzw=",
    version = "v0.0.0-20220722155255-886fb9371eb4",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:w8ZOecv6NaNa/zC8944JTU3vz4u6Lagfk4RPQxv92NQ=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:qoo4akIqOcDME5bhc/NgxUdovd6BSS2uMsVjB56q1xI=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:OLmvp0KP+FVG99Ct/qFiL/Fhk4zp4QQnZ7b2U+5piUM=",
    version = "v0.5.0",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
    version = "v0.1.12",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

go_repository(
    name = "tools_gotest_v3",
    importpath = "gotest.tools/v3",
    sum = "h1:ZazjZUfuVeZGLAmlKKuyv3IKP5orXcwtOwDQH6YVr6o=",
    version = "v3.4.0",
)

go_rules_dependencies()

go_register_toolchains(version = "1.19.3")

gazelle_dependencies()
