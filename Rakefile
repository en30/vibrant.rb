require "bundler/gem_tasks"
require "rspec/core/rake_task"
require 'rake/extensiontask'

RSpec::Core::RakeTask.new(:spec)

task :default => [:compile, :spec]

spec = eval File.read('vibrant.gemspec')
Rake::ExtensionTask.new('vibrant', spec) do |ext|
  ext.ext_dir = 'ext/vibrant'
  ext.lib_dir = File.join(*['lib', 'vibrant', ENV['FAT_DIR']].compact)
  ext.source_pattern = "*.{c,cpp,go}"
end
