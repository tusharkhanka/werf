{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}

Manage the dependencies of a chart.

Helm charts store their dependencies in 'charts/'. For chart developers, it is
often easier to manage dependencies in 'Chart.yaml' which declares all
dependencies.

The dependency commands operate on that file, making it easy to synchronize
between the desired dependencies and the actual dependencies stored in the
'charts/' directory.

For example, this Chart.yaml declares two dependencies:

    # Chart.yaml
    dependencies:
    - name: nginx
      version: "1.2.3"
      repository: "[https://example.com/charts](https://example.com/charts)"
    - name: memcached
      version: "3.2.1"
      repository: "[https://another.example.com/charts](https://another.example.com/charts)"


The 'name' should be the name of a chart, where that name must match the name
in that chart's 'Chart.yaml' file.

The 'version' field should contain a semantic version or version range.

The 'repository' URL should point to a Chart Repository. Helm expects that by
appending '/index.yaml' to the URL, it should be able to retrieve the chart
repository's index. Note: 'repository' can be an alias. The alias must start
with 'alias:' or '@'.

Starting from 2.2.0, repository can be defined as the path to the directory of
the dependency charts stored locally. The path should start with a prefix of
"[file://](file://)". For example,

    # Chart.yaml
    dependencies:
    - name: nginx
      version: "1.2.3"
      repository: "[file://](file://)../dependency_chart/nginx"

If the dependency chart is retrieved locally, it is not required to have the
repository added to helm by "helm add repo". Version matching is also supported
for this case.


{{ header }} Options

```shell
  -h, --help=false:
            help for dependency
```

