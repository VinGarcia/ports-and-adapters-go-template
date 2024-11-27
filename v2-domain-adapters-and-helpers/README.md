# Ports & Adapters Go Template - Domain Adapters and Helpers

If you haven't read yet, I recommend reading the `v1-very-simple/README.md` first.

This template divides the code in an equivalent way to the `v1-very-simple` one, except this one
is slightly more ergonomic by that I mean:

It moves the domain interfaces (ports) to packages closer to the adapters that implement them
which allows for better names, such as:

- `rest.Provider` instead of `domain.RestProvider`
- `cache.Provider` instead of `domain.CacheProvider`

Another difference is that on the `v1-very-simple` version we put both
helpers and adapters inside the `infra/` directory, but in this version
we actually split the infra directory into 2 new directories:

- `adapters/`
- `helpers/`

This is an important distinction because helpers are meant to be very simple
pieces of code that don't depend on the domain and also don't depend on
specific external technologies, except maybe helper libraries like testify
that are also similarly decoupled from technologies other than the stdlib.

But the same is not valid for adapters: they can both see the domain types
and also depend on external technologies such as databases, queues, email providers
and so on.

> Note: Not all adapters need to depend on domain types, and we could create
> a third directory to differentiate between the adapters that do, we
> could call them repositories, and the ones that don't. But in terms
> of logical decoupling rules that benefit the code maintainability
> this distinction is unimportant: all adapters are here to serve
> the needs of your domain as specified by the ports (domain interfaces).

## Reorganizing the infra package

The `infra/` packages still exist in this repo but they are organized
differently, we now have one upper directory for each interface implementation:

- The `infra/http` package was moved to `adapters/rest/http`
- The `infra/memorycache` package was moved to `adapters/cache/memorycache`
- The `infra/redis` package was moved to `adapters/cache/redis`

And so on.

For each of these new packages a new file `contracts.go` was created containing
only the relevant interfaces for that dependency, so now we have 3 new files with that name:

- `adapters/rest/contracts.go`
- `adapters/cache/contracts.go`
- `adapters/cache/contracts.go`

And the old `domain/contracts.go` was deleted, although in some situations
where an interface is implemented by a service or by more than one adapter
it would be ok to keep it on `domain/contracts.go`.

We also moved the two helpers that we used to have inside the infra package
to the helpers directory:

- The `infra/env` pkg was moved to `helpers/env`
- The `infra/maps` pkg was moved to `helpers/maps`

