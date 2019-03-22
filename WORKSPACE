rules_jsonnet_commit = "f863822b6268d8e41435bb32ee720242d6633a20"

git_repository(
    name = "io_bazel_rules_jsonnet",
    commit = rules_jsonnet_commit,
    remote = "https://github.com/bazelbuild/rules_jsonnet.git",
)

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

rules_nodejs_version = "0.15.1"  # 0.15.1

rules_nodejs_sha = "7df5904995960be72730e6e8c5e406dc9e55e531fa894c9836300a4ed1ce78d5"

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = rules_nodejs_sha,
    strip_prefix = "rules_nodejs-%s" % rules_nodejs_version,
    url = "https://github.com/bazelbuild/rules_nodejs/archive/%s.tar.gz" % rules_nodejs_version,
)

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dependencies")

rules_nodejs_dependencies()

load("@build_bazel_rules_nodejs//:defs.bzl", "node_repositories")

node_repositories(
    node_version = "8.9.1",
    package_json = [
        "//rules/ecs:package.json",
    ],
    yarn_version = "1.3.2",
)

load("@build_bazel_rules_nodejs//:defs.bzl", "npm_install", "yarn_install")

yarn_install(
    name = "npm_rules_ecs",
    package_json = "//rules/ecs:package.json",
    yarn_lock = "//rules/ecs:yarn.lock",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6776d68ebb897625dead17ae510eac3d5f6342367327875210df44dbe2aeeb19",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.1/rules_go-0.17.1.tar.gz"],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()
