---
date:   2018-06-04
title:  "Chainable, auto-refined APIs in TypeScript"
# excerpt: "Pow
---

I'm creating a query builder for DynamoDB.
It's a chainable, JavaScript-ey API.  Each method call updates the internal state and returns a reference to `this`.
However, certain method calls only make sense *after* others have been called.
For example, it doesn't make sense to specify row filters before you've specified the table to visit.  Not only would the code be confusing to read,
you would lose the benefits of tab-completion when coding the filter.  When you specify the table first, the API understands the table's schema and
can offer tab-completion for column names.

Also, some methods are contradictory.  For example, in DynamoDB you can either "scan" or "query" a table.  I expose these as a `scan()` and a `query()` method.
Once you've called one, it doesn't make sense to call the other.

I wanted to devise a way of filtering the methods shown in tab-completion.  Initially, you're only shown methods to specify the query type: 'query' or 'scan'.
After calling one of them, you're shown the method to select a table to query.