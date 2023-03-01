def _test_repo_rule_impl(rctx):
  test_random_file = rctx.execute(["cat", rctx.attr.file_to_read])
  print("The random file contains '{}'".format(test_random_file.stdout.strip()))
  rctx.file(
    "BUILD.bazel",
    """
# contents of random file are {}
filegroup(name = 'empty')
""".format(test_random_file.stdout.strip())
  )

test_repo_rule = repository_rule(
  implementation = _test_repo_rule_impl,
  # As per this documentation:
  # https://bazel.build/versions/6.0.0/rules/lib/globals#repository_rule(implementation,%20attrs,%20local,%20environ,%20configure,%20remotable,%20doc)
  # I would expect these two toggles to mean that the rule gets re-evaluated.
  configure = True,
  local = True,
  attrs = {
    "file_to_read": attr.string(default="/tmp/file.txt"),
  },
)
