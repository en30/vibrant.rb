# coding: utf-8
lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'vibrant/version'

Gem::Specification.new do |spec|
  spec.name          = "vibrant"
  spec.version       = Vibrant::VERSION
  spec.authors       = ["en30"]
  spec.email         = ["en30.git@gmail.com"]

  spec.summary       = %q{ruby binding for vibrant}
  spec.description   = %q{ruby binding for vibrant}
  spec.homepage      = "https://github.com/en30/vibrant.rb"
  spec.license       = "MIT"

  spec.files         = `git ls-files -z`.split("\x0").reject { |f| f.match(%r{^(test|spec|features)/}) }
  spec.bindir        = "exe"
  spec.executables   = spec.files.grep(%r{^exe/}) { |f| File.basename(f) }
  spec.require_paths = ["lib"]
  spec.extensions    = ["ext/vibrant/extconf.rb"]

  spec.add_development_dependency "bundler", "~> 1.12"
  spec.add_development_dependency "rake", "~> 10.0"
  spec.add_development_dependency "rspec", "~> 3.0"
  spec.add_development_dependency "rake-compiler"
end
