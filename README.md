你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
<p align="center"><img src="logo.png" width="150" height="150"></p>

<p align="center">
    <a href="https://godoc.org/github.com/gobuffalo/pop"><img src="https://godoc.org/github.com/gobuffalo/pop?status.svg" alt="GoDoc" /></a>
    <a href="https://dev.azure.com/stanislasmichalak/pop/_build/latest?definitionId=1&branchName=master"><img src="https://dev.azure.com/stanislasmichalak/pop/_apis/build/status/gobuffalo.pop?branchName=master" alt="Build Status" /></a>
</p>

# POP

## A Tasty Treat For All Your Database Needs

So what does Pop do exactly? Well, it wraps the absolutely amazing [https://github.com/jmoiron/sqlx](https://github.com/jmoiron/sqlx) library. It cleans up some of the common patterns and work flows usually associated with dealing with databases in Go.

Pop makes it easy to do CRUD operations, run migrations, and build/execute queries.

Pop, by default, follows conventions that were influenced by the [ActiveRecord](http://www.rubyonrails.org) Ruby gem. What does this mean?

* Tables must have an "id" column and a corresponding "ID" field on the `struct` being used.
* If there is a `timestamp` column named `created_at`, and a `CreatedAt time.Time` attribute on the `struct`, it will be set with the current time when the record is created.
* If there is a `timestamp` column named `updated_at`, and a `UpdatedAt time.Time` attribute on the `struct`, it will be set with the current time when the record is updated.
* Default database table names are lowercase, plural, and underscored versions of the `struct` name. Examples: User{} is "users", FooBar{} is "foo_bars", etc...

Want to know more? Take a look at the documentation!

## Documentation

Please visit [http://gobuffalo.io](https://gobuffalo.io/docs/db/getting-started) for the latest documentation, examples, and more.

### Quick Start
* [CLI Installation](https://gobuffalo.io/docs/db/toolbox)
* [Configuration](https://gobuffalo.io/docs/db/configuration)

## Shoulders of Giants

Pop would not be possible if not for all of the great projects it depends on. Please see [SHOULDERS.md](SHOULDERS.md) to see a list of them.

## Contributing

First, thank you so much for wanting to contribute! It means so much that you care enough to want to contribute. We appreciate every PR from the smallest of typos to the be biggest of features.

To contribute, please read the contribution guidelines: [CONTRIBUTING](.github/CONTRIBUTING.md)
