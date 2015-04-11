require 'formula'

class I18n4goBin < Formula
  homepage 'https://github.com/maximilien/i18n4go'
  url 'https://github.com/maximilien/i18n4go/archive/v0.2.0.tar.gz'
  sha1 'f2c0073672be91c185f81d5c5bd747aa1763c393'
  version '0.1.1'

  def install
    system 'curl -O https://raw.github.com/maximilien/i18n4go/v0.2.0/LICENSE'
    system "#{bin}/build"
    bin.install 'out/i18n4go' => 'i18n4go'
    doc.install 'LICENSE'
  end

  test do
    system "#{bin}/i18n4go"
  end
end