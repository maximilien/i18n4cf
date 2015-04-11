require 'formula'

class I18n4goSrc < Formula
  homepage 'https://github.com/maximilien/i18n4go'
  url 'https://github.com/maximilien/i18n4go.git', :tag => 'v0.2.0'

  head 'https://github.com/maximilien/i18n4go.git', :branch => 'master'

  depends_on 'go' => :build

  def install
    system 'bin/build'
    bin.install 'out/i18n4go'
    doc.install 'LICENSE'
  end

  test do
    system "#{bin}/i18n4go"
  end
end