# Releasing and change management

## Minor version

1. Create a new minor version changeset and merge it to main `pnpm changeset`
2. Merge the automated changesets PR `Version Packages for main` this will tag and release a new version

## Patch version

1. Create a patch changeset and merge it to main `pnpm changeset`
2. If you are ready to release this change, Merge the automated changesets PR `Version Packages for main` this will tag and release a new version

## Backporting changes to older minor versions

1. Check for an existing release branch e.g `release/v1.*` for the minor version you are back porting to
2. If the branch doesn't exist, checkout the commit associated with the minor release you want to backport to
3. Run `./create-minor-release.sh` and when prompted set the release branch version e.g `v1.1` or `v1.2` excluding the patch version
4. Push the branch
5. You can now push your changes to this release branch, you may choose to cherrypick the patch from main
6. Create a patch changeset and merge it to the release branch `pnpm changeset`
7. If you are ready to release this change, Merge the automated changesets PR matching your release branch `Version Packages for release/v1.*` this will tag and release a new version

## Adding protobuf validations

We use protovalidate to add proto level validations to our API. There are examples already in the proto files. But you can also use the [docs](https://github.com/bufbuild/protovalidate/blob/main/docs/cel.md)

## Generating code

We use a combination of buf for proto generation and a post processing step written in go which extracts variables from the generated files into a collection for use later.

Use `make generate` to run the code generation steps
