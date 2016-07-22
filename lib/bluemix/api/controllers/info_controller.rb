require 'bluemix/api/controllers/base_controller'
require 'sinatra/swagger-exposer/swagger-exposer'

module Bluemix::BM
  module Api::Controllers
    class InfoController < BaseController

      register Sinatra::SwaggerExposer

      type 'Status', {
        :required => [:name, :version],
        :properties => {
          :name => {
            :type => String,
            :example => 'Name',
            :description => 'The name of this running server',
          },
          :version => {
            :type => String,
            :example => '1.0.0',
            :description => 'The version of this running server',
          },
        },
      }

      endpoint_summary 'info about this server'
      endpoint_description 'Returns basic information about this SoftLayer Baremetal server'
      endpoint_tags 'info'
      endpoint_response 200, ['Status'], 'Standard response'
      get '/info' do
        status = {
          'name' => "Bluemix Provision Server",
          'version' => "0.1"
        }
        Bluemix::BM::Common::Message.success( status )
      end
    end
  end
end
