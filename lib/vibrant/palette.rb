module Vibrant
  class Palette
    attr_reader :swatches

    %w{vibrant light_vibrant dark_vibrant muted light_muted dark_muted}.each do |name|
      define_method(name) do
        @swatches.find {|s| s.name == name.split('_').map(&:capitalize).join }
      end
    end

    def pairs_with_contrast
      @swatches.combination(2).map do |s1, s2|
        [[s1, s2], s1.contrast_ratio(s2)]
      end
    end
  end
end
