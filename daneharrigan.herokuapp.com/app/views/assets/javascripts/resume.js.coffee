$ ->
  $('a[href^=#]').bind 'click', (e) ->
    anchor = $( $(@).attr 'href' )

    $('html, body').animate scrollTop: anchor.offset().top
    e.preventDefault()

  $('a[href^=http]').bind 'click', (e) ->
    window.open @href
    e.preventDefault()
