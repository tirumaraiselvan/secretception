# secretception

A K8s custom controller that encrypts a secret

This is a very stupid controller that takes a secret and md5 hashes it and stores it back.

This was used for DEMO purposes in Koffee With Kubernetes (https://twitter.com/KoffeeWithK8S) meetup on 8 March 2018

## Usage

go build

./secretception --kubeconfig <your-kubeconfig-loc>

```
const algoliasearch = require('algoliasearch');

exports.function = async (req, res) => {
  const { event: { op, data } } = req.body;

  var client = algoliasearch(ALGOLIA_APP_ID, ALGOLIA_ADMIN_API_KEY);
  var index = client.initIndex('demo_serverless_etl_app');

  index.addObjects([data.new], function(err, content) {
      console.log(content);
    });
};
```
