# Module Summary

## Dynamic Routing

- You can pass dynamic path segments by adding a ":" to the Express router path
- The name you add after ":" is the name by which you can extract the adata on req.params
- Optional (query) parameters can also be passed (?params=value&b=2) and extracted (req.query.myParam)

## More on Models

- A Cart model was added - it holds static methods only
- You can interact between models (e.g. delete cart item if a product is deleted)
- Working with files for data storage is suboptimal for bigger amounts of data
