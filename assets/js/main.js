$(function () {
  var count = 0;

  function initSearchBox() {
    var pages = new Bloodhound({
      datumTokenizer: Bloodhound.tokenizers.obj.nonword(['title','content']),
      queryTokenizer: Bloodhound.tokenizers.whitespace,
      sorter: function(a, b) {
        var input_string = $('#search-box').val();
        return distance(a.title, input_string) - distance(b.title, input_string);
      },

      prefetch: {
        url: '/search.json',
        transform : function(response) {
          return pages.sorter(response)
        },
        cache: false
      }
    });

    $('#search-box').typeahead({
      minLength: 0,
      highlight: true
    }, {
        name: 'pages',
        limit: 10,
        display: 'title',
        templates: {
          footer: function(context) {
            if (context.suggestions.length < 10) {
              return "";
            }
 
            var query = $("#search-box").typeahead("val");
            return '<div class="tt-suggestion"><a href="/search?q=' + query + '">See all results...</a></div>';
          }
        },
        source: pages
      });

    $('#search-box').bind('typeahead:select', function (ev, suggestion) {
      window.location.href = suggestion.url;
    });
  }

  function styleContentToMD() {
    $('#markdown-content-container table').addClass('table');
    $('#markdown-content-container img').addClass('img-responsive');
  }

  initSearchBox();
  styleContentToMD();
});
