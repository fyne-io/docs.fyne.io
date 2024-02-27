$(function () {

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
        display: 'title',
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
