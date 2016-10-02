require "vibrant/version"
begin
  require "vibrant/#{RUBY_VERSION[/\d+\.\d+/]}/vibrant"
rescue LoadError
  require "vibrant/vibrant"
end

class Vibrant
end
