require 'bluemix/api/controllers/baremetal_controller'
require 'bluemix/api/controllers/task_controller'
require 'bluemix/api/controllers/info_controller'
require 'bluemix/api/controllers/misc_controller'
require 'bluemix/api/controllers/stemcell_controller'
require 'bluemix/api/controllers/softlayer_controller'
require 'sinatra/swagger-exposer/swagger-exposer'

module Bluemix::BM
  module Api
    class Controller < Sinatra::Base
      register Sinatra::SwaggerExposer
      general_info(
          {
              version: '0.0.1',
              title: 'SoftLayer Baremetal Server',
              description: 'A baremetal server for BOSH SoftLayer CPI',
              license: {
                  name: 'MIT',
                  url: 'http://opensource.org/licenses/MIT'
              }
          }
      )

      use Controllers::BaremetalController
      use Controllers::TaskController
      use Controllers::InfoController
      use Controllers::MiscController
      use Controllers::SoftlayerController
      use Controllers::StemcellController
    end
  end
end
