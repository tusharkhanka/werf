{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}

Search reads through all of the repositories configured on the system, and
looks for matches. Search of these repositories uses the metadata stored on
the system.

It will display the latest stable versions of the charts found. If you
specify the --devel flag, the output will include pre-release versions.
If you want to search using a version constraint, use --version.

Examples:

    # Search for stable release versions matching the keyword "nginx"
    $ helm search repo nginx

    # Search for release versions matching the keyword "nginx", including pre-release versions
    $ helm search repo nginx --devel

    # Search for the latest stable release for nginx-ingress with a major version of 1
    $ helm search repo nginx-ingress --version ^1.0.0

Repositories are managed with 'helm repo' commands.


{{ header }} Syntax

```shell
werf helm-v3 search repo [keyword] [flags] [options]
```

{{ header }} Options

```shell
      --devel=false:
            use development versions (alpha, beta, and release candidate releases), too. Equivalent 
            to version '>0.0.0-0'. If --version is set, this is ignored
  -h, --help=false:
            help for repo
      --max-col-width=50:
            maximum column width for output table
  -o, --output=table:
            prints the output in the specified format. Allowed values: table, json, yaml
  -r, --regexp=false:
            use regular expressions for searching repositories you have added
      --version='':
            search using semantic versioning constraints on repositories you have added
  -l, --versions=false:
            show the long listing, with each version of each chart on its own line, for             
            repositories you have added
```

