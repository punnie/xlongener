var getLongenedURL = function (url, callback) {
  var req = new XMLHttpRequest();
  req.onload = function () {
    callback(window.location + this.responseText);
  }

  req.open('GET', '/create?url=' + url)
  req.send()
}

document.getElementById('new_url').onsubmit = function () {
  var result_url = document.getElementById('result_url');
  var result     = document.getElementById('result');

  result_url.value = 'loading...';
  result.style.display = 'block';

  getLongenedURL(document.getElementById('new_url_url').value, function (url) {
    result_url.value = url;
    result_url.select();
  });

  return false;
}
