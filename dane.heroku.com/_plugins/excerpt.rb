module Jekyll
  module Filters
    def more_link(content, url)
      return content unless content.include? "<!-- more -->"
      excerpt = content.split("<!-- more -->")[0]
      excerpt << %(<p class="read-more"><a href="#{url}#read-more">Read more...</a></p>)
    end

    def full_article(content)
      content.sub("<!-- more -->", %(<span id="read-more"></span>))
    end
  end
end
