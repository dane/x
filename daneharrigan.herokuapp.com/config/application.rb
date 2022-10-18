class Resume < Sinatra::Base
  set :views, "#{APP_ROOT}/app/views"
  set :public, "#{APP_ROOT}/public"

  register Sinatra::Contrib
  helpers JavaScript
end
