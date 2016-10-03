module Vibrant
  class Swatch
    def hex
      "##{color.to_s(16)}"
    end

    def rgb
      (0..2).reverse_each.map {|i| (color  >> (i*8)) % 256 }
    end
  end
end
