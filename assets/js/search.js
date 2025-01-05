
function getQueryParams(k){
 var p={};
 location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi,function(s,k,v){p[k]=v})
 return k?p[k]:p;
}

function doSearch(index){
  // Initialize lunr with the fields it will be searching on. I've given title
  // a boost of 10 to indicate matches on this field are more important.
  var idx = lunr(function () {
    this.ref('id');
    this.field('title', { boost: 10 });
    this.field('keywords');
    this.field('tags');
    this.field('body');

    for (var key in index) { // Add the data to lunr
      this.add({
        'id': key,
        'title': index[key].title,
        'keywords': index[key].keywords,
        'tags': index[key].tags,
        'body': index[key].content
      });
    }
  });

    var results = idx.search("*"+searchTerm+"*"); // Get lunr to perform a search
    displaySearchResults(results, index); // We'll write this in the next section
}

function displaySearchResults(results, index) {
  var searchResults = document.getElementById('search-results');

  if (results.length) { // Are there any results?
    var appendString = '';

    for (var i = 0; i < results.length; i++) {  // Iterate over the results
      var item = index[results[i].ref];
      appendString += '<li><a href="' + item.url + '"><h3>' + item.title + '</h3></a>';
      if (item.tags != "") {
        appendString += '<span class="search-tag">' + item.tags + '</span>';
      }
      appendString += '<p>' + item.content.substring(0, 150) + '...</p></li>';
    }

    searchResults.innerHTML = appendString;
  } else {
    searchResults.innerHTML = '<li>No results found</li>';
  }
}

var searchTerm = getQueryParams('q');

if (searchTerm) {
  document.getElementById('search-box-full').setAttribute("value", searchTerm);
  document.getElementById('search-box').setAttribute("value", searchTerm);

$.getJSON('/search.json', function(data) {
  doSearch(data);
});
}
