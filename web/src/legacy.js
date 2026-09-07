const RETRY_LIMIT = 3;

function retry(fn) {
  for (var i = 0; i < RETRY_LIMIT; i++) {
    fn();
  }
}

module.exports = { retry, RETRY_LIMIT };
