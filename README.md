# Vibrant

This is a ruby binding of [https://github.com/generaltso/vibrant](https://github.com/generaltso/vibrant), which is a go port of the Android Palette class.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'vibrant', github: 'en30/vibrant.rb'
```

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install vibrant

## Usage

```ruby
palette = Vibrant::Palette.from_file(PATH_TO_IMAGE_FILE) # or from_url(URL)
=> #<Vibrant::Palette:0x007f7f741884b8 @swatches=[#<Vibrant::Swatch:0x007f7f74188468>, #<Vibrant::Swatch:0x007f7f74188440>, #<Vibrant::Swatch:0x007f7f741883c8>, #<Vibrant::Swatch:0x007f7f741883a0>]>

palette.swatches #=> [#<Vibrant::Swatch:id>, #<Vibrant::Swatch:id>, #<Vibrant::Swatch:id>, #<Vibrant::Swatch:id>]

palette.vibrant       #=> #<Vibrant::Swatch:id>
palette.light_vibrant #=> #<Vibrant::Swatch:id>
palette.dark_vibrant  #=> #<Vibrant::Swatch:id>
palette.muted         #=> #<Vibrant::Swatch:id>
palette.light_muted   #=> #<Vibrant::Swatch:id>
palette.dark_muted    #=> #<Vibrant::Swatch:id>

highest_contrast_swathes, c = palette.pairs_with_contrast.sort_by {|((s1, s2), c)| c }.last

swatch = pallete.swatches[0]
swatch.population #=> 100
swatch.name  #=> "LightVibrant"
swatch.color #=> 12032073
swatch.rgb   #=> [183, 152, 73]
swatch.hex   #=> "#b79849"
```


## Development

After checking out the repo, run `bin/setup` to install dependencies. Then, run `rake spec` to run the tests. You can also run `bin/console` for an interactive prompt that will allow you to experiment.

To install this gem onto your local machine, run `bundle exec rake install`. To release a new version, update the version number in `version.rb`, and then run `bundle exec rake release`, which will create a git tag for the version, push git commits and tags, and push the `.gem` file to [rubygems.org](https://rubygems.org).

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/en30/vibrant.rb


## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
