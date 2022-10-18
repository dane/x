APP_ROOT = File.expand_path('.')
require 'bundler/setup'
require 'sinatra'
require 'sinatra/contrib'
require 'sass/plugin'
require 'coffee_script'
require  File.expand_path('app/helpers/javascript')
require  File.expand_path('config/application')
require  File.expand_path('app/controllers/resume_controller')
require  File.expand_path('app/controllers/assets_controller')

run Resume
