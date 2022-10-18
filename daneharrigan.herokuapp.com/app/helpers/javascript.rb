module JavaScript
  def javascript(path)
    %{<script type="text/javascript" src="#{path}" charset="utf-8"></script>}
  end
end
