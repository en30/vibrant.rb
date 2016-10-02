def find_go_package(gopath, package)
  print "checking for #{package}... "
  if Dir.exist?("#{gopath}/src/#{package}")
    puts 'yes'
  else
    puts 'not'
    $stderr.puts <<-EOS

#{package} is not installed. Please
go get #{package}

EOS
    exit 1
  end
end

require 'mkmf'
find_executable('go')

gopath = ENV['GOPATH'] || arg_config('gopath')
unless gopath.is_a?(String)
  $stderr.puts 'Please set GOPATH environment variable or pass --with-gopath=GOPATH'
  exit 1
end

find_go_package(gopath, 'github.com/generaltso/vibrant')

$objs = []
def $objs.empty?; false ;end
create_makefile("vibrant/vibrant")
case `#{CONFIG['CC']} --version`
when /Free Software Foundation/
  ldflags = '-Wl,--unresolved-symbols=ignore-all'
when /clang/
  ldflags = '-undefined dynamic_lookup'
end
File.open('Makefile', 'a') do |f|
  f.write <<eom.gsub(/^ {8}/, "\t")
$(DLLIB): Makefile $(srcdir)/vibrant.go $(srcdir)/wrapper.go
        CGO_CFLAGS='$(INCFLAGS)' CGO_LDFLAGS='#{ldflags}' \
          go build -p 4 -buildmode=c-shared -o $(DLLIB) $(srcdir)/vibrant.go $(srcdir)/wrapper.go
eom
end
