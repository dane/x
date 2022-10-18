class Resume
  get '/assets/javascripts/:path' do
    coffee :"assets/javascripts/#{params[:path]}"
  end

  get '/assets/stylesheets/:path' do
    scss :"assets/stylesheets/#{params[:path]}"
  end
end
