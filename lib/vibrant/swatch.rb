module Vibrant
  class Swatch
    def hex
      ?# + color.to_s(16).rjust(6, '0')
    end

    def rgb
      (0..2).reverse_each.map {|i| (color  >> (i*8)) % 256 }
    end

    # https://www.w3.org/TR/2008/REC-WCAG20-20081211/#contrast-ratiodef
    def contrast_ratio(the_other)
      r1, r2 = [relative_luminance + 0.05, the_other.relative_luminance + 0.05].minmax
      r2 / r1
    end

    # https://www.w3.org/TR/2008/REC-WCAG20-20081211/#relativeluminancedef
    def relative_luminance
      r, g, b = rgb.map do |c|
        sc = c / 255.0
        sc <= 0.03928 ? (sc / 12.92) : ((sc + 0.055) / 1.055) ** 2.4
      end
      0.2126 * r + 0.7152 * g + 0.0722 * b
    end
  end
end
