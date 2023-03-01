# Repository Rules not being re-evaluated

Reproduction for (#issue)

Tested with bazel versions:
- `6.0.0`
- `7.0.0-pre.20230215.2`

## Current Behavior:

```console
➜ echo "contents" > /tmp/file.txt

➜ bazel build @test_read_file//...
INFO: Invocation ID: c6ab4a93-90bf-49df-a077-54bbbd56aef5
DEBUG: /Volumes/code/github.com/blorente/standalone-repros/test_repo_rule.bzl:3:8: The random file contains 'contents'
INFO: Analyzed target @test_read_file//:empty (1 packages loaded, 1 target configured).
INFO: Found 1 target...
Target @test_read_file//:empty up-to-date (nothing to build)
INFO: Elapsed time: 0.153s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action

➜ cat bazel-standalone-repros/external/test_read_file/BUILD.bazel

# contents of random file are contents
filegroup(name = 'empty')

➜ echo "other contents" > /tmp/file.txt

➜ bazel build @test_read_file//...
INFO: Invocation ID: a47e85b7-a69f-4773-b239-0e1db3967f0b
INFO: Analyzed target @test_read_file//:empty (1 packages loaded, 1 target configured).
INFO: Found 1 target...
Target @test_read_file//:empty up-to-date (nothing to build)
INFO: Elapsed time: 0.089s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action

➜ cat bazel-standalone-repros/external/test_read_file/BUILD.bazel

# contents of random file are contents
filegroup(name = 'empty')
```

## Expected Behavior:

As per [the `repository_rule` documentation](https://bazel.build/versions/6.0.0/rules/lib/globals#repository_rule(implementation,%20attrs,%20local,%20environ,%20configure,%20remotable,%20doc)), setting either `local = True` or `config = True` should cause the rule to re-run on every invocation.

Here we set both, and it doesn't seem to run.

If I run `bazel shutdown`, the rule gets reevaluated correctly.
