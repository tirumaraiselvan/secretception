const algoliasearch = require('algoliasearch');

exports.function = async (req, res) => {
  const { event: { op, data } } = req.body;

  var client = algoliasearch(ALGOLIA_APP_ID, ALGOLIA_ADMIN_API_KEY);
  var index = client.initIndex('demo_serverless_etl_app');

  index.addObjects([data.new], function(err, content) {
      console.log(content);
    });
};
