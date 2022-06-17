# To Boostrap and Run a Node

```shell
./scripts/start-test-node.sh
```

# Blockchain Engineer Challenge

This challenge consists of adding a new `commenting` functionality to `x/blog`, a demo Cosmos SDK module that can be used to run a decentralized blog.

## The Challenge

At its current state in this repo, `x/blog` allows users to create blog posts and to list all existing blog posts. The challenge consists of adding comments to `x/blog`: we want users to be able to comment on those blog posts.

Concretely, we would like to see the following components added:

- The addition of a new `Msg` service method called `CreateComment`. This method's input should take:
  - `post_slug` as string,
  - `author` as string (Bech32-encoded address),
  - `body` as string.
- New comments should be persisted to the blockchain state. A comment cannot be inserted into state if its corresponding `Post` does not exist in state.
- The addition of a new `AllComments` gRPC query method. This method's input should take a `post_slug` (a string), and returns all the comments on a post.
- The addition of two CLI subcommands, `tx create-comment` and `query list-comments`, which call the `CreateComment` and `AllComments` service methods under the hood.

## Requirements:

- Go v1.17+: https://golang.org/
- Docker: https://www.docker.com/

## Getting Started with `x/blog`

A good place to start learning about `x/blog` is the [`./proto/blog/v1`](./proto/blog/v1) folder. It contains 3 files:

- `types.proto`: defines the shared messages that may be used in other files. We define inside it a `Post`, which represents a blog post. The `slug` field is a string that represents a human-readable identifier for each post (see [definition](https://yoast.com/slug/)), and the `author` field is a bech32-encoded address.
- `query.proto`: defines the `Query` service, or how to query the state. It contains for now a single method that allows to query all posts.
- `tx.proto`: defines the `Msg` service, or how to handle state transitions. It contains for now a single method that allows to create a new `Post`.

```bash
x/blog
    |- client/
        |- cli/             # CLI commands related to x/blog
    |- module/              # Implementation of AppModule, so that x/blog can be wired up to the app.
    |- server/
        |- msg_server.go    # `Msg` server: handling state transitions. Its defines the implementation of the `Msg`
                            # service defined in `./proto/blog/v1/tx.proto`.
        |- query_server.go  # `Query` server: querying state. Its defines the implementation of the `Query` service
                            # defined in `./proto/blog/v1/tx.proto`.
    |- codec.go
    |- keys.go              # Prefix keys used in the store.
    |- requests.go
    |- *.pb.go              # Files generated by gogoproto
```

If you modify the `./proto/blog/v1/*.proto` files and wish to re-generate the associated `*.pb.go` files, run the command:

```bash
make proto-gen
```

> Note: You can safely ignore the warning `No HttpRule found for method` for this command.

## Build and Run the Node (Optional)

If you would like to build the node to be able to test it, just run:

```bash
make build
```

A new binary will be created under `./build/regen`.

To run the node, you can refer to the documentation [here](https://docs.cosmos.network/master/run-node/).

## Resources

- Cosmos SDK Documentation: https://docs.cosmos.network/master/.
- Having a look at how existing modules are implemented can help. Here are the modules we're currently maintaining inside `regen-ledger`: https://github.com/regen-network/regen-ledger/tree/master/x.
- Ask us questions! Shoot us an email, or talk to us on [Discord](https://discord.gg/stujhkkhvk), if you need help.

## Submission

For submission, please send us an email with a link to your project, ideally as a Github repo. Oh, and don't forget to add a nice README.md (or edit this one) so we know how to build & run it, and what changes you made :)

## Additional Information

### Expected Timeline

The entire task should only take 2-3 hours, but you’re free to take it as far as you like. We don't expect you to come up with a perfect solution, nor do we want to exploit the idea of a take home task by requiring you to build the module with full functionality. Whenever you run out of time, add a comment describing what you would have done if you had more time.

### Documentation

With this take home task we would like to understand how you tackle tasks like the above. In order for us to easier understand what your reasoning for certain decisions is, please make sure to write good code comments and documentation and maintain a proper Git history with commit messages explaining each step.

Feel free to include a list of ideas on how to improve your final project under the assumption that you have unlimited time and resources to spend on it.
